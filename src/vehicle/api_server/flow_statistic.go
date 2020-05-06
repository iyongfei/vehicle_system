package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func FlowStatistics(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId nill", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty vehicleId nill", util.RunFuncName())
		return
	}

	logger.Logger.Info("%s vehicleId:%s", util.RunFuncName(), vehicleId)
	logger.Logger.Print("%s vehicleId:%s", util.RunFuncName(), vehicleId)

	flowStatistics := &model.FlowStatistic{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(flowStatistics)

	err, recordNotFound := modelBase.GetModelByCondition(
		"vehicle_id = ?", []interface{}{flowStatistics.VehicleId}...)

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s recordNotFound,err:%s", util.RunFuncName(), flowStatistics.VehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s recordNotFound,err:%s", util.RunFuncName(), flowStatistics.VehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), flowStatistics.VehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), flowStatistics.VehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	flowStatisticsList := []*model.FlowStatistic{}

	err = modelBase.GetModelListByCondition(&flowStatisticsList, "vehicle_id = ?", []interface{}{flowStatistics.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), flowStatistics.VehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), flowStatistics.VehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleStatisticListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	responseData := map[string]interface{}{
		"flow_statistics": flowStatisticsList,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetVehicleStatisticListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}
