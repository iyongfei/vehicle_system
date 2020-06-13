package model

import (
	"sort"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/model/model_base"
)

const TOTAL = "t"

func GetAssetFprintProtolRate(mac string) (map[string]uint64, map[string]uint64) {
	var totalProtos uint64
	protos := map[string]uint64{}
	//protoRate := map[string]float64{}

	var totalFlowBytes uint64
	protosFlowBytes := map[string]uint64{}

	//获取指纹信息
	fprintInfo := &FprintInfo{
		DeviceMac: mac,
	}
	fprintInfoModelBase := model_base.ModelBaseImpl(fprintInfo)

	err, fprintExist := fprintInfoModelBase.GetModelByCondition("device_mac = ?", []interface{}{fprintInfo.DeviceMac}...)

	if err != nil {
		return protos, protosFlowBytes
	}

	if fprintExist {
		return protos, protosFlowBytes
	}

	fprintInfoId := fprintInfo.FprintInfoId

	//捕获的指纹列表
	fprintPassiveInfos := []*FprintPassiveInfo{}

	fprintPassiveInfoModelBase := model_base.ModelBaseImpl(&FprintPassiveInfo{})

	err = fprintPassiveInfoModelBase.GetModelListByCondition(&fprintPassiveInfos, "fprint_info_id = ?", []interface{}{fprintInfoId}...)

	if err != nil {
		return protos, protosFlowBytes
	}

	/**
	{"DIAMETER":1,"FREE_205":1,"NETFLOW":1,"RADIUS":1,"RX":1,"SOULSEEK":1,"SSDP":1,"STEAM":1,"SYSLOG":1,"t":15}
	{"FREE_205":17173,"NETFLOW":15120,"RADIUS":11934,"SMBV1":10380,"SSDP":10224,"STEALTHNET":12220,"STEAM":13530,"TELEGRAM":11538,"WHOIS_DAS":16362,"t":146951}
	*/
	//种类，占比
	for _, fprintPassiveInfo := range fprintPassiveInfos {
		protocol := fprintPassiveInfo.Protocol

		protocolStr := protobuf.GetFlowProtocols(int(protocol))

		if v, ok := protos[protocolStr]; ok {
			protos[protocolStr] = v + 1
		} else {
			protos[protocolStr] = 1
		}

		//每条会话的流量和
		srcDstBytes := fprintPassiveInfo.SrcDstBytes
		dstSrcBytes := fprintPassiveInfo.DstSrcBytes

		if v, ok := protosFlowBytes[protocolStr]; ok {
			protosFlowBytes[protocolStr] = v + srcDstBytes + dstSrcBytes
		} else {
			protosFlowBytes[protocolStr] = srcDstBytes + dstSrcBytes
		}
		totalFlowBytes += srcDstBytes + dstSrcBytes
		if protocolStr != "" {
			totalProtos += 1
		}

	}

	protos[TOTAL] = totalProtos
	protosFlowBytes[TOTAL] = totalFlowBytes

	fprotos := sortProtos(protos)
	fprotosFlowBytes := sortProtos(protosFlowBytes)

	return fprotos, fprotosFlowBytes
}

const REMAIN_MAX = 100
const REMAIN_MIN = 1

/////////////////////////////////////////////////////////////////

func sortProtos(protos map[string]uint64) map[string]uint64 {
	fprotos := map[string]uint64{}
	//排序
	var ProtoListData ProtoList
	for protoId, protoCount := range protos {
		obj := Protos{Key: protoId, Value: protoCount}
		ProtoListData = append(ProtoListData, obj)
	}

	sort.Sort(ProtoListData)

	var tmpProtoListData ProtoList
	if len(ProtoListData) <= REMAIN_MAX && len(ProtoListData) >= REMAIN_MIN {
		tmpProtoListData = ProtoListData[0:]
	}
	if len(ProtoListData) > REMAIN_MAX {
		tmpProtoListData = ProtoListData[0:REMAIN_MAX]
	}

	for _, v := range tmpProtoListData {
		key := v.Key
		value := v.Value
		fprotos[key] = value
	}
	return fprotos

}

type ProtoList []Protos

type Protos struct {
	Key   string
	Value uint64
}

func (list ProtoList) Len() int {
	return len(list)
}

func (list ProtoList) Less(i, j int) bool {

	return list[i].Value > list[j].Value
}

func (list ProtoList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
