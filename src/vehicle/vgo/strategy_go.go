package vgo

import (
	"fmt"
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
			strategyCmd:=model.GetVehicleRecentStrategy(vehicleIdInitStrategy)
			fmt.Printf("strategyCmd:%+v\n", strategyCmd)
			topic_publish_handler.GetPublishService().PutMsg2PublicChan(strategyCmd)
		}
	}
}
