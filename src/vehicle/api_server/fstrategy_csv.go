package api_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
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

func GetFStrategyCsv(c *gin.Context) {
	fstrategyId := c.Param("fstrategy_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(fstrategyId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty fstrategyId:%s,argsTrimsEmpty", util.RunFuncName(), fstrategyId)
		logger.Logger.Print("%s argsTrimsEmpty fstrategyId:%s,argsTrimsEmpty", util.RunFuncName(), fstrategyId)
		return
	}
	logger.Logger.Print("%s vehicle_id:%s,fstrategyId:%s,", util.RunFuncName(), fstrategyId)
	logger.Logger.Info("%s vehicle_id:%s,fstrategyId:%s", util.RunFuncName(), fstrategyId)

	fstrategy := &model.Fstrategy{
		FstrategyId: fstrategyId,
	}
	fstrategyModelBase := model_base.ModelBaseImpl(fstrategy)

	err, recordNotFound := fstrategyModelBase.GetModelByCondition("fstrategy_id = ?", fstrategy.FstrategyId)
	if err != nil {
		logger.Logger.Error("%s fstrategyId:%s,err:%+v", util.RunFuncName(), fstrategyId, err)

		logger.Logger.Print("%s fstrategyId:%s,err:%+v", util.RunFuncName(), fstrategyId, err)

		ret := response.StructResponseObj(response.VStatusServerError, response.ReqFstrategyCsvFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if recordNotFound {
		logger.Logger.Error("%s fstrategyId:%s,recordNotFound", util.RunFuncName(), fstrategyId)
		logger.Logger.Print("%s fstrategyId:%s,recordNotFound", util.RunFuncName(), fstrategyId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFstrategyCsvUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//获取csv文件
	csvPath := fstrategy.CsvPath
	fStrategyCsvFolderIndex := strings.Index(csvPath, csv.FStrategyCsvFolder)

	var csvFileName string
	if fStrategyCsvFolderIndex != -1 {
		csvFileName = csvPath[fStrategyCsvFolderIndex:]
	}

	fmt.Println(csvFileName, "csvFileName")
	c.File(csvFileName)
}

/**
上传scv
*/
func UploadFStrategyCsv(c *gin.Context) {
	uploadCsv, err := c.FormFile("upload_csv")
	vehicleId := c.PostForm("vehicle_id")

	//文件获取失败
	if err != nil {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s upload fstrategy csv formfile err:%+v", util.RunFuncName(), err)
		logger.Logger.Print("%s upload fstrategy csv formfile err:%+v", util.RunFuncName(), err)
		return
	}
	logger.Logger.Print("%s vehicle_id:%s,uploadCsv:%s,", util.RunFuncName(), vehicleId, uploadCsv.Filename)
	logger.Logger.Info("%s vehicle_id:%s,uploadCsv:%s", util.RunFuncName(), vehicleId, uploadCsv.Filename)

	//终端id为空
	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)

	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s upload fstrategy vehicleId:%s,argsTrimsEmpty", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s upload fstrategy vehicleId:%s,argsTrimsEmpty", util.RunFuncName(), vehicleId)
		return
	}
	//查看vehicle是否存在
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s upload fstrategy csv vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s upload fstrategy csv vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s upload fstrategy csv vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s upload fstrategy csv vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	uploadFileName := uploadCsv.Filename

	logger.Logger.Info("%s fileName:%s, vehicleId:%s,err:%+v", util.RunFuncName(), uploadFileName, vehicleId, err)
	logger.Logger.Print("%s fileName:%s, vehicleId:%s,err:%+v", util.RunFuncName(), uploadFileName, vehicleId, err)

	//创建文件
	tempCsvName := util.RandomString(16)
	tempCsvFileFolderPath, _ := csv.CreateCsvFolder()
	tempCsvPathName := tempCsvFileFolderPath + "/" + tempCsvName

	if err := c.SaveUploadedFile(uploadCsv, tempCsvPathName); err != nil {
	}
	//解析
	csvReaderModel := csv.CreateCsvReader(tempCsvPathName)
	//map[string]map[string][]uint32
	parseData, _ := csvReaderModel.ParseAddCsvFile(vehicleId)

	if len(parseData) == 0 {
		//删除文件
		if csv.IsExists(tempCsvPathName) {
			os.Remove(tempCsvPathName)
		}
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s vehicleId:%s,csv format err", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicleId:%s,csv format err", util.RunFuncName(), vehicleId)
		return
	}

	//fstrategy_items table
	fstrategyItems := map[string][]string{}

	for vehicleIdK, vehicleDipPortMap := range parseData {
		var vehicleFstrategyItems []string
		for dip, portSlice := range vehicleDipPortMap {
			for _, dport := range portSlice {

				fstrategyItem := &model.FstrategyItem{
					FstrategyItemId: util.RandomString(32),
					VehicleId:       vehicleIdK,
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
		fstrategyItems[vehicleIdK] = vehicleFstrategyItems
	}
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
	vehicleFstrategyVehicleIdMap := map[string]string{}
	for vehicleIdK, _ := range parseData {
		fstrategyVehicle := &model.FstrategyVehicle{
			FstrategyVehicleId: util.RandomString(32),
			FstrategyId:        fstrategy.FstrategyId,
			VehicleId:          vehicleIdK,
		}
		fstrategyVehicleModelBase := model_base.ModelBaseImpl(fstrategyVehicle)
		if err := fstrategyVehicleModelBase.InsertModel(); err != nil {
			logger.Logger.Print("%s vehicle_id:%s insert FstrategyVehicle err:%+v", util.RunFuncName(), vehicleIdK, err)
			logger.Logger.Info("%s vehicle_id:%s insert FstrategyVehicle err:%+v", util.RunFuncName(), vehicleIdK, err)
			continue
		}

		vehicleFstrategyVehicleIdMap[vehicleIdK] = fstrategyVehicle.FstrategyVehicleId
	}
	//fstrategy_vehicle_items table

	//fstrategy_items table
	//fstrategyItems := map[string][]string{}
	//fstrategyItems[vehicleIdK] = vehicleFstrategyItems

	//FstrategyVehicle table
	//vehicleFstrategyVehicleIdMap := map[string]string{}
	//vehicleFstrategyVehicleIdMap[vehicleIdK] = fstrategyVehicle.FstrategyVehicleId

	for vehicleIdK, _ := range parseData {
		vehicleFsItems := fstrategyItems[vehicleIdK]
		FstrategyVehicleIdMapvehicleItem := vehicleFstrategyVehicleIdMap[vehicleIdK]

		for _, item := range vehicleFsItems {
			fstrategyVehicleItem := &model.FstrategyVehicleItem{
				FstrategyVehicleId: FstrategyVehicleIdMapvehicleItem,
				FstrategyItemId:    item,
			}

			fstrategyVehicleItemModelBase := model_base.ModelBaseImpl(fstrategyVehicleItem)

			if err := fstrategyVehicleItemModelBase.InsertModel(); err != nil {
				logger.Logger.Print("%s vehicle_id:%s insert fstrategyVehicleItem err:%+v", util.RunFuncName(), vehicleIdK, err)
				logger.Logger.Info("%s vehicle_id:%s insert fstrategyVehicleItem err:%+v", util.RunFuncName(), vehicleIdK, err)

				continue
			}
		}
	}
	//插入csv
	csvModel := csv.NewCsvWriter(fstrategy.FstrategyId, csv.FileAppend)
	fCsvHeader := csv.CreateCsvFstrategyHeader()

	//parseData：map[754d2728b4e549c5a16c0180fcacb800:map[192.167.1.3:[123 125 23] 192.168.1.5:[123 125 23]]]

	var csvFstrategyModelBodyList []csv.CsvFstrategyModelHeader
	for vehicleIdK, vehicleDipPorts := range parseData {
		fCsvBody := csv.CreateCsvFstrategyBody(vehicleIdK, fstrategy.FstrategyId, vehicleDipPorts)
		csvFstrategyModelBodyList = append(csvFstrategyModelBodyList, fCsvBody.CsvFstrategyModelBody...)
	}
	csvFstrategyModelBody := csv.CsvFstrategyModelBody{
		CsvFstrategyModelBody: csvFstrategyModelBodyList,
	}

	csvModel.SetCsvWritData(fCsvHeader, csvFstrategyModelBody)

	attrs := map[string]interface{}{
		"csv_path": csvModel.CsvFilePath,
	}
	if err := fstrategyModelBase.UpdateModelsByCondition(attrs, "fstrategy_id = ?", fstrategy.FstrategyId); err != nil {
		logger.Logger.Print("%s vehicle_id:%s insert fstrategy csv_path err:%+v", util.RunFuncName(), csvModel.CsvFilePath, err)
		logger.Logger.Info("%s vehicle_id:%s insert fstrategy csv_path err:%+v", util.RunFuncName(), csvModel.CsvFilePath, err)
	} else {
		fstrategy.CsvPath = csvModel.CsvFilePath
	}

	//删除文件

	if csv.IsExists(tempCsvPathName) {
		os.Remove(tempCsvPathName)
	}

	//下发策略
	for vehicleIdK, _ := range parseData {
		fstrategyCmd := &emq_cmd.FStrategySetCmd{
			VehicleId: vehicleIdK,
			TaskType:  int(protobuf.Command_FLOWSTRATEGY_ADD),

			FstrategyId: fstrategy.FstrategyId,
			Type:        int(protobuf.FlowStrategyAddParam_FLWOWHITEMODE),
			HandleMode:  int(protobuf.FlowStrategyAddParam_WARNING),
			Enable:      true,
			GroupId:     "", //目前不实现
		}
		topic_publish_handler.GetPublishService().PutMsg2PublicChan(fstrategyCmd)
	}

	vehicleSingleFlowStrategyItemsReult := model.VehicleSingleFlowStrategyItemsReult{
		FstrategyId:              fstrategy.FstrategyId,
		Type:                     fstrategy.Type,
		HandleMode:               fstrategy.HandleMode,
		Enable:                   fstrategy.Enable,
		VehicleId:                vehicleId,
		ScvPath:                  fstrategy.CsvPath,
		VehicleFStrategyItemsMap: parseData[vehicleId],
	}

	responseData := map[string]interface{}{
		"fstrategy": vehicleSingleFlowStrategyItemsReult,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddFStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}

func EditFStrategyCsv(c *gin.Context) {
	uploadCsv, err := c.FormFile("upload_csv")
	fstrategyId := c.Param("fstrategy_id")
	vehicleId := c.PostForm("vehicle_id")

	fmt.Println(uploadCsv, fstrategyId, vehicleId, "hahahahah")
	logger.Logger.Print("%s vehicle_id:%s,uploadCsv:%s,fstrategyId:%s", util.RunFuncName(), vehicleId, uploadCsv.Filename, fstrategyId)
	logger.Logger.Info("%s vehicle_id:%s,uploadCsv:%s,fstrategyId:%s", util.RunFuncName(), vehicleId, uploadCsv.Filename, fstrategyId)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s formfile err:%+v", util.RunFuncName(), err)
		logger.Logger.Print("%s formfile err:%+v", util.RunFuncName(), err)
		return
	}
	//终端id为空
	argsTrimsEmpty := util.RrgsTrimsEmpty(fstrategyId, vehicleId)

	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s fstrategyId:%s,vehicleId:%s,argsTrimsEmpty", util.RunFuncName(), fstrategyId, vehicleId)
		logger.Logger.Print("%s fstrategyId:%s,vehicleId:%s,argsTrimsEmpty", util.RunFuncName(), fstrategyId, vehicleId)
		return
	}

	//查看该vehicle是否存在
	vehicleFStrategy, err := model.GetVehicleFStrategy(
		"fstrategy_vehicles.vehicle_id = ? and fstrategies.fstrategy_id = ?", []interface{}{vehicleId, fstrategyId}...)
	if vehicleFStrategy.FstrategyVehicleId == "" {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFStrtegyUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	uploadFileName := uploadCsv.Filename

	logger.Logger.Info("%s fileName:%s, vehicleId:%s,err:%+v", util.RunFuncName(), uploadFileName, err)
	logger.Logger.Print("%s fileName:%s, vehicleId:%s,err:%+v", util.RunFuncName(), uploadFileName, err)

	//创建文件
	tempCsvName := util.RandomString(16)
	tempCsvFileFolderPath, _ := csv.CreateCsvFolder()
	tempCsvPathName := tempCsvFileFolderPath + "/" + tempCsvName

	if err := c.SaveUploadedFile(uploadCsv, tempCsvPathName); err != nil {
	}
	//解析
	csvReaderModel := csv.CreateCsvReader(tempCsvPathName)
	//map[string]map[string][]uint32
	parseData, _ := csvReaderModel.ParseEditCsvFile(fstrategyId, vehicleId)

	if len(parseData) == 0 {
		//删除文件
		if csv.IsExists(tempCsvPathName) {
			os.Remove(tempCsvPathName)
		}

		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s fstrategyId:%s,vehicleId:%s,csv format err", util.RunFuncName(), fstrategyId, vehicleId)
		logger.Logger.Print("%s fstrategyId:%s,vehicleId:%s,csv format err", util.RunFuncName(), fstrategyId, vehicleId)
		return
	}

	//FstrategyVehicle,Fstrategy不更改

	//fstrategy_items table
	fstrategyItems := map[string][]string{}

	for vehicleIdK, vehicleDipPortMap := range parseData {
		var vehicleFstrategyItems []string
		for dip, portSlice := range vehicleDipPortMap {
			for _, dport := range portSlice {

				fstrategyItem := &model.FstrategyItem{
					FstrategyItemId: util.RandomString(32),
					VehicleId:       vehicleIdK,
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
		fstrategyItems[vehicleIdK] = vehicleFstrategyItems
	}

	/**
	VehicleId,FstrategyId,Ip,Port
	754d2728b4e549c5a16c0180fcacb800,HFiYobVy2dqYiVcGpcsrk6GVRxUdqpuy,192.167.1.9,123:124
	754d2728b4e549c5a16c0180fcacb800,HFiYobVy2dqYiVcGpcsrk6GVRxUdqpuy,192.167.1.8,123:121
	*/
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
	fCsvBody := csv.CreateCsvFstrategyBody(vehicleId, vehicleFStrategy.FstrategyId, parseData[vehicleId])
	csvModel.SetCsvWritData(fCsvHeader, fCsvBody)

	//attrs := map[string]interface{}{
	//	"scv_path": csvModel.CsvFilePath,
	//}
	//if err := fstrategyModelBase.UpdateModelsByCondition(attrs, "fstrategy_id = ?", fstrategy.FstrategyId); err != nil {
	//	logger.Logger.Print("%s vehicle_id:%s insert fstrategy scv_path err:%+v", util.RunFuncName(), csvModel.CsvFilePath, err)
	//	logger.Logger.Info("%s vehicle_id:%s insert fstrategy scv_path err:%+v", util.RunFuncName(), csvModel.CsvFilePath, err)
	//}

	//删除文件

	if csv.IsExists(tempCsvPathName) {
		os.Remove(tempCsvPathName)
	}

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
		ScvPath:                  vehicleFStrategy.ScvPath,
		VehicleFStrategyItemsMap: parseData[vehicleId],
	}

	responseData := map[string]interface{}{
		"put_fstrategy": vehicleSingleFlowStrategyItemsReult,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateStrategySuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}
