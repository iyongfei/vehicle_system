package model

import (
	"fmt"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

type StrategySetCmd struct {
	VehicleId string
	CmdId     int
	TaskType  int

	StrategyId string
	Type       int
	HandleMode int
	Enable     bool
	GroupId    string
}

func GetVehicleRecentStrategy(vehicleId string) *StrategySetCmd {
	vehicleAllStrategys ,err := GetVehicleAllStrategys(
		"strategy_vehicles.vehicle_id = ?",
		[]interface{}{vehicleId}...)

	if err!=nil{
		return nil
	}
	strategySetCmd := &StrategySetCmd{}

	if  len(vehicleAllStrategys) == 0{
		//strategy table
		strategy := &Strategy{
			StrategyId:util.RandomString(32),
			Type:      uint8(protobuf.StrategyParam_WHITEMODE),
			HandleMode:     uint8(protobuf.StrategyParam_WARNING),
			Enable: true,
		}
		strategyModelBase := model_base.ModelBaseImpl(strategy)

		if err := strategyModelBase.InsertModel(); err != nil {
			return nil
		}
		//strategyVehicle table
		strategyVehicle := &StrategyVehicle{
			StrategyVehicleId:util.RandomString(32),
			StrategyId:strategy.StrategyId,
			VehicleId:vehicleId,
		}
		strategyVehicleModelBase := model_base.ModelBaseImpl(strategyVehicle)
		if err := strategyVehicleModelBase.InsertModel(); err != nil {
			return nil
		}

		fmt.Printf("strategy%+v",strategy)
		fmt.Printf("strategyVehicle%+v",strategyVehicle)
		strategySetCmd.VehicleId = vehicleId
		strategySetCmd.TaskType = int(protobuf.Command_STRATEGY_ADD)
		strategySetCmd.StrategyId = strategy.StrategyId
		strategySetCmd.Type =  int(strategy.Type)
		strategySetCmd.HandleMode =  int(strategy.HandleMode)
		strategySetCmd.Enable =  true
		strategySetCmd.GroupId = ""
	}else{
		vehicleRecentStrategy := vehicleAllStrategys[0]

		strategySetCmd.VehicleId = vehicleId
		strategySetCmd.TaskType = int(protobuf.Command_STRATEGY_ADD)
		strategySetCmd.StrategyId = vehicleRecentStrategy.StrategyId
		strategySetCmd.Type =   int(vehicleRecentStrategy.Type)
		strategySetCmd.HandleMode =  int(vehicleRecentStrategy.HandleMode)
		strategySetCmd.Enable =  true
		strategySetCmd.GroupId = ""
	}
	return strategySetCmd
}
