package vgo

import (
	"fmt"
	"vehicle_system/src/vehicle/emq/emq_cmd"
	"vehicle_system/src/vehicle/emq/topic_publish_handler"
	"vehicle_system/src/vehicle/model"
)

func Setup()  {
	startInitStrategyGo()
}
func startInitStrategyGo() {
	go startInitStrategyGoService()
}

func startInitStrategyGoService() {
	for {
		select {
		case vehicleIdInitStrategy := <-model.InitVehicleStrategyChan:
			fmt.Println("vehicleIdInitStrategy", vehicleIdInitStrategy)
			strategyRecent:=model.GetVehicleRecentStrategy(vehicleIdInitStrategy)
			fmt.Printf("strategyCmd:%+v\n", strategyRecent)

			strategyCmd := &emq_cmd.StrategySetCmd{
				VehicleId: strategyRecent.VehicleId,
				TaskType:  strategyRecent.TaskType,

				StrategyId:strategyRecent.StrategyId,
				Type:      strategyRecent.Type,
				HandleMode:strategyRecent.HandleMode,
				Enable:strategyRecent.Enable,
				GroupId:"", //目前不实现
			}

			topic_publish_handler.GetPublishService().PutMsg2PublicChan(strategyCmd)
		}
	}
}
