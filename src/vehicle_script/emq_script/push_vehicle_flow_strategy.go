package main

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle_script/emq_service"
	"vehicle_system/src/vehicle_script/emq_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
 */
const (
	InsertFlowStrategyVehicleId = "insert_flow_strategy_vehicle_id"
	FlowStrategyId = "flow_strategy_id"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertFlowStrategyVehicleId]
	flowStrategyId := configMap[FlowStrategyId]

	emqx:= emq_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatFlowStrategyProtobuf(insertVehicleId,flowStrategyId))
}

func creatFlowStrategyProtobuf(vehicleId string,flowStrategyId string)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_FLOWSTRATEGYSTAT,
		GUID:vehicleId,
	}
	flowStrategyParams := &protobuf.FlowStrategyParam{
		FlowStrategyId:flowStrategyId,
		HandleMode:protobuf.FlowStrategyParam_WARNING,
		DefenseType:protobuf.FlowStrategyParam_FLWOWHITEMODE,
		Enable:false,
	}


	portMapParamsBytes,_:=proto.Marshal(flowStrategyParams)
	pushReq.Param = portMapParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

