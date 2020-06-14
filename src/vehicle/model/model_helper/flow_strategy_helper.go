package model_helper

import (
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

type FStrategySetCmd struct {
	VehicleId string
	CmdId     int
	TaskType  int

	FstrategyId string
	Type        int
	HandleMode  int
	Enable      bool
	GroupId     string
}

func GetVehicleRecentFStrategy(vehicleId string) *FStrategySetCmd {
	vehicleAllFStrategys, err := model.GetFStrategyVehicles(
		"fstrategy_vehicles.vehicle_id = ?",
		[]interface{}{vehicleId}...)

	if err != nil {
		return nil
	}
	fstrategySetCmd := &FStrategySetCmd{}

	if len(vehicleAllFStrategys) == 0 {
		//fstrategy table
		fstrategy := &model.Fstrategy{
			FstrategyId: util.RandomString(32),
			Type:        uint8(protobuf.FlowStrategyAddParam_FLWOWHITEMODE),
			HandleMode:  uint8(protobuf.FlowStrategyAddParam_WARNING),
			Enable:      true,
		}
		strategyModelBase := model_base.ModelBaseImpl(fstrategy)

		if err := strategyModelBase.InsertModel(); err != nil {
			return nil
		}

		//fstrategyVehicle table
		fstrategyVehicle := &model.FstrategyVehicle{
			FstrategyVehicleId: util.RandomString(32),
			FstrategyId:        fstrategy.FstrategyId,
			VehicleId:          vehicleId,
		}
		strategyVehicleModelBase := model_base.ModelBaseImpl(fstrategyVehicle)
		if err := strategyVehicleModelBase.InsertModel(); err != nil {
			return nil
		}

		fstrategySetCmd.VehicleId = vehicleId
		fstrategySetCmd.TaskType = int(protobuf.Command_FLOWSTRATEGY_ADD)
		fstrategySetCmd.FstrategyId = fstrategy.FstrategyId
		fstrategySetCmd.Type = int(fstrategy.Type)
		fstrategySetCmd.HandleMode = int(fstrategy.HandleMode)
		fstrategySetCmd.Enable = true
		fstrategySetCmd.GroupId = ""
	} else {
		vehicleRecentStrategy := vehicleAllFStrategys[0]

		fstrategySetCmd.VehicleId = vehicleId
		fstrategySetCmd.TaskType = int(protobuf.Command_FLOWSTRATEGY_ADD)
		fstrategySetCmd.FstrategyId = vehicleRecentStrategy.FstrategyId
		fstrategySetCmd.Type = int(vehicleRecentStrategy.Type)
		fstrategySetCmd.HandleMode = int(vehicleRecentStrategy.HandleMode)
		fstrategySetCmd.Enable = true
		fstrategySetCmd.GroupId = ""
	}
	return fstrategySetCmd
}
