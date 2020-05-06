package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func GetMonitor(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s argsTrimsEmpty", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s argsTrimsEmpty", util.RunFuncName(), vehicleId)
		return
	}
	logger.Logger.Info("%s vehicleId:%s", util.RunFuncName(), vehicleId)
	logger.Logger.Print("%s vehicleId:%s", util.RunFuncName(), vehicleId)
	//判断有无vehicle
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

	//disk
	diskList := []*model.Disk{}
	disk := &model.Disk{
		MonitorId: vehicleId,
	}

	diskModelBase := model_base.ModelBaseImpl(disk)

	diskErr := diskModelBase.GetModelListByCondition(&diskList, "monitor_id = ?", disk.MonitorId)
	if diskErr != nil {
		logger.Logger.Error("%s argsTrimsEmpty diskErr:%+v", util.RunFuncName(), diskErr)
		logger.Logger.Print("%s argsTrimsEmpty diskErr:%+v", util.RunFuncName(), diskErr)

	}

	//redis
	redisInfo := &model.RedisInfo{
		MonitorId: vehicleId,
	}
	redisInfoErr := mysql.QueryModelOneRecordByWhereConditionOrderBy(redisInfo, "id desc", "monitor_id = ?", redisInfo.MonitorId)
	if redisInfoErr != nil {
		logger.Logger.Error("%s argsTrimsEmpty redisInfoErr:%+v", util.RunFuncName(), redisInfoErr)
		logger.Logger.Print("%s argsTrimsEmpty redisInfoErr:%+v", util.RunFuncName(), redisInfoErr)

	}

	//vhalonets
	vhaloInfo := &model.VhaloNets{
		MonitorId: vehicleId,
	}

	vhaloInfoErr := mysql.QueryModelOneRecordByWhereConditionOrderBy(vhaloInfo, "id desc", "monitor_id = ?", vhaloInfo.MonitorId)
	if vhaloInfoErr != nil {
		logger.Logger.Error("%s argsTrimsEmpty vhaloInfoErr:%+v", util.RunFuncName(), vhaloInfoErr)
		logger.Logger.Print("%s argsTrimsEmpty vhaloInfoErr:%+v", util.RunFuncName(), vhaloInfoErr)

	}
	vehicleMonitorItemsResponse := model.VehicleMonitorItemsResponse{
		Disks:     diskList,
		RedisInfo: *redisInfo,
		VhaloNets: *vhaloInfo,
	}

	responseData := map[string]interface{}{
		"mointor": vehicleMonitorItemsResponse,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetMonitorsSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}
