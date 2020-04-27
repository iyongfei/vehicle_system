package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
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
	vehicleMonitorJoinItems, err := model.GetVehicleMonitorItems("monitors.monitor_id = ?", []interface{}{vehicleId}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetMonitorsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	var monitorModel model.Monitor
	if len(vehicleMonitorJoinItems) > 0 {
		monitor := vehicleMonitorJoinItems[0]
		monitorModel = monitor.Monitor
	}

	var vehicleMonitorItemList []model.Disk
	for _, monitorItem := range vehicleMonitorJoinItems {
		disk := model.Disk{
			Path:     monitorItem.Path,
			DiskRate: monitorItem.DiskRate,
		}

		vehicleMonitorItemList = append(vehicleMonitorItemList, disk)
	}

	vehicleMonitorItemsResponse := model.VehicleMonitorItemsResponse{
		Monitor:                monitorModel,
		VehicleMonitorItemList: vehicleMonitorItemList,
	}

	responseData := map[string]interface{}{
		"mointor": vehicleMonitorItemsResponse,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetMonitorsSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}
