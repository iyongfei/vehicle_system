package main

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle_script/emq_service"
	"vehicle_system/src/vehicle_script/emq_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
insert_vehicle_count
 */
const (
	InsertStrategyVehicleId = "insert_strategy_vehicle_id"
	StrategyId = "strategy_id"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertStrategyVehicleId]
	strategyId := configMap[StrategyId]

	emqx:= emq_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatStrategyProtobuf(insertVehicleId,strategyId))
}

func creatStrategyProtobuf(vehicleId string,strategyId string)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_STRATEGY,
		GUID:vehicleId,
	}
	strategyParams := &protobuf.StrategyParam{
		StrategyId:strategyId,
		HandleMode:protobuf.StrategyParam_WARNING,
		DefenseType:protobuf.StrategyParam_WHITEMODE,
		Enable:false,
	}


	portMapParamsBytes,_:=proto.Marshal(strategyParams)
	pushReq.Param = portMapParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

