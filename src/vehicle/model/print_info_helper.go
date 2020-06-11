package model

import (
	"fmt"
	"sort"
	"strconv"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/model/model_base"
)

func GetAssetFprintProtolRate(mac string) (map[string]int, map[string]float64) {

	protos := map[string]int{}
	protoRate := map[string]float64{}

	protosFlowBytes := map[string]uint64{}
	var totalFlowBytes uint64

	fprintInfo := &FprintInfo{
		DeviceMac: mac,
	}
	fprintInfoModelBase := model_base.ModelBaseImpl(fprintInfo)

	err, fprintExist := fprintInfoModelBase.GetModelByCondition("device_mac = ?", []interface{}{fprintInfo.DeviceMac}...)

	if err != nil {
		return protos, protoRate
	}

	if fprintExist {
		return protos, protoRate
	}

	fprintInfoId := fprintInfo.FprintInfoId

	fprintPassiveInfos := []*FprintPassiveInfo{}

	fprintPassiveInfoModelBase := model_base.ModelBaseImpl(&FprintPassiveInfo{})

	err = fprintPassiveInfoModelBase.GetModelListByCondition(&fprintPassiveInfos, "fprint_info_id = ?", []interface{}{fprintInfoId}...)

	if err != nil {
		return protos, protoRate
	}
	/**
		/**
	{"HTTP":2,"KONTIKI":2,"MDNS":3,"NFS":1,"NTP":3,"SNMP":3,"SSDP":1}
	{"HTTP":2,"KONTIKI":2,"NFS":1,"NTP":3,"SSDP":1}
	{"HTTP":"0.012","KONTIKI":"0.043","MDNS":"0.459","NFS":"0.059","NTP":"0.060","SNMP":"0.264","SSDP":"0.102"}
	{"HTTP":0.012,"KONTIKI":0.043,"NFS":0.059,"NTP":0.06,"SSDP":0.102}
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
	}

	for protocol, bytes := range protosFlowBytes {
		rate := fmt.Sprintf("%.3f", float64(bytes)/float64(totalFlowBytes))

		floatRate, _ := strconv.ParseFloat(rate, 64)

		protoRate[protocol] = floatRate
	}

	fprotos := sortProtos(protos)
	fprotoRate := sortProtosRate(protoRate)

	return fprotos, fprotoRate
}

const REMAIN_MAX = 5
const REMAIN_MIN = 1

///////////////////////////////////////////////////////////////////////////////////////
func sortProtosRate(protoRate map[string]float64) map[string]float64 {
	fprotoRate := map[string]float64{}

	var ProtoRateListData ProtoRateList
	for protoId, protoCount := range protoRate {
		obj := ProtoRate{Key: protoId, Value: protoCount}
		ProtoRateListData = append(ProtoRateListData, obj)
	}

	sort.Sort(ProtoRateListData)

	var tmpProtoRateListData ProtoRateList
	if len(ProtoRateListData) <= REMAIN_MAX && len(ProtoRateListData) >= REMAIN_MIN {
		tmpProtoRateListData = ProtoRateListData[0:]
	}
	if len(ProtoRateListData) > REMAIN_MAX {
		tmpProtoRateListData = ProtoRateListData[0:REMAIN_MAX]
	}

	fmt.Println("tmpProtoRateListData::", tmpProtoRateListData)
	for _, v := range tmpProtoRateListData {
		key := v.Key
		value := v.Value
		fprotoRate[key] = value
	}

	return fprotoRate
}

type ProtoRateList []ProtoRate

type ProtoRate struct {
	Key   string
	Value float64
}

func (list ProtoRateList) Len() int {
	return len(list)
}

func (list ProtoRateList) Less(i, j int) bool {
	return list[i].Value > list[j].Value
}

func (list ProtoRateList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

/////////////////////////////////////////////////////////////////

func sortProtos(protos map[string]int) map[string]int {
	fprotos := map[string]int{}
	//排序
	var ProtoListData ProtoList
	for protoId, protoCount := range protos {
		obj := Protos{Key: protoId, Value: protoCount}
		ProtoListData = append(ProtoListData, obj)
	}

	sort.Sort(ProtoListData)
	fmt.Println("tmpProtoListData::", ProtoListData)

	var tmpProtoListData ProtoList
	if len(ProtoListData) <= REMAIN_MAX && len(ProtoListData) >= REMAIN_MIN {
		tmpProtoListData = ProtoListData[0:]
	}
	if len(ProtoListData) > REMAIN_MAX {
		tmpProtoListData = ProtoListData[0:REMAIN_MAX]
	}

	fmt.Println("tmpProtoListData::", tmpProtoListData)
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
	Value int
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
