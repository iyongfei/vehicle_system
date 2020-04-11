package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	"vehicle_system/src/vehicle_script/emq_service"
	"vehicle_system/src/vehicle_script/emq_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
insert_vehicle_count
insert_asset_vehicle_id=5
insert_asset_count=5
 */
const (
	InsertAssetVehicleId = "insert_asset_vehicle_id"
	InsertAssetCount = "insert_asset_count"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertAssetVehicleId]
	insertVehicleAssetCount := configMap[InsertAssetCount]
	defaultInsertVehicleAssetCount ,_ := strconv.Atoi(insertVehicleAssetCount)
	fmt.Println("defaultVehicleCount:",defaultInsertVehicleAssetCount)

	emqx:= emq_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatAssetProtobuf(insertVehicleId,defaultInsertVehicleAssetCount))
}


func creatAssetProtobuf(vehicleId string,threatCount int)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_DEVICE,
		GUID:vehicleId,
	}
	threatParams := &protobuf.DeviceParam{}
	//添加ThreatItem
	items:=[]*protobuf.DeviceParam_Item{}

	for i:=0;i<threatCount;i++{
		moduleItem := &protobuf.DeviceParam_Item{
			Ip:tool.GenIpAddr(),
			Mac:tool.RandomString(8),
			Name:tool.RandomString(8),
			Trademark:tool.RandomString(8),
			IsOnline:true,
			LastOnline:tool.TimeNowToUnix(),
			InternetSwitch:false,
			ProtectSwitch:true,
			LanVisitSwitch:false,
		}
		items = append(items,moduleItem)
	}
	threatParams.DeviceItem = items

	deviceParamsBytes,_:=proto.Marshal(threatParams)
	pushReq.Param = deviceParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

