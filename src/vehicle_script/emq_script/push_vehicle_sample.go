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
 = kQ8XKqP57cNwt0CwXLWoXWF0UfwaYFX8
 = 5
 */
const (
	InsertSampleVehicleId = "insert_sample_vehicle_id"
	InsertSampleId = "insert_sample_id"
	InsertSampleItemCount = "sample_item_count"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertSampleVehicleId]
	sampleId := configMap[InsertSampleItemCount]
	inserSampleItemCount := configMap[InsertSampleItemCount]

	defaultSampleItemCount ,_ := strconv.Atoi(inserSampleItemCount)
	fmt.Println("defaultVehicleCount:",defaultSampleItemCount)

	emqx:= emq_service.NewEmqx()
	emqx.Publish(insertVehicleId,creatSampleProtobuf(insertVehicleId,sampleId,defaultSampleItemCount))
}


func creatSampleProtobuf(vehicleId string,sampleId string,defaultSampleItemCount int)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_SAMPLE,
		GUID:vehicleId,
	}
	sampleParams := &protobuf.SampleParam{
		Id:sampleId,
		StartTime:tool.TimeNowToUnix(),
		TimeRemain:30,
		Status:protobuf.SampleParam_COLLECT_OK,
		Timeout:30,
	}
	//添加ThreatItem
	items:=[]*protobuf.SampleParam_Item{}

	for i:=0;i<defaultSampleItemCount;i++{
		moduleItem := &protobuf.SampleParam_Item{
			Sm:tool.RandomString(8),
			Sip:tool.GenIpAddr(),
			Sp:200,
			Dip:tool.GenIpAddr(),
			Dp:200,
			U:"www.bai.com",
		}
		items = append(items,moduleItem)
	}
	sampleParams.SampleItem = items


	deviceParamsBytes,_:=proto.Marshal(sampleParams)
	pushReq.Param = deviceParamsBytes
	ret,_ := proto.Marshal(pushReq)
	return ret
}

