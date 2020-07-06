package emq_client

import (
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/emq_cacha"
	"vehicle_system/src/vehicle/emq/topic_subscribe_handler"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/util"
)

func EmqReConnectTokenError() {

	PushAllVehicleOffLine()

	t := time.NewTicker(time.Second * 60)
	select {
	case <-t.C:
		if !EmqClient.IsConnected() {
			logger.Logger.Print("%s,emqReConnectTokenError:%v", util.RunFuncName(), &EmqClient)
			logger.Logger.Info("%s,emqReConnectTokenError:%v", util.RunFuncName(), &EmqClient)
			GetEmqInstance().InitEmqClient()
		}
		t.Stop()
		return
	}
}
func PushAllVehicleOffLine() {
	vehicleCache := emq_cacha.GetVehicleCache()
	vehicleCache.CleanAllKey()
	//发送请求

	var vehicleIds []string
	_ = mysql.QueryPluckByModelWhere(&model.VehicleInfo{}, "vehicle_id", &vehicleIds,
		"", []interface{}{}...)

	for _, vehicleId := range vehicleIds {
		_ = topic_subscribe_handler.HandleVehicleOfflineStatus(vehicleId, false)
	}
}
