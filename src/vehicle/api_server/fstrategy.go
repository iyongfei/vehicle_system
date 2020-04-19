package api_server

import (
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
	strategyId := c.Param("strategy_id")
	vehicleId := c.PostForm("vehicle_id")
	setTypeP := c.PostForm("type")
	handleModeP := c.PostForm("handle_mode")

	argsTrimsEmpty := util.RrgsTrimsEmpty(strategyId, setTypeP, handleModeP, vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
	}

	setType, _ := strconv.Atoi(setTypeP)
	handleMode, _ := strconv.Atoi(handleModeP)

	strategyVehicleLearningResultJoins, err := model.GetStrategyJoinVehicles(
		"strategy_vehicles.strategy_id = ? and strategy_vehicles.vehicle_id = ?",
		[]interface{}{strategyId, vehicleId}...)

	if err != nil || strategyVehicleLearningResultJoins.StrategyId == "" || strategyVehicleLearningResultJoins.VehicleId == "" {
		logger.Logger.Error("%s strategyId:%s,recordNotFounder", util.RunFuncName(), strategyId)
		logger.Logger.Print("%s strategyId:%s,recordNotFounder", util.RunFuncName(), strategyId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//查询是否存在
	strategyInfo := &model.Strategy{
		StrategyId: strategyVehicleLearningResultJoins.StrategyId,
	}
	modelBase := model_base.ModelBaseImpl(strategyInfo)
	err, recordNotFound := modelBase.GetModelByCondition("strategy_id = ?", []interface{}{strategyInfo.StrategyId}...)

	if err != nil {
		logger.Logger.Error("%s strategy_id:%s,err:%s", util.RunFuncName(), strategyId, err)
		logger.Logger.Print("%s strategy_id:%s,err:%s", util.RunFuncName(), strategyId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s strategyId:%s,recordNotFound", util.RunFuncName(), strategyId)
		logger.Logger.Print("%s strategyId:%s,recordNotFound", util.RunFuncName(), strategyId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	strategyInfo.HandleMode = uint8(handleMode)
	strategyInfo.Type = uint8(setType)

	attrs := map[string]interface{}{
		"handle_mode": strategyInfo.HandleMode,
		"type":        strategyInfo.Type,
	}
	if err := modelBase.UpdateModelsByCondition(attrs, "strategy_id = ?",
		[]interface{}{strategyInfo.StrategyId}...); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//更新
	strategyCmd := &emq_cmd.StrategySetCmd{
		VehicleId: vehicleId,
		TaskType:  int(protobuf.Command_STRATEGY_ADD),

		StrategyId: strategyId,
		Type:       setType,
		HandleMode: handleMode,
		Enable:     true,
		GroupId:    "", //目前不实现
	}
	topic_publish_handler.GetPublishService().PutMsg2PublicChan(strategyCmd)

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

	strategys := []*model.Strategy{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.Strategy{})

	err := modelBase.GetModelPaginationByCondition(pageIndex, pageSize,
		&total, &strategys, "",
		[]interface{}{}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrategyListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"strategys":  strategys,
		"totalCount": total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetStrategyListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func GetFStrategy(c *gin.Context) {
	strategyId := c.Param("strategy_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(strategyId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), strategyId)
		logger.Logger.Print("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), strategyId)
	}
	strategyInfo := &model.Strategy{
		StrategyId: strategyId,
	}

	modelBase := model_base.ModelBaseImpl(strategyInfo)

	err, recordNotFound := modelBase.GetModelByCondition("strategy_id = ?", []interface{}{strategyInfo.StrategyId}...)

	if err != nil {
		logger.Logger.Error("%s strategy_id:%s,err:%s", util.RunFuncName(), strategyId, err)
		logger.Logger.Print("%s strategy_id:%s,err:%s", util.RunFuncName(), strategyId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s strategy_id:%s,recordNotFound", util.RunFuncName(), strategyId)
		logger.Logger.Print("%s strategy_id:%s,recordNotFound", util.RunFuncName(), strategyId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	responseData := map[string]interface{}{
		"strategy": strategyInfo,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func AddFStrategy(c *gin.Context) {
	vehicleIdsP := c.PostForm("vehicle_ids")
	sType := c.PostForm("type")
	handleMode := c.PostForm("handle_mode")
	dips := c.PostForm("dips")
	dstPorts := c.PostForm("dst_ports")

	sTypeValid := util.IsEleExistInSlice(sType, []interface{}{
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOTYPEDEFAULT)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOWHITEMODE)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOBLACKMODE))})

	handleModeValid := util.IsEleExistInSlice(sType, []interface{}{
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_MODEDEFAULT)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_PREVENTWARNING)),
		strconv.Itoa(int(protobuf.FlowStrategyAddParam_WARNING))})

	logger.Logger.Print("%s vehicleIdsP:%s,sType:%s,handleMode:%s,dip:%s,dstPort:%s",
		util.RunFuncName(), vehicleIdsP, sType, handleMode, dips, dstPorts)
	logger.Logger.Info("%s vehicleIdsP:%s,sType:%s,handleMode:%s,dip:%s,dstPort:%s",
		util.RunFuncName(), vehicleIdsP, sType, handleMode, dips, dstPorts)

	argsTrimsEmpty := util.RrgsTrimsEmpty(sType, handleMode, dips, dstPorts, vehicleIdsP)
	if argsTrimsEmpty ||
		!sTypeValid ||
		!handleModeValid {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Print("%s vehicleIdsP:%s,sType:%s,handleMode:%s,dip:%s,dstPort:%s",
			util.RunFuncName(), vehicleIdsP, sType, handleMode, dips, dstPorts)

		logger.Logger.Info("%s vehicleIdsP:%s,sType:%s,handleMode:%s,dip:%s,dstPort:%s",
			util.RunFuncName(), vehicleIdsP, sType, handleMode, dips, dstPorts)
		return
	}

	vStype, _ := strconv.Atoi(sType)
	vHandleMode, _ := strconv.Atoi(handleMode)

	fdipSlice := []string{}
	//筛选dip
	dipSlice := strings.Split(dips, ",")
	for _, dip := range dipSlice {
		destIpValid := util.IpFormat(dip)
		if destIpValid && !util.IsExistInSlice(dip,fdipSlice) {
			fdipSlice = append(fdipSlice, dip)
		}
	}

	fdportSlice := []uint32{}
	//筛选dstPort
	dstPortSlice := strings.Split(dstPorts, ",")

	for _, dport := range dstPortSlice {
		destPortValid := util.VerifyIpPort(dport)
		dpInt,_:=strconv.Atoi(dport)
		if destPortValid && !util.IsExistInSlice(uint32(dpInt),fdportSlice)  {
			fdportSlice = append(fdportSlice, uint32(dpInt))
		}
	}

	//找出合法的vehicle
	vehicleIdSlice := strings.Split(vehicleIdsP, ",")
	var vehicleIds []string
	_ = mysql.QueryPluckByModelWhere(&model.VehicleInfo{}, "vehicle_id", &vehicleIds, "vehicle_id in (?)", vehicleIdSlice)


	//fstrategy_items table
	fstrategyItems := map[string][]string{}
	for _,vehicleId := range vehicleIds{
		var vehicleFstrategyItems []string
		for _,dip := range fdipSlice{
			for _,dport:=range fdportSlice{
				fstrategyItem := &model.FstrategyItem{
					FstrategyItemId: 	util.RandomString(32),
					VehicleId:       	vehicleId,
					DstIp: 				dip,
					DstPort:     		dport,
				}
				modelBase := model_base.ModelBaseImpl(fstrategyItem)

				err,fstrategyItemrecordNotFound :=modelBase.GetModelByCondition(
					"vehicle_id = ? and dst_ip = ? and dst_port = ?",
					[]interface{}{fstrategyItem.VehicleId,fstrategyItem.DstIp,fstrategyItem.DstPort}...)
				if err!=nil{
					continue
				}

				if fstrategyItemrecordNotFound{
					if err := modelBase.InsertModel(); err != nil {
						continue
					}
				}
				if !util.IsExistInSlice(fstrategyItem.FstrategyItemId,vehicleFstrategyItems){
					vehicleFstrategyItems = append(vehicleFstrategyItems,fstrategyItem.FstrategyItemId)
				}
			}
		}
		fstrategyItems[vehicleId] = vehicleFstrategyItems
	}

	//fstrategy table
	fstrategy := &model.Fstrategy{
		FstrategyId:util.RandomString(32),
		Type:uint8(vStype),
		HandleMode:uint8(vHandleMode),
		Enable:true,
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
			VehicleId:         vehicleId,
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
				FstrategyItemId:  item,
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
			Type:       vStype,
			HandleMode: vHandleMode,
			Enable:     true,
			GroupId:    "", //目前不实现
		}
		topic_publish_handler.GetPublishService().PutMsg2PublicChan(fstrategyCmd)
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddStrategySuccessMsg, fstrategy)
	c.JSON(http.StatusOK, retObj)
}

func DeleFStrategy(c *gin.Context) {
	strategyId := c.Param("strategy_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(strategyId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty assetId:%s argsTrimsEmpty", util.RunFuncName(), strategyId)
		logger.Logger.Print("%s argsTrimsEmpty assetId:%s argsTrimsEmpty", util.RunFuncName(), strategyId)
		return
	}

	strategyObj := &model.Strategy{
		StrategyId: strategyId,
	}

	modelBase := model_base.ModelBaseImpl(strategyObj)
	err, recordNotFound := modelBase.GetModelByCondition("strategy_id = ?", []interface{}{strategyObj.StrategyId}...)

	if err != nil {
		logger.Logger.Error("%s strategy_id:%s err:%s", util.RunFuncName(), strategyObj.StrategyId, err)
		logger.Logger.Print("%s strategy_id:%s err:%s", util.RunFuncName(), strategyObj.StrategyId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if recordNotFound {
		logger.Logger.Error("%s asset_id:%s,record not exist", util.RunFuncName(), strategyObj.StrategyId)
		logger.Logger.Print("%s asset_id:%s,record not exist", util.RunFuncName(), strategyObj.StrategyId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err := modelBase.DeleModelsByCondition("strategy_id = ?",
		[]interface{}{strategyObj.StrategyId}...); err != nil {
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqDeleStrategySuccessMsg, "")
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