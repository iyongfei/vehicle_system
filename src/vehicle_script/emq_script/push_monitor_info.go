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
 */
const (
	Monitor_Id = "monitor_id"
	Disk_Path_Count = "disk_path_count"
)

func main()  {
	configMap := tool.InitConfig("conf.txt")
	MonitorId := configMap[Monitor_Id]
	DiskPathCount := configMap[Disk_Path_Count]
	defaultDiskPathCount ,_ := strconv.Atoi(DiskPathCount)

	emqx:= emq_service.NewEmqx()
	//vid:=tool.RandomString(32)
	emqx.Publish(MonitorId,createMonitorProbuf(MonitorId,defaultDiskPathCount))
}

func createMonitorProbuf(vId string,defaultDiskPathCount int)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_MONITORINFO,
		GUID:vId,
	}

	params := &protobuf.MonitorInfoParam{
		CpuRate:12,
		MemRate:10,
		GatherTime:1231231231,
	}



	//module begin
	items:=[]*protobuf.MonitorInfoParam_DiskOverFlow{}

	for i:=0;i<defaultDiskPathCount;i++{
		fmt.Println("lwejl",i)
		moduleItem := &protobuf.MonitorInfoParam_DiskOverFlow{
			//Path:tool.RandomString(4),
			Path:"x9Es",
			DiskRate:70,
		}
		items = append(items,moduleItem)
	}

	params.DiskItem = items

	fmt.Println("jsldfks",params)

	bys,_:=proto.Marshal(params)
	///////////////////////////////////

	pushReq.Param = bys
	ret,_:=proto.Marshal(pushReq)
	return  ret
}