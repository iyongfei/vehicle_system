package main

import (
	"github.com/golang/protobuf/proto"
	"strconv"
	"vehicle_system/src/vehicle_script/emq_service"
	"vehicle_system/src/vehicle_script/emq_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
insert_vehicle_count
 */
const (
	InsertFlowVehicleId = "insert_flow_vehicle_id"
	InsertFlowCount = "insert_flow_count"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertFlowVehicleId]
	insertFlowCount := configMap[InsertFlowCount]
	defaultVehicleFlowCount ,_ := strconv.Atoi(insertFlowCount)

	emqx:= emq_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatFlowProtobuf(insertVehicleId,defaultVehicleFlowCount))
}


func creatFlowProtobuf(vehicleId string,flowCount int)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_FLOWSTAT,
		GUID:vehicleId,
	}
	flowParams := &protobuf.FlowParam{}
	//添加ThreatItem

	hashList:=[]uint32{1,2,3,4,5}
	list:=[]*protobuf.FlowParam_FItem{}
	for i:=0;i<flowCount;i++{
		moduleItem := &protobuf.FlowParam_FItem{
			Hash:hashList[i],
			SrcIp:131,
			SrcPort:23,
			DstIp:23,
			DstPort:23,
			Protocol:protobuf.FlowProtos(32),
			FlowInfo:"wklejl",
			SafeType:protobuf.FlowSafetype(33),
			SafeInfo:"jwek",
			StartTime:tool.TimeNowToUnix(),
			LastSeenTime:tool.TimeNowToUnix(),
			Src2DstBytes:3233,
			Dst2SrcBytes:43444,
			FlowStat:protobuf.FlowStat_FST_FINISH,
		}
		list = append(list,moduleItem)
	}
	flowParams.FlowItem = list

	deviceParamsBytes,_:=proto.Marshal(flowParams)
	pushReq.Param = deviceParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

