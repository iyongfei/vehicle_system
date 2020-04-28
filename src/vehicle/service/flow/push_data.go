package flow

const (
	ActionType = "action_type"
	VehicleId  = "vehicle_id"
	PushData   = "push_data"
)

func CreatePushData(actionType string, vehicleId string, pushData interface{}) map[string]interface{} {
	fPushData := map[string]interface{}{}
	fPushData[ActionType] = actionType
	fPushData[VehicleId] = vehicleId
	fPushData[PushData] = pushData
	return fPushData
}
