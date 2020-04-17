package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func EditStrategy(c *gin.Context) {
	strategyId := c.Param("strategy_id")
	setTypeP := c.PostForm("type")
	handleModeP := c.PostForm("handle_mode")

	argsTrimsEmpty := util.RrgsTrimsEmpty(strategyId, setTypeP, handleModeP)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
	}

	setType, _ := strconv.Atoi(setTypeP)
	handleMode, _ := strconv.Atoi(handleModeP)

	//查询是否存在
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
		"type": strategyInfo.Type,
	}
	if err:=modelBase.UpdateModelsByCondition(attrs,"strategy_id = ?",
		[]interface{}{strategyInfo.StrategyId}...);err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqUpdateStrategyFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	//更新
	//assetCmd := &emq_cmd.AssetSetCmd{
	//	VehicleId: assetInfo.VehicleId,
	//	TaskType:  int(protobuf.Command_DEVICE_SET),
	//
	//	Switch: setSwitch,
	//	Type:   setType,
	//	Mac:    assetId,
	//}
	//
	//topic_publish_handler.GetPublishService().PutMsg2PublicChan(assetCmd)

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateStrategySuccessMsg, "")
	c.JSON(http.StatusOK, retObj)

}

func GetStrategys(c *gin.Context) {
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
		"strategys":   strategys,
		"totalCount": total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetStrategyListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func GetStrategy(c *gin.Context) {
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


func AddStrategy(c *gin.Context) {
	sType := c.PostForm("type")
	handleMode := c.PostForm("handle_mode")
	learningResultIds := c.PostForm("learning_result_ids")
	vehicleId := c.PostForm("vehicle_id")

	vStype, tErr := strconv.Atoi(sType)
	vHandleMode, hErr := strconv.Atoi(handleMode)

	logger.Logger.Print("%s sType:%s,handleMode:%s,learningResultIds:%s", util.RunFuncName(), sType, handleMode, learningResultIds)
	logger.Logger.Info("%s sType:%s,handleMode:%s,learningResultIds:%s", util.RunFuncName(), sType, handleMode, learningResultIds)

	argsTrimsEmpty := util.RrgsTrimsEmpty(sType, handleMode, learningResultIds,vehicleId)
	if argsTrimsEmpty || tErr != nil || hErr != nil {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty sType:%s,handleMode:%s,learningResultIds:%s", util.RunFuncName(), sType, handleMode, learningResultIds)
		logger.Logger.Print("%s argsTrimsEmpty sType:%s,handleMode:%s,learningResultIds:%s", util.RunFuncName(), sType, handleMode, learningResultIds)
		return
	}


	strategy := &model.Strategy{
		StrategyId:util.RandomString(32),
		Type:      uint8(vStype),
		HandleMode:     uint8(vHandleMode),
		Enable: true,
	}
	modelBase := model_base.ModelBaseImpl(strategy)

	if err := modelBase.InsertModel(); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	strategyVehicle := &model.StrategyVehicle{
		StrategyId:strategy.StrategyId,
		VehicleId:vehicleId,
	}
	strategyVehicleModelBase := model_base.ModelBaseImpl(strategyVehicle)

	if err := strategyVehicleModelBase.InsertModel(); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddStrategyFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//查groupId
	learningResultIdSlice := strings.Split(learningResultIds,",")
	for _,learningResultId:=range learningResultIdSlice{
		strategyVehicleLearningResult := &model.StrategyVehicleLearningResult{
			VehicleId:vehicleId,
			LearningResultId:learningResultId,
		}

		strategyVehicleLearningResultModelBase := model_base.ModelBaseImpl(strategyVehicleLearningResult)

		if err := strategyVehicleLearningResultModelBase.InsertModel(); err != nil {
			ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddStrategyFailMsg, "")
			c.JSON(http.StatusOK, ret)
			return
		}
	}


	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddStrategySuccessMsg, strategy)
	c.JSON(http.StatusOK, retObj)
}

func DeleStrategy(c *gin.Context) {
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

func GetStrategyVehicle(c *gin.Context) {
	strategyId := c.Param("strategy_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(strategyId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), strategyId)
		logger.Logger.Print("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), strategyId)
	}
	strategyVehicleInfo := &model.StrategyVehicle{
		StrategyId: strategyId,
	}

	modelBase := model_base.ModelBaseImpl(strategyVehicleInfo)

	strategyVehicleInfos := []*model.StrategyVehicle{}
	err:=modelBase.GetModelListByCondition(&strategyVehicleInfos,"strategy_id = ?",[]interface{}{strategyVehicleInfo.StrategyId}...)

	if err != nil {
		logger.Logger.Error("%s strategy_id:%s,err:%s", util.RunFuncName(), strategyId, err)
		logger.Logger.Print("%s strategy_id:%s,err:%s", util.RunFuncName(), strategyId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrategyVehicleListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"strategy_vehicles": strategyVehicleInfos,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetStrategyVehicleListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}



/****************************************StrategyVehicleResult********************************************************/

func GetStrategyVehicleLearningResults(c *gin.Context) {
	vehicleId := c.Param("vehicle_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty strategy_id:%s", util.RunFuncName(), vehicleId)
	}
	strategyVehicleLearnResultInfo := &model.StrategyVehicleLearningResult{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(strategyVehicleLearnResultInfo)

	strategyVehicleLearnResultInfos := []*model.StrategyVehicleLearningResult{}
	err:=modelBase.GetModelListByCondition(&strategyVehicleLearnResultInfos,"vehicle_id = ?",[]interface{}{strategyVehicleLearnResultInfo.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), strategyVehicleLearnResultInfo.VehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), strategyVehicleLearnResultInfo.VehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetStrategyVehicleResultListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"strategy_vehicle_results": strategyVehicleLearnResultInfos,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetStrategyVehicleResultListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}