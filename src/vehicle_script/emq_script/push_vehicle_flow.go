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
	items:=[]*protobuf.ThreatParam_Item{}

	for i:=0;i<flowCount;i++{
		moduleItem := &protobuf.ThreatParam_Item{
			SrcMac:tool.RandomString(8),
			ThreatType:protobuf.ThreatParam_Item_SITE,
			Content:"威胁内容"+tool.RandomString(8),
			ThreatStatus:protobuf.ThreatParam_Item_PREVENT,
			AttactTime:tool.TimeNowToUnix(),
			SrcIp:tool.GenIpAddr(),
			DstIp:tool.GenIpAddr(),
		}
		items = append(items,moduleItem)
	}
	threatParams.ThreatItem = items

	deviceParamsBytes,_:=proto.Marshal(threatParams)
	pushReq.Param = deviceParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

