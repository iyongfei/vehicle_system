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
#monitor
monitor_id = 1233
disk_path_count = 5
*/
const (
	Monitor_Id      = "monitor_id"
	Disk_Path_Count = "disk_path_count"
)

func main() {
	configMap := tool.InitConfig("conf.txt")
	MonitorId := configMap[Monitor_Id]
	DiskPathCount := configMap[Disk_Path_Count]
	defaultDiskPathCount, _ := strconv.Atoi(DiskPathCount)

	emqx := emq_service.NewEmqx()
	//vid:=tool.RandomString(32)
	emqx.Publish(MonitorId, createMonitorProbuf(MonitorId, defaultDiskPathCount))
}

func createMonitorProbuf(vId string, defaultDiskPathCount int) []byte {
	pushReq := &protobuf.GWResult{
		ActionType: protobuf.GWResult_MONITORINFO,
		GUID:       vId,
	}

	params := &protobuf.MonitorInfoParam{
		GatherTime: 1231231231,
	}

	items := []*protobuf.MonitorInfoParam_DiskOverFlow{}

	for i := 0; i < defaultDiskPathCount; i++ {
		moduleItem := &protobuf.MonitorInfoParam_DiskOverFlow{
			Path:     tool.RandomString(8),
			DiskRate: 0.4,
		}
		items = append(items, moduleItem)
	}

	params.DiskItem = items

	redisInfo := &protobuf.MonitorInfoParam_RedisInfo{
		Active:  true,
		CpuRate: 0.4,
		MemRate: 0.2,
		Mem:     23222323232332,
	}

	vhaloInfo := &protobuf.MonitorInfoParam_VHaloNets{
		Active:  false,
		CpuRate: 0.4,
		MemRate: 0.2,
		Mem:     2322332,
	}
	params.RedisInfo = redisInfo
	params.VhaloInfo = vhaloInfo

	bys, _ := proto.Marshal(params)
	///////////////////////////////////

	pushReq.Param = bys
	ret, _ := proto.Marshal(pushReq)
	return ret
}
