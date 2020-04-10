package topic_subscribe_handler

import (
	"fmt"
	"time"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
)

func HandleVehicleCommonAction(vehicleResult protobuf.GWResult) error {

	vehicleId :=vehicleResult.GetGUID()

	vehicleInfo:=&model.VehicleInfo{
		VehicleId:vehicleId,
		StartTime:time.Now(),
	}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err,recordNotFound :=modelBase.GetModelByCondition("vehicle_id = ?",vehicleInfo.VehicleId)

	if err!=nil{
		return fmt.Errorf("vehicleId %s not exist",vehicleId)
	}
	if recordNotFound{
		fmt.Println("HandleVehicleCommonAction",recordNotFound,vehicleId)
		err:= modelBase.InsertModel()
		if err!=nil{
			return err
		}
	}

	return nil
}





