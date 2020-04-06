package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	"vehicle_system/src/vehicle_script/emp_service"
	"vehicle_system/src/vehicle_script/emp_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
insert_vehicle_count
 */
const (
	InsertVehicleId = "insert_vehicle_id"
	InsertVehicleThreatCount = "insert_vehicle_threat_count"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertVehicleId]
	insertVehicleThreatCount := configMap[InsertVehicleThreatCount]
	defaultVehicleThreatCount ,_ := strconv.Atoi(insertVehicleThreatCount)
	fmt.Println("defaultVehicleCount:",defaultVehicleThreatCount)

	emqx:=emp_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatThreatProtobuf(insertVehicleId,defaultVehicleThreatCount))
}


func creatThreatProtobuf(vehicleId string,threatCount int)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_THREAT,
		GUID:vehicleId,
	}
	threatParams := &protobuf.ThreatParam{}
	//添加ThreatItem
	items:=[]*protobuf.ThreatParam_Item{}

	for i:=0;i<threatCount;i++{
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

