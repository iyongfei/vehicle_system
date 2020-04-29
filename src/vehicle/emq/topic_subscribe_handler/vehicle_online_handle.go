package topic_subscribe_handler

import "vehicle_system/src/vehicle/db/tdata"

func HandleVehicleOnline(vehicleId string, onLineStatus bool) error {
	err := tdata.VehicleAssetCheck(vehicleId, onLineStatus)
	return err
}
