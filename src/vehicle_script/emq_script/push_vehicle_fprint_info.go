package main

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle_script/emq_service"
	"vehicle_system/src/vehicle_script/emq_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
insert_vehicle_count
*/
const (
	insert_fprint_vehicle_id = "insert_fprint_vehicle_id"
	insert_fprint_mac        = "insert_fprint_mac"
)

func main() {
	configMap := tool.InitConfig("conf.txt")
	insert_fprint_vehicle_id := configMap[insert_fprint_vehicle_id]
	insert_fprint_mac := configMap[insert_fprint_mac]

	emqx := emq_service.NewEmqx()
	emqx.Publish(insert_fprint_vehicle_id, createGwFprintProbuf(insert_fprint_vehicle_id, insert_fprint_mac))
}

func createGwFprintProbuf(vId string, mac string) []byte {
	pushReq := &protobuf.GWResult{
		ActionType: protobuf.GWResult_FINGERPRINT,
		GUID:       vId,
	}

	active := &protobuf.FingerprintParam_ActiveDetect{
		Os: "centos",
	}

	items := []*protobuf.PassiveInfoItem{}
	for i := 0; i < 5; i++ {
		moduleItem := &protobuf.PassiveInfoItem{
			Hash:         uint32(tool.RandomNumber(5)),
			SrcIp:        131,
			SrcPort:      23,
			DstIp:        23,
			DstPort:      23,
			Protocol:     protobuf.FlowProtos(32),
			FlowInfo:     "wklejl",
			SafeType:     protobuf.FlowSafetype(33),
			SafeInfo:     "jwek",
			StartTime:    tool.TimeNowToUnix(),
			LastSeenTime: tool.TimeNowToUnix(),
			Src2DstBytes: 3233,
			Dst2SrcBytes: 43444,
			FlowStat:     protobuf.FlowStat_FST_FINISH,
		}

		items = append(items, moduleItem)
	}

	//PassiveInfoItems

	passive := &protobuf.FingerprintParam_PassiveLearn{
		DstPort: 200,
		Item:    items,
	}
	params := &protobuf.FingerprintParam{
		Mac:     mac,
		Active:  active,
		Passive: passive,
	}

	bys, _ := proto.Marshal(params)
	///////////////////////////////////

	pushReq.Param = bys
	ret, _ := proto.Marshal(pushReq)
	return ret
}
