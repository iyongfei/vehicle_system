package main

import (
	"github.com/golang/protobuf/proto"
	"strconv"
	"strings"
	"vehicle_system/src/vehicle_script/emq_service"
	"vehicle_system/src/vehicle_script/emq_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
insert_flow_vehicle_id = 754d2728b4e549c5a16c0180fcacb800
insert_flow_count = 5
insert_assetids=DfQWLAOw,F80D69
*/
const (
	InsertFlowVehicleId = "insert_flow_vehicle_id"
	InsertFlowCount     = "insert_flow_count"
	InsertAssetids      = "insert_assetids"
)

func main() {
	configMap := tool.InitConfig("conf.txt")
	insertVehicleId := configMap[InsertFlowVehicleId]
	insertFlowCount := configMap[InsertFlowCount]
	InsertAssetids := configMap[InsertAssetids]
	defaultVehicleFlowCount, _ := strconv.Atoi(insertFlowCount)

	emqx := emq_service.NewEmqx()
	emqx.Publish(insertVehicleId, creatFlowProtobuf(insertVehicleId, InsertAssetids, defaultVehicleFlowCount))
}

func creatFlowProtobuf(vehicleId string, InsertAssetids string, flowCount int) []byte {

	fAssetids := strings.Split(InsertAssetids, ",")
	pushReq := &protobuf.GWResult{
		ActionType: protobuf.GWResult_FLOWSTAT,
		GUID:       vehicleId,
	}
	flowParams := &protobuf.FlowParam{}
	//添加ThreatItem

	list := []*protobuf.FlowParam_FMacItems{}

	protocols := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	}

	for _, assetId := range fAssetids {
		flows := []*protobuf.FItem{}
		for i := 0; i < flowCount; i++ {
			moduleItem := &protobuf.FItem{
				Hash:             uint32(tool.RandomNumber(5)),
				SrcIp:            131,
				SrcPort:          23,
				DstIp:            23,
				DstPort:          23,
				Protocol:         protobuf.FlowProtos(protocols[tool.RandOneToMaxNumberT(2)]),
				FlowInfo:         "wklejl",
				SafeType:         protobuf.FlowSafetype(33),
				SafeInfo:         "jwek",
				StartTime:        tool.TimeNowToUnix(),
				LastSeenTime:     tool.TimeNowToUnix(),
				Src2DstBytes:     uint64(tool.RandomNumber(6)),
				Dst2SrcBytes:     uint64(tool.RandomNumber(6)),
				Src2DstPackets:   323232,
				Dst2SrcPackets:   200,
				FlowStat:         protobuf.FlowStat_FST_FINISH,
				HostName:         "",
				HasPassive:       true,
				IatFlowAvg:       3.2,
				IatFlowStddev:    2.3,
				DataRatio:        1.2,
				StrDataRatio:     protobuf.FItem_DR_DOWNLOAD,
				PktlenCToSAvg:    100.2,
				PktlenCToSStddev: 233.3,
				PktlenSToCAvg:    1.333,
				PktlenSToCStddev: 3.43,
				TlsClientInfo:    "gggg",
				Ja3C:             "23_232,kk23",
			}

			flows = append(flows, moduleItem)
		}

		fmacItem := &protobuf.FlowParam_FMacItems{
			Mac:      assetId,
			FlowItem: flows,
		}
		list = append(list, fmacItem)

	}

	flowParams.MacItems = list

	deviceParamsBytes, _ := proto.Marshal(flowParams)
	pushReq.Param = deviceParamsBytes
	ret, _ := proto.Marshal(pushReq)
	return ret
}
