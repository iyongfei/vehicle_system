package topic_subscribe_handler

import (
	"vehicle_system/src/vehicle/service/push"
)

//离线
func HandleVehicleOfflineStatus(vehicleId string, onLineStatus bool) error {
	//err := tdata.VehicleAssetCheck(vehicleId, onLineStatus)
	//if err == nil {
	pushActionTypeName := push.ONLINE_STATUS
	pushVehicleid := vehicleId
	pushData := map[string]interface{}{
		"online_status": onLineStatus,
	}

	fPushData := push.CreatePushData(pushActionTypeName, pushVehicleid, pushData)
	push.GetPushervice().SetPushData(fPushData).Write()
	//}
	return nil
}
