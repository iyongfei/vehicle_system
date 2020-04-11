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
	InsertProtectVehicleId = "insert_protect_vehicle_id"
	ProtectStatus = "protect_status"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertProtectVehicleId]
	protectStatus := configMap[ProtectStatus]
	defaultProtectStatus ,_ := strconv.Atoi(protectStatus)

	emqx:= emq_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatProtectProtobuf(insertVehicleId,defaultProtectStatus))
}

func creatProtectProtobuf(vehicleId string,defaultProtectStatus int)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_PROTECT,
		GUID:vehicleId,
	}
	protectParams := &protobuf.GWProtectInfoParam{}
	//添加ThreatItem
	protectParams.ProtectStatus = protobuf.GWProtectInfoParam_Status(defaultProtectStatus)

	portMapParamsBytes,_:=proto.Marshal(protectParams)
	pushReq.Param = portMapParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

