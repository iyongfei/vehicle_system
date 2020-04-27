package main

import (
	"fmt"
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
	insert_flowstatistic_vehicle_id = "insert_flowstatistic_vehicle_id"
)

func main() {
	configMap := tool.InitConfig("conf.txt")
	insertFlowStatisticVehicleId := configMap[insert_flowstatistic_vehicle_id]

	fmt.Println(insertFlowStatisticVehicleId)
	emqx := emq_service.NewEmqx()
	emqx.Publish(insertFlowStatisticVehicleId, creatFlowStatisticProtobuf(insertFlowStatisticVehicleId))
}

func creatFlowStatisticProtobuf(vehicleId string) []byte {
	pushReq := &protobuf.GWResult{
		ActionType: protobuf.GWResult_FLOWSTATISTIC,
		GUID:       vehicleId,
	}
	flowStatisticParamParams := &protobuf.FlowStatisticParam{}
	//添加ThreatItem
	flowStatisticParamParams.InterfaceName = "ne2"
	flowStatisticParamParams.Rx = 12312
	flowStatisticParamParams.Tx = 34
	flowStatisticParamParams.FlowCount = 2342
	flowStatisticParamParams.Pub = 232
	flowStatisticParamParams.Notlocal = 232
	flowStatisticParamParams.White = 232

	portMapParamsBytes, _ := proto.Marshal(flowStatisticParamParams)

	pushReq.Param = portMapParamsBytes
	ret, _ := proto.Marshal(pushReq)
	return ret
}
