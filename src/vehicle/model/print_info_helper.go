package model

import (
	"fmt"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/model/model_base"
)

func GetAssetFprintProtolRate(mac string) (map[string]int, map[string]interface{}) {

	protos := map[string]int{}
	protoRate := map[string]interface{}{}

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
	{"HTTP":"0.012","KONTIKI":"0.043","MDNS":"0.459","NFS":"0.059","NTP":"0.060","SNMP":"0.264","SSDP":"0.102"}
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
		protoRate[protocol] = rate
	}

	return protos, protoRate
}
