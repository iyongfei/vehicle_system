package cron

import (
	"time"
	"vehicle_system/src/vehicle/logger"
)

func perMinuteFun()  {

	logger.Logger.Print("PerMinuteFun %v",time.Now())


	//strategyCmd := &emq_cmd.StrategySetCmd{
	//	VehicleId: vehicleId,
	//	TaskType:  int(protobuf.Command_STRATEGY_ADD),
	//
	//	StrategyId:strategy.StrategyId,
	//	Type:      vStype,
	//	HandleMode:vHandleMode,
	//	Enable:true,
	//	GroupId:"", //目前不实现
	//}
	//topic_publish_handler.GetPublishService().PutMsg2PublicChan(strategyCmd)

}