package cron

import (
	"time"
	"vehicle_system/src/vehicle/emq/emq_cmd"
	"vehicle_system/src/vehicle/emq/topic_publish_handler"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
)

func perMinuteFun() {

	logger.Logger.Print("PerMinuteFun %v", time.Now())

	vehicleInfos := []*model.VehicleInfo{}
	err := model_base.ModelBaseImpl(&model.VehicleInfo{}).
		GetModelListByCondition(&vehicleInfos, "", []interface{}{}...)
	if err != nil {
		return
	}

	for _, vehicle := range vehicleInfos {
		recentStrategy := model.GetVehicleRecentStrategy(vehicle.VehicleId)
		strategyCmd := &emq_cmd.StrategySetCmd{
			VehicleId: recentStrategy.VehicleId,
			TaskType:  recentStrategy.TaskType,

			StrategyId: recentStrategy.StrategyId,
			Type:       recentStrategy.Type,
			HandleMode: recentStrategy.HandleMode,
			Enable:     recentStrategy.Enable,
			GroupId:    "", //目前不实现
		}
		topic_publish_handler.GetPublishService().PutMsg2PublicChan(strategyCmd)
	}
}
