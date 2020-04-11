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
	InsertFirmVehicleId = "insert_firm_vehicle_id"
	Deploy_Id = "deploy_id"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertFirmVehicleId]
	deployId := configMap[Deploy_Id]

	emqx:= emq_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatFirmwareProtobuf(insertVehicleId,deployId))
}


func creatFirmwareProtobuf(vehicleId string,deployId string)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_FIRMWARE,
		GUID:vehicleId,
		CmdID:deployId,
	}
	firmwareParams := &protobuf.FirwareParam{
		Version:tool.GenVersion(),
		UpgradeTimestamp:tool.TimeNowToUnix(),
		UpgradeStatus:protobuf.FirwareParam_Status(1),
		Timeout:30,
	}
	//添加ThreatItem


	deviceParamsBytes,_:=proto.Marshal(firmwareParams)
	pushReq.Param = deviceParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

