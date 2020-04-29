package topic_subscribe_handler

import (
	"vehicle_system/src/vehicle/db/tdata"
	"vehicle_system/src/vehicle/service/flow"
)

//离线
func HandleVehicleOfflineStatus(vehicleId string, onLineStatus bool) error {
	err := tdata.VehicleAssetCheck(vehicleId, onLineStatus)

	if err == nil {

		pushActionTypeName := flow.ONLINE_STATUS
		pushVehicleid := vehicleId
		pushData := map[string]interface{}{
			"online_status": onLineStatus,
		}

		fPushData := flow.CreatePushData(pushActionTypeName, pushVehicleid, pushData)

		flow.GetFlowService().SetFlowData(fPushData).WriteFlow()

	}

	return err
}
