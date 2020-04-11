package main

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle_script/emq_service"
	"vehicle_system/src/vehicle_script/emq_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
insert_leader_vehicle_id = ff
name=safly
phone=13581922339
dev_name =小V上线
 */
const (
	InsertLeaderVehicleId = "insert_leader_vehicle_id"
	Name = "name"
	Phone = "phone"
	DevName = "dev_name"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	InsertLeaderVehicleId := configMap[InsertLeaderVehicleId]
	name := configMap[Name]
	phone := configMap[Phone]
	devName := configMap[DevName]


	emqx:= emq_service.NewEmqx()
	emqx.Publish(InsertLeaderVehicleId,creatLeaderProtobuf(InsertLeaderVehicleId,name,phone,devName))
}

func creatLeaderProtobuf(vehicleId string,name ,phone,devName string)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_DEPLOYER,
		GUID:vehicleId,
	}
	deployerParams := &protobuf.DeployerParam{
		Name:name,
		Phone:phone,
		DevName:devName,
	}

	portMapParamsBytes,_:=proto.Marshal(deployerParams)
	pushReq.Param = portMapParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

