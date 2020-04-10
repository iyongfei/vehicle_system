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
	InsertPortMapVehicleId = "insert_port_vehicle_id"
	PortMapCount = "port_map_count"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertPortMapVehicleId]
	portMapCount := configMap[PortMapCount]
	defaultPortMapCount ,_ := strconv.Atoi(portMapCount)

	emqx:= emq_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatPortMapProtobuf(insertVehicleId,defaultPortMapCount))
}


func creatPortMapProtobuf(vehicleId string,defaultPortMapCount int)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_PORTREDIRECT,
		GUID:vehicleId,
	}
	portMapParams := &protobuf.PortRedirectParam{}
	//添加ThreatItem

	list:=[]*protobuf.PortRedirectParam_Item{}
	for i:=0;i<defaultPortMapCount;i++{
		moduleItem := &protobuf.PortRedirectParam_Item{
			SrcPort:strconv.Itoa(i),
			DestPort:strconv.Itoa(i),
			DestIp:tool.GenIpAddr(),
			Switch:false,
			Proto:protobuf.PortRedirectParam_Item_TCP,
		}
		list = append(list,moduleItem)
	}
	portMapParams.PortRedirect = list

	portMapParamsBytes,_:=proto.Marshal(portMapParams)
	pushReq.Param = portMapParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

