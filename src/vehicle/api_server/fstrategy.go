package api_server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"vehicle_system/src/vehicle/csv"
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
	diports := c.PostForm("dip_ports")

	logger.Logger.Print("%s vehicle_id:%s,fstrategy_id:%s,diports:%v", util.RunFuncName(), vehicleId, fstrategyId, diports)
	logger.Logger.Info("%s vehicle_id:%s,fstrategy_id:%s,diports:%v", util.RunFuncName(), vehicleId, fstrategyId, diports)

	//setTypeP := c.PostForm("type")
	//handleModeP := c.PostForm("handle_mode")
	//sTypeValid := util.IsEleExistInSlice(setTypeP, []interface{}{
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOTYPEDEFAULT)),
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOWHITEMODE)),
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOBLACKMODE))})
	//handleModeValid := util.IsEleExistInSlice(handleModeP, []interface{}{
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_MODEDEFAULT)),
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_PREVENTWARNING)),
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_WARNING))})
	argsTrimsEmpty := util.RrgsTrimsEmpty(fstrategyId, vehicleId, diports)
	diportsMap := map[string][]uint32{}
	err := json.Unmarshal([]byte(diports), &diportsMap)

	if argsTrimsEmpty || err != nil {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
		return
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

	finalDiportsMap := map[string][]uint32{}
	for dip, ports := range diportsMap {
		destIpValid := util.IpFormat(dip)
		if destIpValid {
			for _, port := range ports {
				portStr := strconv.Itoa(int(port))
				if util.VerifyIpPort(portStr) {
					finalDiportsMap[dip] = append(finalDiportsMap[dip], port)
				}
			}
		}
	}

	logger.Logger.Print("%s vehicle_id:%s finalDiportsMap:%+v", util.RunFuncName(), vehicleId, finalDiportsMap)
	logger.Logger.Info("%s vehicle_id:%s finalDiportsMap:%+v", util.RunFuncName(), vehicleId, finalDiportsMap)

	//FstrategyVehicle,Fstrategy不更改

	//遍历FstrategyVehicleItem表
	fstrategyItems := map[string][]string{}
	var vehicleFstrategyItems []string
	for dip, ports := range finalDiportsMap {
		for _, dport := range ports {
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

	logger.Logger.Print("%s vehicle_id:%s fstrategyItems:%+v", util.RunFuncName(), vehicleId, fstrategyItems)
	logger.Logger.Info("%s vehicle_id:%s fstrategyItems:%+v", util.RunFuncName(), vehicleId, fstrategyItems)

	//找到FstrategyItemId(FstrategyVehicleItem表中)
	//在FstrategyItemId(FstrategyItem表中)不存在的值

	var fstrategyVehicleItemIds []string
	_ = mysql.QueryPluckByModelWhere(&model.FstrategyVehicleItem{}, "fstrategy_item_id", &fstrategyVehicleItemIds,
		"fstrategy_vehicle_id = ?", vehicleFStrategy.FstrategyVehicleId)

	logger.Logger.Print("%s fstrategyItemIds:%+v", util.RunFuncName(), fstrategyVehicleItemIds)
	logger.Logger.Info("%s fstrategyItemIds:%+v", util.RunFuncName(), fstrategyVehicleItemIds)

	//如果没有在里面，就是被删除的，需要改delete标志位
	newFstrategyItemIds := fstrategyItems[vehicleId]
	var needDeleFstrategyItemIds []string
	for _, fstrategyItemId := range fstrategyVehicleItemIds {
		if !util.IsExistInSlice(fstrategyItemId, newFstrategyItemIds) {
			needDeleFstrategyItemIds = append(needDeleFstrategyItemIds, fstrategyItemId)
		}
	}
	logger.Logger.Print("%s needDeleFstrategyItemIds:%+v", util.RunFuncName(), needDeleFstrategyItemIds)
	logger.Logger.Info("%s needDeleFstrategyItemIds:%+v", util.RunFuncName(), needDeleFstrategyItemIds)

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

	//插入csv
	csvModel := csv.NewCsvWriter(vehicleFStrategy.FstrategyId, csv.FileTruncate)
	fCsvHeader := csv.CreateCsvFstrategyHeader()
	fCsvBody := csv.CreateCsvFstrategyBody(vehicleId, vehicleFStrategy.FstrategyId, finalDiportsMap)
	csvModel.SetCsvWritData(fCsvHeader, fCsvBody)

	//更新
	fstrategyCmd := &emq_cmd.FStrategySetCmd{
		VehicleId: vehicleId,
		TaskType:  int(protobuf.Command_FLOWSTRATEGY_ADD),

		FstrategyId: fstrategyId,
		Type:        int(protobuf.FlowStrategyAddParam_FLWOWHITEMODE),
		HandleMode:  int(protobuf.FlowStrategyAddParam_WARNING),

		Enable:  true,
		GroupId: "", //目前不实现
	}
	topic_publish_handler.GetPublishService().PutMsg2PublicChan(fstrategyCmd)

	vehicleSingleFlowStrategyItemsReult := model.VehicleSingleFlowStrategyItemsReult{
		FstrategyId:              vehicleFStrategy.FstrategyId,
		Type:                     vehicleFStrategy.Type,
		HandleMode:               vehicleFStrategy.HandleMode,
		Enable:                   vehicleFStrategy.Enable,
		VehicleId:                vehicleFStrategy.VehicleId,
		CsvPath:                  vehicleFStrategy.CsvPath,
		VehicleFStrategyItemsMap: finalDiportsMap,
	}

	responseData := map[string]interface{}{
		"put_fstrategy": vehicleSingleFlowStrategyItemsReult,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateStrategySuccessMsg, responseData)
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
		&total, &fstrategys, "", "",
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

func AddFStrategy(c *gin.Context) {
	vehicleId := c.PostForm("vehicle_id")
	diports := c.PostForm("dip_ports")

	//默认白名单
	//sType := c.PostForm("type")
	//默认告警
	//handleMode := c.PostForm("handle_mode")
	//此版本不实现
	//sTypeValid := util.IsEleExistInSlice(sType, []interface{}{
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOTYPEDEFAULT)),
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOWHITEMODE)),
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_FLWOBLACKMODE))})
	//
	//handleModeValid := util.IsEleExistInSlice(handleMode, []interface{}{
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_MODEDEFAULT)),
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_PREVENTWARNING)),
	//	strconv.Itoa(int(protobuf.FlowStrategyAddParam_WARNING))})

	logger.Logger.Print("%s vehicle_ids:%s,diports:%v",
		util.RunFuncName(), vehicleId, diports)
	logger.Logger.Info("%s vehicle_ids:%s,diports:%v",
		util.RunFuncName(), vehicleId, diports)

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, diports)

	diportsMap := map[string][]uint32{}
	err := json.Unmarshal([]byte(diports), &diportsMap)
	if argsTrimsEmpty || err != nil {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s vehicle_ids:%s,diports:%v", util.RunFuncName(), vehicleId, diports)
		logger.Logger.Error("%s vehicle_ids:%s,diports:%v", util.RunFuncName(), vehicleId, diports)
		return
	}
	logger.Logger.Info("%s vehicleId:%s,diports:%s", util.RunFuncName(), vehicleId, diports)
	logger.Logger.Print("%s vehicleId:%s,diports:%s", util.RunFuncName(), vehicleId, diports)
	//找出合法的vehicle
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}
	vehicleInfoModelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := vehicleInfoModelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)
	if err != nil || recordNotFound {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s vehicle_id:%s recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Error("%s vehicle_id:%s recordNotFound", util.RunFuncName(), vehicleId)
		return
	}

	finalDiportsMap := map[string][]uint32{}
	for dip, ports := range diportsMap {
		destIpValid := util.IpFormat(dip)
		if destIpValid {
			for _, port := range ports {
				portStr := strconv.Itoa(int(port))
				if util.VerifyIpPort(portStr) {
					finalDiportsMap[dip] = append(finalDiportsMap[dip], port)
				}
			}
		}
	}

	//fstrategy_items table
	fstrategyItems := map[string][]string{}
	var vehicleFstrategyItems []string

	for dip, ports := range finalDiportsMap {
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

	logger.Logger.Print("%s vehicle_id:%s fstrategyItems:%+v", util.RunFuncName(), vehicleId, fstrategyItems)
	logger.Logger.Info("%s vehicle_id:%s fstrategyItems:%+v", util.RunFuncName(), vehicleId, fstrategyItems)

	//fstrategy table
	fstrategy := &model.Fstrategy{
		FstrategyId: util.RandomString(32),
		Type:        uint8(protobuf.FlowStrategyAddParam_FLWOWHITEMODE),
		HandleMode:  uint8(protobuf.FlowStrategyAddParam_WARNING),
		Enable:      true,
	}

	fstrategyModelBase := model_base.ModelBaseImpl(fstrategy)
	if err := fstrategyModelBase.InsertModel(); err != nil {
		//todo
	}

	//FstrategyVehicle table
	fstrategyVehicle := &model.FstrategyVehicle{
		FstrategyVehicleId: util.RandomString(32),
		FstrategyId:        fstrategy.FstrategyId,
		VehicleId:          vehicleId,
	}
	fstrategyVehicleModelBase := model_base.ModelBaseImpl(fstrategyVehicle)
	if err := fstrategyVehicleModelBase.InsertModel(); err != nil {
		logger.Logger.Print("%s vehicle_id:%s insert FstrategyVehicle err:%+v", util.RunFuncName(), vehicleId, err)
		logger.Logger.Info("%s vehicle_id:%s insert FstrategyVehicle err:%+v", util.RunFuncName(), vehicleId, err)

	}

	vehicleFsItems := fstrategyItems[vehicleId]
	//fstrategy_vehicle_items table
	for _, item := range vehicleFsItems {
		fstrategyVehicleItem := &model.FstrategyVehicleItem{
			FstrategyVehicleId: fstrategyVehicle.FstrategyVehicleId,
			FstrategyItemId:    item,
		}

		fstrategyVehicleItemModelBase := model_base.ModelBaseImpl(fstrategyVehicleItem)

		if err := fstrategyVehicleItemModelBase.InsertModel(); err != nil {
			logger.Logger.Print("%s vehicle_id:%s insert fstrategyVehicleItem err:%+v", util.RunFuncName(), vehicleId, err)
			logger.Logger.Info("%s vehicle_id:%s insert fstrategyVehicleItem err:%+v", util.RunFuncName(), vehicleId, err)

			continue
		}
	}
	//插入csv
	csvModel := csv.NewCsvWriter(fstrategy.FstrategyId, csv.FileAppend)
	fCsvHeader := csv.CreateCsvFstrategyHeader()
	fCsvBody := csv.CreateCsvFstrategyBody(vehicleId, fstrategy.FstrategyId, finalDiportsMap)
	csvModel.SetCsvWritData(fCsvHeader, fCsvBody)

	attrs := map[string]interface{}{
		"csv_path": csvModel.CsvFilePath,
	}
	if err := fstrategyModelBase.UpdateModelsByCondition(attrs, "fstrategy_id = ?", fstrategy.FstrategyId); err != nil {
		logger.Logger.Print("%s insert fstrategy csv_path:%s, err:%+v", util.RunFuncName(), csvModel.CsvFilePath, err)
		logger.Logger.Info("%s insert fstrategy csv_path:%s, err:%+v", util.RunFuncName(), csvModel.CsvFilePath, err)
	} else {
		fstrategy.CsvPath = csvModel.CsvFilePath
	}

	//下发策略
	fstrategyCmd := &emq_cmd.FStrategySetCmd{
		VehicleId: vehicleId,
		TaskType:  int(protobuf.Command_FLOWSTRATEGY_ADD),

		FstrategyId: fstrategy.FstrategyId,
		Type:        int(protobuf.FlowStrategyAddParam_FLWOWHITEMODE),
		HandleMode:  int(protobuf.FlowStrategyAddParam_WARNING),
		Enable:      true,
		GroupId:     "", //目前不实现
	}
	topic_publish_handler.GetPublishService().PutMsg2PublicChan(fstrategyCmd)

	vehicleSingleFlowStrategyItemsReult := model.VehicleSingleFlowStrategyItemsReult{
		FstrategyId:              fstrategy.FstrategyId,
		Type:                     fstrategy.Type,
		HandleMode:               fstrategy.HandleMode,
		Enable:                   fstrategy.Enable,
		VehicleId:                vehicleInfo.VehicleId,
		CsvPath:                  fstrategy.CsvPath,
		VehicleFStrategyItemsMap: finalDiportsMap,
	}

	responseData := map[string]interface{}{
		"fstrategy": vehicleSingleFlowStrategyItemsReult,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddFStrategySuccessMsg, responseData)
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
	logger.Logger.Info("%s fstrategyId:%s", util.RunFuncName(), fstrategyId)
	logger.Logger.Print("%s fstrategyId:%s", util.RunFuncName(), fstrategyId)

	fStrategyObj := &model.Fstrategy{
		FstrategyId: fstrategyId,
	}

	modelBase := model_base.ModelBaseImpl(fStrategyObj)
	err, frecordNotFound := modelBase.GetModelByCondition("fstrategy_id = ?", []interface{}{fStrategyObj.FstrategyId}...)

	if err != nil {
		logger.Logger.Error("%s fstrategy_id:%s err:%s", util.RunFuncName(), fStrategyObj.FstrategyId, err)
		logger.Logger.Print("%s fstrategy_id:%s err:%s", util.RunFuncName(), fStrategyObj.FstrategyId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleFStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if frecordNotFound {
		logger.Logger.Error("%s asset_id:%s,record not exist", util.RunFuncName(), fStrategyObj.FstrategyId)
		logger.Logger.Print("%s asset_id:%s,record not exist", util.RunFuncName(), fStrategyObj.FstrategyId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
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
		fVehicleIdMapSlice = append(fVehicleIdMapSlice, k)
	}
	for k := range fstrategyVehicleIdMap {
		fstrategyVehicleIdMapSlice = append(fstrategyVehicleIdMapSlice, k)
	}
	for k := range fstrategyItemIdMap {
		fstrategyItemIdMapSlice = append(fstrategyItemIdMapSlice, k)
	}

	//dele Fstrategy
	fstrategyObj := &model.Fstrategy{
		FstrategyId: fstrategyId,
	}

	fstrategyModelBase := model_base.ModelBaseImpl(fstrategyObj)

	err, recordNotFound := fstrategyModelBase.GetModelByCondition("fstrategy_id = ?", fstrategyObj.FstrategyId)

	if !recordNotFound {

		err := fstrategyModelBase.DeleModelsByCondition("fstrategy_id = ?", []interface{}{fstrategyObj.FstrategyId}...)

		if err != nil {
			logger.Logger.Error("%s fstrategy_id:%s err:%s", util.RunFuncName(), fstrategyObj.FstrategyId, err)
			logger.Logger.Print("%s fstrategy_id:%s err:%s", util.RunFuncName(), fstrategyObj.FstrategyId, err)
			ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleFStrategyFailMsg, "")
			c.JSON(http.StatusOK, ret)
			return
		}
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
	err = fstrategyVehicleItemModelBase.DeleModelsByCondition(
		"fstrategy_item_id in (?) and fstrategy_vehicle_id in (?)",
		fstrategyItemIdMapSlice, fstrategyVehicleIdMapSlice)
	if err != nil {
	}

	//软删除FstrategyItem
	fstrategyItem := &model.FstrategyItem{}
	err = fstrategyItem.SoftDeleModelImpl("fstrategy_item_id in (?)", fstrategyItemIdMapSlice)
	if err != nil {
	}
	//删除scv
	//http://192.168.100.2:7001/fstrategy_csv/N5gqNSN0lpV30gKJOfBkYvGudNUfj1V5.csv
	csvPath := fstrategyObj.CsvPath
	fStrategyCsvFolderIndex := strings.Index(csvPath, csv.FStrategyCsvFolder)

	var csvFileName string
	if fStrategyCsvFolderIndex != -1 {
		csvFileName = csvPath[fStrategyCsvFolderIndex:]
	}
	isExist := csv.IsExists(csvFileName)
	if isExist {
		csvRemoveErr := os.Remove(csvFileName)
		if csvRemoveErr != nil {
			logger.Logger.Error("%s remove csvFile:%s,err:%s", util.RunFuncName(), csvFileName, err)
			logger.Logger.Print("%s remove csvFile:%s,err:%s", util.RunFuncName(), csvFileName, err)
		}
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

/**
查看所有的策略id
*/
func GetAllFstrategys(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s", util.RunFuncName(), vehicleId)
		return
	}
	logger.Logger.Print("%s vehicle_id:%s", util.RunFuncName(), vehicleId)
	logger.Logger.Info("%s vehicle_id:%s", util.RunFuncName(), vehicleId)

	//查询终端id是否存在
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//查询所有的策略
	//查看该vehicle是否存在
	vehicleFStrategys, err := model.GetFStrategyVehicles(
		"fstrategy_vehicles.vehicle_id = ?", []interface{}{vehicleId}...)

	var fstrategys []interface{}

	for _, v := range vehicleFStrategys {
		fstrategyId := v.FstrategyId
		createdAt := v.CreatedAt
		fstrategyIds := map[string]interface{}{}

		fstrategyIds["CreatedAt"] = createdAt
		fstrategyIds["FstrategyId"] = fstrategyId
		fstrategys = append(fstrategys, fstrategyIds)
	}

	responseData := map[string]interface{}{
		"Vehicle_Id":     vehicleId,
		"All_Fstrategys": fstrategys,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}

/**
查看所有的策略id
*/
func GetPartFstrategyIds(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	fstrategyIds := c.Query("fstrategy_ids")
	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, fstrategyIds)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s,fstrategyIds%s", util.RunFuncName(), vehicleId, fstrategyIds)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s,fstrategyIds%s", util.RunFuncName(), vehicleId, fstrategyIds)
		return
	}
	logger.Logger.Print("%s vehicle_id:%s", util.RunFuncName(), vehicleId)
	logger.Logger.Info("%s vehicle_id:%s", util.RunFuncName(), vehicleId)

	//过滤非法的fstrategyIds
	fstrategyIdsRrgsTrim := util.RrgsTrim(fstrategyIds)
	fstrategyIdSlice := strings.Split(fstrategyIdsRrgsTrim, ",")

	fstrategyVehicles, err := model.GetFStrategyVehicles("fstrategy_vehicles.vehicle_id = ? and fstrategy_vehicles.fstrategy_id in (?)",
		[]interface{}{vehicleId, fstrategyIdSlice}...)

	if err != nil {

	}
	var fstrategyVehicleIds []string
	for _, fstrategyVehicle := range fstrategyVehicles {

		fVehicleId := fstrategyVehicle.FstrategyVehicleId
		exist := util.IsExistInSlice(fVehicleId, fstrategyVehicleIds)
		if !exist {
			fstrategyVehicleIds = append(fstrategyVehicleIds, fVehicleId)
		}
	}

	logger.Logger.Print("%s fstrategyVehicleIds:%+v", util.RunFuncName(), fstrategyVehicleIds)
	logger.Logger.Info("%s fstrategyVehicleIds:%+v", util.RunFuncName(), fstrategyVehicleIds)

	//fstrategyVehicleIds
	vehicleFStrategyItems, err := model.GetVehicleFStrategyItems(
		"fstrategy_vehicle_items.fstrategy_vehicle_id in (?) and fstrategy_items.deleted_at is null",
		[]interface{}{fstrategyVehicleIds}...)

	vehicleFStrategyItemsMap := map[string]map[string][]uint32{}

	for _, vehicleFStrategyItem := range vehicleFStrategyItems {
		_, ok := vehicleFStrategyItemsMap[vehicleFStrategyItem.FstrategyVehicleId]
		if !ok {
			vehicleFStrategyItemsMap[vehicleFStrategyItem.FstrategyVehicleId] = map[string][]uint32{}

			_, oker := vehicleFStrategyItemsMap[vehicleFStrategyItem.FstrategyVehicleId][vehicleFStrategyItem.DstIp]
			if !oker {
				vehicleFStrategyItemsMap[vehicleFStrategyItem.FstrategyVehicleId][vehicleFStrategyItem.DstIp] = []uint32{vehicleFStrategyItem.DstPort}
			}
		} else {
			vehicleFStrategyItemsMap[vehicleFStrategyItem.FstrategyVehicleId][vehicleFStrategyItem.DstIp] =
				append(vehicleFStrategyItemsMap[vehicleFStrategyItem.FstrategyVehicleId][vehicleFStrategyItem.DstIp], vehicleFStrategyItem.DstPort)
		}
	}

	var list []model.VehicleSingleFlowStrategyItemsReult
	for _, vehicleFStrategy := range fstrategyVehicles {
		vehicleSingleFlowStrategyItemsReult := model.VehicleSingleFlowStrategyItemsReult{
			FstrategyId:              vehicleFStrategy.FstrategyId,
			Type:                     vehicleFStrategy.Type,
			HandleMode:               vehicleFStrategy.HandleMode,
			Enable:                   vehicleFStrategy.Enable,
			VehicleId:                vehicleFStrategy.VehicleId,
			CsvPath:                  vehicleFStrategy.CsvPath,
			VehicleFStrategyItemsMap: vehicleFStrategyItemsMap[vehicleFStrategy.FstrategyVehicleId],
		}
		list = append(list, vehicleSingleFlowStrategyItemsReult)
	}

	responseData := map[string]interface{}{
		"fstrategys": list,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

/**
获取所有的策略
*/
func GetPaginationFstrategys(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")
	startTimeP := c.Query("start_time")
	endTimeP := c.Query("end_time")

	logger.Logger.Info("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)
	logger.Logger.Print("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		return
	}

	fpageSize, _ := strconv.Atoi(pageSizeP)
	fpageIndex, _ := strconv.Atoi(pageIndexP)

	startTime, _ := strconv.Atoi(startTimeP)
	endTime, _ := strconv.Atoi(endTimeP)

	var fStartTime time.Time
	var fEndTime time.Time

	//默认20
	defaultPageSize := 20
	if fpageSize == 0 {
		fpageSize = defaultPageSize
	}
	//默认第一页
	defaultPageIndex := 1
	if fpageIndex == 0 {
		fpageIndex = defaultPageIndex
	}
	//默认2天前
	defaultStartTime := util.GetFewDayAgo(2) //2
	if startTime == 0 {
		fStartTime = defaultStartTime
	} else {
		fStartTime = util.StampUnix2Time(int64(startTime))
	}

	//默认当前时间
	defaultEndTime := time.Now()
	if endTime == 0 {
		fEndTime = defaultEndTime
	} else {
		fEndTime = util.StampUnix2Time(int64(endTime))
	}

	logger.Logger.Info("%s frequest params vehicle_id:%s,fpageSize:%d,fpageIndex:%d,fStartTime%d,fEndTime%d",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)
	logger.Logger.Print("%s frequest params vehicle_id:%s,fpageSize:%d,fpageIndex:%d,fStartTime%d,fEndTime%d",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)

	//查询终端id是否存在
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	var totalCount int

	//终端-策略
	vehicleFStrategys, err := model.GetPaginFStrategyVehicles(fpageIndex, fpageSize, &totalCount,
		"fstrategy_vehicles.vehicle_id = ? and fstrategies.created_at BETWEEN ? AND ?", []interface{}{vehicleId, fStartTime, fEndTime}...)

	if len(vehicleFStrategys) == 0 {
		logger.Logger.Error("%s vehicle_id:%s,vehicleFStrategys null", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,vehicleFStrategys null", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,vehicleFStrategys err:%+v", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,vehicleFStrategys err:%+v", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//获取fstrategyVehicleIds
	var fstrategyVehicleIds []string
	for _, fstrategyVehicle := range vehicleFStrategys {

		fVehicleId := fstrategyVehicle.FstrategyVehicleId
		exist := util.IsExistInSlice(fVehicleId, fstrategyVehicleIds)
		if !exist {
			fstrategyVehicleIds = append(fstrategyVehicleIds, fVehicleId)
		}
	}
	logger.Logger.Print("%s vehicle_id:%s,fstrategyVehicleIds:%+v", util.RunFuncName(), vehicleId, fstrategyVehicleIds)
	logger.Logger.Info("%s vehicle_id:%s,fstrategyVehicleIds:%+v", util.RunFuncName(), vehicleId, fstrategyVehicleIds)

	//fstrategyVehicleIds
	vehicleFStrategyItems, err := model.GetVehicleFStrategyItems(
		"fstrategy_vehicle_items.fstrategy_vehicle_id in (?) and fstrategy_items.deleted_at is null",
		[]interface{}{fstrategyVehicleIds}...)

	var list []model.VehicleSingleFlowStrategyItemsReult
	for _, vehicleFStrategy := range vehicleFStrategys {
		fVid := vehicleFStrategy.FstrategyVehicleId
		//每个策略的ip:port列表
		vehicleFStrategyItemsMap := map[string][]uint32{}
		for _, vehicleFStrategyItem := range vehicleFStrategyItems {
			fVehicleId := vehicleFStrategyItem.FstrategyVehicleId
			if fVehicleId == fVid {
				dip := util.RrgsTrim(vehicleFStrategyItem.DstIp)
				dPort := vehicleFStrategyItem.DstPort
				if _, ok := vehicleFStrategyItemsMap[dip]; ok {
					if !util.IsExistInSlice(dPort, vehicleFStrategyItemsMap[dip]) {
						vehicleFStrategyItemsMap[dip] = append(vehicleFStrategyItemsMap[dip], dPort)
					}
				} else {

					vehicleFStrategyItemsMap[dip] = []uint32{dPort}
				}

			}

		}
		vehicleSingleFlowStrategyItemsReult := model.VehicleSingleFlowStrategyItemsReult{
			FstrategyId:              vehicleFStrategy.FstrategyId,
			Type:                     vehicleFStrategy.Type,
			HandleMode:               vehicleFStrategy.HandleMode,
			Enable:                   vehicleFStrategy.Enable,
			VehicleId:                vehicleFStrategy.VehicleId,
			CsvPath:                  vehicleFStrategy.CsvPath,
			VehicleFStrategyItemsMap: vehicleFStrategyItemsMap,
		}
		list = append(list, vehicleSingleFlowStrategyItemsReult)
	}

	responseData := map[string]interface{}{
		"fstrategys": list,
		"totalCount": totalCount,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

/**
查询目前正在执行的策略
*/
func GetActiveFstrategy(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s", util.RunFuncName(), vehicleId)
		return
	}
	logger.Logger.Print("%s vehicle_id:%s", util.RunFuncName(), vehicleId)
	logger.Logger.Info("%s vehicle_id:%s", util.RunFuncName(), vehicleId)

	//查询终端id是否存在
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//查询策略
	recentFStrategy := model.GetVehicleRecentFStrategy(vehicleId)

	strategyCmd := &emq_cmd.FStrategySetCmd{
		VehicleId: recentFStrategy.VehicleId,
		TaskType:  int(protobuf.Command_FLOWSTRATEGY_ADD),

		FstrategyId: recentFStrategy.FstrategyId,
		Type:        recentFStrategy.Type,
		HandleMode:  recentFStrategy.HandleMode,
		Enable:      true,
		GroupId:     "", //目前不实现
	}

	_, dipPortMap := emq_cmd.FetchDipPortList(strategyCmd)

	logger.Logger.Print("%s dipPortMap:%+v", util.RunFuncName(), dipPortMap)
	logger.Logger.Info("%s dipPortMap:%+v", util.RunFuncName(), dipPortMap)

	vehicleSingleFlowStrategyItemsReult := model.VehicleSingleFlowStrategyItemsReult{
		FstrategyId:              recentFStrategy.FstrategyId,
		Type:                     uint8(recentFStrategy.Type),
		HandleMode:               uint8(recentFStrategy.HandleMode),
		Enable:                   recentFStrategy.Enable,
		VehicleId:                recentFStrategy.VehicleId,
		VehicleFStrategyItemsMap: dipPortMap,
	}
	responseData := map[string]interface{}{
		"active_fstrategy": vehicleSingleFlowStrategyItemsReult,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

// @Summary GetFStrategy
// @Description GetFStrategy
// @Produce json
// @Accept multipart/form-data
// @Param  vehicle_id query string true "vehicle_id"
// @Param fstrategy_id path string true "fstrategy_id"
// @Success 200 {object} model.VehicleSingleFlowStrategyItemsReult
// @Failure 400 {object} response.Response
// @Router /api/v1/fstrategys/{fstrategy_id} [get]
func GetFStrategy(c *gin.Context) {
	fstrategyId := c.Param("fstrategy_id")
	vehicleId := c.Query("vehicle_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(fstrategyId, vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty fstrategyId:%s", util.RunFuncName(), fstrategyId)
		logger.Logger.Print("%s argsTrimsEmpty fstrategyId:%s", util.RunFuncName(), fstrategyId)
		return
	}

	logger.Logger.Print("%s vehicle_id:%s,fstrategyId:%s,vehicleId:%v", util.RunFuncName(), fstrategyId, vehicleId)
	logger.Logger.Info("%s vehicle_id:%s,fstrategyId:%s,vehicleId:%v", util.RunFuncName(), fstrategyId, vehicleId)

	//查看该vehicle是否存在
	vehicleFStrategy, err := model.GetVehicleFStrategy(
		"fstrategy_vehicles.vehicle_id = ? and fstrategies.fstrategy_id = ?",
		[]interface{}{vehicleId, fstrategyId}...)

	logger.Logger.Print("%s vehicleFStrategy:%+v", util.RunFuncName(), vehicleFStrategy)
	logger.Logger.Info("%s vehicleFStrategy:%+v", util.RunFuncName(), vehicleFStrategy)

	if vehicleFStrategy.FstrategyVehicleId == "" {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s fstrategy join fstrategy_vehicles fstrategyVehicleId null", util.RunFuncName())
		logger.Logger.Error("%s fstrategy join fstrategy_vehicles fstrategyVehicleId null", util.RunFuncName())

		return
	}

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Print("%s fstrategy join fstrategy_vehicles err:%+v", util.RunFuncName(), err)
		logger.Logger.Error("%s fstrategy join fstrategy_vehicles err:%+v", util.RunFuncName(), err)

		return
	}

	vehicleFStrategyItems, err := model.GetVehicleFStrategyItems(
		"fstrategy_vehicle_items.fstrategy_vehicle_id = ?",
		[]interface{}{vehicleFStrategy.FstrategyVehicleId}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	vehicleFStrategyItemsMap := map[string][]uint32{}

	for _, vehicleFStrategyItem := range vehicleFStrategyItems {
		if len(vehicleFStrategyItemsMap[vehicleFStrategyItem.DstIp]) == 0 {
			vehicleFStrategyItemsMap[vehicleFStrategyItem.DstIp] = []uint32{vehicleFStrategyItem.DstPort}
		} else {
			vehicleFStrategyItemsMap[vehicleFStrategyItem.DstIp] =
				append(vehicleFStrategyItemsMap[vehicleFStrategyItem.DstIp], vehicleFStrategyItem.DstPort)
		}
	}

	vehicleSingleFlowStrategyItemsReult := model.VehicleSingleFlowStrategyItemsReult{
		FstrategyId:              vehicleFStrategy.FstrategyId,
		Type:                     vehicleFStrategy.Type,
		HandleMode:               vehicleFStrategy.HandleMode,
		Enable:                   vehicleFStrategy.Enable,
		VehicleId:                vehicleFStrategy.VehicleId,
		CsvPath:                  vehicleFStrategy.CsvPath,
		VehicleFStrategyItemsMap: vehicleFStrategyItemsMap,
	}
	responseData := map[string]interface{}{
		"fstrategy": vehicleSingleFlowStrategyItemsReult,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}
