package api_server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/emq_cmd"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/emq/topic_publish_handler"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func EditFStrategy(c *gin.Context) {
	vehicleId := c.PostForm("vehicle_id")
	fstrategyId := c.Param("fstrategy_id")
	setTypeP := c.PostForm("type")
	handleModeP := c.PostForm("handle_mode")
	dips := c.PostForm("dips")
	dstPorts := c.PostForm("dst_ports")

	//logger.Logger.Info("%s vehicle_id:%s,fstrategy_id:%s,type:%s,handle_mode:%s,dips:%s,dst_ports:%s",
	//	util.RunFuncName(),vehicleId,fstrategyId,setTypeP,handleModeP,dips,dstPorts)
	//
	//logger.Logger.Print("%s vehicle_id:%s,fstrategy_id:%s,type:%s,handle_mode:%s,dips:%s,dst_ports:%s",
	//	util.RunFuncName(),vehicleId,fstrategyId,setTypeP,handleModeP,dips,dstPorts)

	argsTrimsEmpty := util.RrgsTrimsEmpty(fstrategyId, vehicleId, setTypeP, handleModeP, dips, dstPorts)
	sTypeValid := util.IsEleExistInSlice(setTypeP, []interface{}{
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOTYPEDEFAULT)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOWHITEMODE)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOBLACKMODE))})
	handleModeValid := util.IsEleExistInSlice(handleModeP, []interface{}{
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_MODEDEFAULT)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_PREVENTWARNING)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_WARNING))})
	if argsTrimsEmpty || !sTypeValid || !handleModeValid {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
	}
	setType, _ := strconv.Atoi(setTypeP)
	handleMode, _ := strconv.Atoi(handleModeP)

	fdipSlice := []string{}
	//筛选dip
	dipSlice := strings.Split(dips, ",")
	for _, dip := range dipSlice {
		destIpValid := util.IpFormat(dip)
		if destIpValid && !util.IsExistInSlice(dip, fdipSlice) {
			fdipSlice = append(fdipSlice, dip)
		}
	}

	fdportSlice := []uint32{}
	//筛选dstPort
	dstPortSlice := strings.Split(dstPorts, ",")

	for _, dport := range dstPortSlice {
		destPortValid := util.VerifyIpPort(dport)
		dpInt, _ := strconv.Atoi(dport)
		if destPortValid && !util.IsExistInSlice(uint32(dpInt), fdportSlice) {
			fdportSlice = append(fdportSlice, uint32(dpInt))
		}
	}
	//////////////////////////////////////////////////////////////
	//查看该vehicle是否存在
	vehicleFStrategy, err := model.GetVehicleFStrategy(
		"fstrategy_vehicles.vehicle_id = ? and fstrategies.fstrategy_id = ?", []interface{}{vehicleId, fstrategyId}...)
	if vehicleFStrategy.FstrategyVehicleId == "" {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	logger.Logger.Print("%s vehicleFStrategy:%+v", util.RunFuncName(), vehicleFStrategy)

	//更新策略信息
	vehicleFstrategy := &model.Fstrategy{
		FstrategyId: fstrategyId,
		Type:        uint8(setType),
		HandleMode:  uint8(handleMode),
	}
	vehicleFstrategyModelBase := model_base.ModelBaseImpl(vehicleFstrategy)

	attrs := map[string]interface{}{
		"handle_mode": vehicleFstrategy.HandleMode,
		"type":        vehicleFstrategy.Type,
	}

	if err := vehicleFstrategyModelBase.UpdateModelsByCondition(attrs, "fstrategy_id = ?",
		[]interface{}{vehicleFstrategy.FstrategyId}...); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//FstrategyVehicle不更改

	//遍历FstrategyVehicleItem表
	fstrategyItems := map[string][]string{}
	var vehicleFstrategyItems []string

	for _, dip := range fdipSlice {
		for _, dport := range fdportSlice {
			fstrategyItem := &model.FstrategyItem{
				FstrategyItemId: util.RandomString(32),
				VehicleId:       vehicleId,
				DstIp:           dip,
				DstPort:         dport,
			}
			modelBase := model_base.ModelBaseImpl(fstrategyItem)

			err, fstrategyItemRecordNotFound := modelBase.GetModelByCondition(
				"vehicle_id = ? and dst_ip = ? and dst_port = ?",
				[]interface{}{fstrategyItem.VehicleId, fstrategyItem.DstIp, fstrategyItem.DstPort}...)
			if err != nil {
				continue
			}

			if fstrategyItemRecordNotFound {
				if err := modelBase.InsertModel(); err != nil {
					continue
				}
			}
			if !util.IsExistInSlice(fstrategyItem.FstrategyItemId, vehicleFstrategyItems) {
				vehicleFstrategyItems = append(vehicleFstrategyItems, fstrategyItem.FstrategyItemId)
			}
		}
	}
	fstrategyItems[vehicleId] = vehicleFstrategyItems
	logger.Logger.Print("%s fstrategyItems:::%+v", util.RunFuncName(), fstrategyItems)

	//找到FstrategyItemId(FstrategyVehicleItem表中)
	//在FstrategyItemId(FstrategyItem表中)不存在的值

	var fstrategyVehicleItemIds []string
	_ = mysql.QueryPluckByModelWhere(&model.FstrategyVehicleItem{}, "fstrategy_item_id", &fstrategyVehicleItemIds,
		"fstrategy_vehicle_id = ?", vehicleFStrategy.FstrategyVehicleId)

	logger.Logger.Print("%s fstrategyItemIds:::%+v", util.RunFuncName(), fstrategyVehicleItemIds)

	//如果没有在里面，就是被删除的，需要改delete标志位
	newFstrategyItemIds := fstrategyItems[vehicleId]
	var needDeleFstrategyItemIds []string
	for _, fstrategyItemId := range fstrategyVehicleItemIds {

		if !util.IsExistInSlice(fstrategyItemId, newFstrategyItemIds) {
			needDeleFstrategyItemIds = append(needDeleFstrategyItemIds, fstrategyItemId)
		}
	}
	//置成标志位
	fstrategyItem := &model.FstrategyItem{}
	err = fstrategyItem.SoftDeleModelImpl("fstrategy_item_id in (?)", needDeleFstrategyItemIds)

	//删除FstrategyVehicleItem表
	fstrategyVehicleItem := &model.FstrategyVehicleItem{}
	fstrategyVehicleItemModelBase := model_base.ModelBaseImpl(fstrategyVehicleItem)
	err = fstrategyVehicleItemModelBase.DeleModelsByCondition("fstrategy_vehicle_id = ?", vehicleFStrategy.FstrategyVehicleId)
	if err != nil {
		return
	}

	//添加FstrategyVehicleItem表
	for _, fstrategyItemId := range newFstrategyItemIds {
		fstrategyVehicleItem := &model.FstrategyVehicleItem{
			FstrategyVehicleId: vehicleFStrategy.FstrategyVehicleId,
			FstrategyItemId:    fstrategyItemId,
		}

		fstrategyVehicleItemModelBase := model_base.ModelBaseImpl(fstrategyVehicleItem)
		if err := fstrategyVehicleItemModelBase.InsertModel(); err != nil {
			continue
		}
	}

	//更新
	fstrategyCmd := &emq_cmd.FStrategySetCmd{
		VehicleId: vehicleId,
		TaskType:  int(protobuf.Command_FLOWSTRATEGY_ADD),

		FstrategyId: fstrategyId,
		Type:        setType,
		HandleMode:  handleMode,
		Enable:      true,
		GroupId:     "", //目前不实现
	}
	topic_publish_handler.GetPublishService().PutMsg2PublicChan(fstrategyCmd)

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateStrategySuccessMsg, "")
	c.JSON(http.StatusOK, retObj)

}

func GetFStrategys(c *gin.Context) {
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")

	argsTrimsEmpty := util.RrgsTrimsEmpty(pageSizeP, pageIndexP)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty pageSizeP:%s,pageIndexP:%s", util.RunFuncName(), pageSizeP, pageIndexP)
		logger.Logger.Print("%s argsTrimsEmpty pageSizeP:%s,pageIndexP:%s", util.RunFuncName(), pageSizeP, pageIndexP)
	}

	pageSize, _ := strconv.Atoi(pageSizeP)
	pageIndex, _ := strconv.Atoi(pageIndexP)

	fstrategys := []*model.Fstrategy{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.Fstrategy{})

	err := modelBase.GetModelPaginationByCondition(pageIndex, pageSize,
		&total, &fstrategys, "",
		[]interface{}{}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrategyListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"fstrategys":  fstrategys,
		"total_count": total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFStrategyListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func GetFStrategy(c *gin.Context) {
	fstrategyId := c.Param("fstrategy_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(fstrategyId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty fstrategyId:%s", util.RunFuncName(), fstrategyId)
		logger.Logger.Print("%s argsTrimsEmpty fstrategyId:%s", util.RunFuncName(), fstrategyId)
	}
	fstrategyInfo := &model.Fstrategy{
		FstrategyId: fstrategyId,
	}

	modelBase := model_base.ModelBaseImpl(fstrategyInfo)

	err, recordNotFound := modelBase.GetModelByCondition("fstrategy_id = ?", []interface{}{fstrategyInfo.FstrategyId}...)

	if err != nil {
		logger.Logger.Error("%s fstrategy_id:%s,err:%s", util.RunFuncName(), fstrategyInfo.FstrategyId, err)
		logger.Logger.Print("%s fstrategy_id:%s,err:%s", util.RunFuncName(), fstrategyInfo.FstrategyId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s strategy_id:%s,recordNotFound", util.RunFuncName(), fstrategyInfo.FstrategyId)
		logger.Logger.Print("%s strategy_id:%s,recordNotFound", util.RunFuncName(), fstrategyInfo.FstrategyId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	responseData := map[string]interface{}{
		"fstrategy": fstrategyInfo,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func AddFStrategy(c *gin.Context) {
	vehicleIdsP := c.PostForm("vehicle_ids")
	sType := c.PostForm("type")
	handleMode := c.PostForm("handle_mode")
	//dips := c.PostForm("dips")
	//dstPorts := c.PostForm("dst_ports")
	diports := c.PostForm("diports")

	sTypeValid := util.IsEleExistInSlice(sType, []interface{}{
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOTYPEDEFAULT)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOWHITEMODE)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOBLACKMODE))})

	handleModeValid := util.IsEleExistInSlice(handleMode, []interface{}{
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_MODEDEFAULT)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_PREVENTWARNING)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_WARNING))})

	logger.Logger.Print("%s vehicleIdsP:%s,sType:%s,handleMode:%s,diports:%v",
		util.RunFuncName(), vehicleIdsP, sType, handleMode, diports)
	logger.Logger.Info("%s vehicleIdsP:%s,sType:%s,handleMode:%s,diports:%v",
		util.RunFuncName(), vehicleIdsP, sType, handleMode, diports)

	argsTrimsEmpty := util.RrgsTrimsEmpty(sType, handleMode, diports, vehicleIdsP)

	diportsMap := map[string][]uint32{}
	err := json.Unmarshal([]byte(diports), &diportsMap)
	if argsTrimsEmpty ||
		!sTypeValid ||
		!handleModeValid ||
		err != nil {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Print("%s vehicleIdsP:%s,sType:%s,handleMode:%s,diports:%v",
			util.RunFuncName(), vehicleIdsP, sType, handleMode, diports)
		logger.Logger.Info("%s vehicleIdsP:%s,sType:%s,handleMode:%s,diports:%v",
			util.RunFuncName(), vehicleIdsP, sType, handleMode, diports)
		return
	}

	vStype, _ := strconv.Atoi(sType)
	vHandleMode, _ := strconv.Atoi(handleMode)

	//for dip, ports := range diportsMap {
	//	destIpValid := util.IpFormat(dip)
	//	if !destIpValid {
	//		delete(diportsMap, dip)
	//	} else {
	//		for _, dport := range ports {
	//			dporStr := strconv.Itoa(int(dport))
	//			destPortValid := util.VerifyIpPort(dporStr)
	//			if !destPortValid {
	//
	//			}
	//		}
	//	}
	//
	//}

	//fdipSlice := []string{}
	////筛选dip
	//dipSlice := strings.Split(dips, ",")
	//for _, dip := range dipSlice {
	//	destIpValid := util.IpFormat(dip)
	//	if destIpValid && !util.IsExistInSlice(dip, fdipSlice) {
	//		fdipSlice = append(fdipSlice, dip)
	//	}
	//}
	//
	//fdportSlice := []uint32{}
	////筛选dstPort
	//dstPortSlice := strings.Split(dstPorts, ",")
	//
	//for _, dport := range dstPortSlice {
	//	destPortValid := util.VerifyIpPort(dport)
	//	dpInt, _ := strconv.Atoi(dport)
	//	if destPortValid && !util.IsExistInSlice(uint32(dpInt), fdportSlice) {
	//		fdportSlice = append(fdportSlice, uint32(dpInt))
	//	}
	//}

	//找出合法的vehicle
	vehicleIdSlice := strings.Split(vehicleIdsP, ",")
	var vehicleIds []string
	_ = mysql.QueryPluckByModelWhere(&model.VehicleInfo{}, "vehicle_id", &vehicleIds, "vehicle_id in (?)", vehicleIdSlice)

	//fstrategy_items table
	fstrategyItems := map[string][]string{}
	for _, vehicleId := range vehicleIds {
		var vehicleFstrategyItems []string
		for dip, ports := range diportsMap {
			for _, dport := range ports {
				fstrategyItem := &model.FstrategyItem{
					FstrategyItemId: util.RandomString(32),
					VehicleId:       vehicleId,
					DstIp:           dip,
					DstPort:         dport,
				}
				modelBase := model_base.ModelBaseImpl(fstrategyItem)

				err, fstrategyItemrecordNotFound := modelBase.GetModelByCondition(
					"vehicle_id = ? and dst_ip = ? and dst_port = ?",
					[]interface{}{fstrategyItem.VehicleId, fstrategyItem.DstIp, fstrategyItem.DstPort}...)
				if err != nil {
					continue
				}

				if fstrategyItemrecordNotFound {
					if err := modelBase.InsertModel(); err != nil {
						continue
					}
				}
				if !util.IsExistInSlice(fstrategyItem.FstrategyItemId, vehicleFstrategyItems) {
					vehicleFstrategyItems = append(vehicleFstrategyItems, fstrategyItem.FstrategyItemId)
				}
			}
		}
		fstrategyItems[vehicleId] = vehicleFstrategyItems
	}

	//fstrategy table
	fstrategy := &model.Fstrategy{
		FstrategyId: util.RandomString(32),
		Type:        uint8(vStype),
		HandleMode:  uint8(vHandleMode),
		Enable:      true,
	}

	fstrategyModelBase := model_base.ModelBaseImpl(fstrategy)
	if err := fstrategyModelBase.InsertModel(); err != nil {
		//todo
	}

	for _, vehicleId := range vehicleIds {
		//FstrategyVehicle table
		fstrategyVehicle := &model.FstrategyVehicle{
			FstrategyVehicleId: util.RandomString(32),
			FstrategyId:        fstrategy.FstrategyId,
			VehicleId:          vehicleId,
		}
		fstrategyVehicleModelBase := model_base.ModelBaseImpl(fstrategyVehicle)
		if err := fstrategyVehicleModelBase.InsertModel(); err != nil {
			continue
		}

		vehicleFsItems := fstrategyItems[vehicleId]
		//learningResultIds table
		for _, item := range vehicleFsItems {
			fstrategyVehicleItem := &model.FstrategyVehicleItem{
				FstrategyVehicleId: fstrategyVehicle.FstrategyVehicleId,
				FstrategyItemId:    item,
			}

			fstrategyVehicleItemModelBase := model_base.ModelBaseImpl(fstrategyVehicleItem)

			if err := fstrategyVehicleItemModelBase.InsertModel(); err != nil {
				continue
			}
		}
	}

	//下发策略
	for _, vehicleId := range vehicleIds {
		fstrategyCmd := &emq_cmd.FStrategySetCmd{
			VehicleId: vehicleId,
			TaskType:  int(protobuf.Command_FLOWSTRATEGY_ADD),

			FstrategyId: fstrategy.FstrategyId,
			Type:        vStype,
			HandleMode:  vHandleMode,
			Enable:      true,
			GroupId:     "", //目前不实现
		}
		topic_publish_handler.GetPublishService().PutMsg2PublicChan(fstrategyCmd)
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddFStrategySuccessMsg, fstrategy)
	c.JSON(http.StatusOK, retObj)
}

func DeleFStrategy(c *gin.Context) {
	fstrategyId := c.Param("fstrategy_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(fstrategyId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty fstrategy_id:%s argsTrimsEmpty", util.RunFuncName(), fstrategyId)
		logger.Logger.Print("%s argsTrimsEmpty fstrategy_id:%s argsTrimsEmpty", util.RunFuncName(), fstrategyId)
		return
	}

	//连表查询
	ftrategyVehicleItems, _ := model.GetFlowStrategyVehicleItems(
		"fstrategies.fstrategy_id = ?", []interface{}{fstrategyId}...)

	fVehicleIdMap := map[string]string{}
	fstrategyVehicleIdMap := map[string]string{}
	fstrategyItemIdMap := map[string]string{}

	for _, ftrategyVehicleItem := range ftrategyVehicleItems {
		fVehicleIdMap[ftrategyVehicleItem.VehicleId] = "1"
		fstrategyVehicleIdMap[ftrategyVehicleItem.FstrategyVehicleId] = "1"
		fstrategyItemIdMap[ftrategyVehicleItem.FstrategyItemId] = "1"
	}
	var fVehicleIdMapSlice []string
	var fstrategyVehicleIdMapSlice []string
	var fstrategyItemIdMapSlice []string

	for k := range fVehicleIdMap {
		fmt.Println("fVehicleIdMap::", k)
		fVehicleIdMapSlice = append(fVehicleIdMapSlice, k)
	}
	for k := range fstrategyVehicleIdMap {
		fmt.Println("fstrategyVehicleIdMap::", k)
		fstrategyVehicleIdMapSlice = append(fstrategyVehicleIdMapSlice, k)
	}
	for k := range fstrategyItemIdMap {
		fmt.Println("fstrategyItemIdMap::", k)
		fstrategyItemIdMapSlice = append(fstrategyItemIdMapSlice, k)
	}

	//dele Fstrategy
	fstrategyObj := &model.Fstrategy{
		FstrategyId: fstrategyId,
	}

	fstrategyModelBase := model_base.ModelBaseImpl(fstrategyObj)
	err := fstrategyModelBase.DeleModelsByCondition("fstrategy_id = ?", []interface{}{fstrategyObj.FstrategyId}...)

	if err != nil {
		logger.Logger.Error("%s fstrategy_id:%s err:%s", util.RunFuncName(), fstrategyObj.FstrategyId, err)
		logger.Logger.Print("%s fstrategy_id:%s err:%s", util.RunFuncName(), fstrategyObj.FstrategyId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleFStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//dele FstrategyVehicleItem
	fstrategyVehicle := &model.FstrategyVehicle{}
	fstrategyVehicleModelBase := model_base.ModelBaseImpl(fstrategyVehicle)
	err = fstrategyVehicleModelBase.DeleModelsByCondition("fstrategy_vehicle_id in (?)", fstrategyVehicleIdMapSlice)
	if err != nil {

	}
	//dele FstrategyVehicleItem
	fstrategyVehicleItem := &model.FstrategyVehicleItem{}
	fstrategyVehicleItemModelBase := model_base.ModelBaseImpl(fstrategyVehicleItem)
	err = fstrategyVehicleItemModelBase.DeleModelsByCondition("fstrategy_item_id in (?)", fstrategyItemIdMapSlice)
	if err != nil {
	}

	//软删除FstrategyItem
	fstrategyItem := &model.FstrategyItem{}
	err = fstrategyItem.SoftDeleModelImpl("fstrategy_item_id in (?)", fstrategyItemIdMapSlice)
	if err != nil {
	}

	//下发会话策略
	for k := range fVehicleIdMap {
		fStrategySetCmd := model.GetVehicleRecentFStrategy(k)
		strategyCmd := &emq_cmd.FStrategySetCmd{
			VehicleId: fStrategySetCmd.VehicleId,
			TaskType:  int(protobuf.Command_FLOWSTRATEGY_ADD),

			FstrategyId: fStrategySetCmd.FstrategyId,
			Type:        fStrategySetCmd.Type,
			HandleMode:  fStrategySetCmd.HandleMode,
			Enable:      true,
			GroupId:     "", //目前不实现
		}
		topic_publish_handler.GetPublishService().PutMsg2PublicChan(strategyCmd)
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqDeleFStrategySuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}

/****************************************StrategyVehicle********************************************************/
//
//func GetStrategyVehicle(c *gin.Context) {
//	strategyId := c.Param("strategy_id")
//	argsTrimsEmpty := util.RrgsTrimsEmpty(strategyId)
//	if argsTrimsEmpty {
//		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
//		c.JSON(http.StatusOK, ret)
//		logger.Logger.Error("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), strategyId)
//		logger.Logger.Print("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), strategyId)
//	}
//	strategyVehicleInfo := &model.StrategyVehicle{
//		StrategyId: strategyId,
//	}
//
//	modelBase := model_base.ModelBaseImpl(strategyVehicleInfo)
//
//	strategyVehicleInfos := []*model.StrategyVehicle{}
//	err:=modelBase.GetModelListByCondition(&strategyVehicleInfos,"strategy_id = ?",[]interface{}{strategyVehicleInfo.StrategyId}...)
//
//	if err != nil {
//		logger.Logger.Error("%s strategy_id:%s,err:%s", util.RunFuncName(), strategyId, err)
//		logger.Logger.Print("%s strategy_id:%s,err:%s", util.RunFuncName(), strategyId, err)
//		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrategyVehicleListFailMsg, "")
//		c.JSON(http.StatusOK, ret)
//		return
//	}
//
//	responseData := map[string]interface{}{
//		"strategy_vehicles": strategyVehicleInfos,
//	}
//
//	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetStrategyVehicleListSuccessMsg, responseData)
//	c.JSON(http.StatusOK, retObj)
//}
//
//
//
///****************************************StrategyVehicleResult********************************************************/
//
//func GetVehicleLearningResults(c *gin.Context) {
//	strategyVehicleId := c.Param("strategy_vehicle_id")
//	argsTrimsEmpty := util.RrgsTrimsEmpty(strategyVehicleId)
//	if argsTrimsEmpty {
//		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
//		c.JSON(http.StatusOK, ret)
//		logger.Logger.Error("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), strategyVehicleId)
//		logger.Logger.Print("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), strategyVehicleId)
//	}
//	vehicleLearnResultInfo := &model.StrategyVehicleLearningResult{
//		StrategyVehicleId:strategyVehicleId,
//	}
//
//	modelBase := model_base.ModelBaseImpl(vehicleLearnResultInfo)
//
//	strategyVehicleLearnResultInfos := []*model.StrategyVehicleLearningResult{}
//	err:=modelBase.GetModelListByCondition(&strategyVehicleLearnResultInfos,"strategy_vehicle_id = ?",[]interface{}{vehicleLearnResultInfo.StrategyVehicleId}...)
//
//	if err != nil {
//		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleLearnResultInfo.StrategyVehicleId, err)
//		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleLearnResultInfo.StrategyVehicleId, err)
//		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrategyVehicleResultListFailMsg, "")
//		c.JSON(http.StatusOK, ret)
//		return
//	}
//
//	responseData := map[string]interface{}{
//		"vehicle_results": strategyVehicleLearnResultInfos,
//	}
//
//	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetStrategyVehicleResultListSuccessMsg, responseData)
//	c.JSON(http.StatusOK, retObj)
//}
//

func GetVehicleFStrategyItem(c *gin.Context) {
	strategyId := c.Param("strategy_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(strategyId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), strategyId)
		logger.Logger.Print("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), strategyId)
	}

	results, _ := model.GetStrategyVehicleLearningResults("strategy_vehicles.strategy_id = ?", []interface{}{strategyId}...)

	responseData := map[string]interface{}{
		"strategy_vehicle_results": results,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetStrategyVehicleResultListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}
