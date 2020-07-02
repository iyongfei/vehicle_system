package model_helper

import (
	"fmt"
	"sort"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

/****************************************************************************************************
判断某个设备采集的总流量
*/

func JudgeAssetCollectByteTotalRate(assetId string) (uint64, float64) {
	collectTotalRate := conf.CollectBytesRate //0.2
	collectTotal := conf.CollectBytes         //字节104857600

	var ftatalBytesRate float64

	totalBytes := GetAssetCollectByteTotal(assetId)

	if totalBytes > collectTotal {
		ftatalBytesRate = collectTotalRate
	} else {
		ftatalBytesRate = util.Decimal(float64(totalBytes) / float64(collectTotal) * collectTotalRate)

	}
	logger.Logger.Print("%s assetId:%s,totalBytes:%d,ftatalBytesRate:%f", util.RunFuncName(), assetId, totalBytes, ftatalBytesRate)
	logger.Logger.Info("%s assetId:%s,totalBytes:%d,ftatalBytesRate:%f", util.RunFuncName(), assetId, totalBytes, ftatalBytesRate)

	return totalBytes, ftatalBytesRate
}

func GetAssetCollectByteTotal(assetId string) uint64 {
	var totalBytes uint64

	fprintFlows := []*model.FprintFlow{}
	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ?", []interface{}{assetId}...)

	if err != nil {
		return totalBytes
	}

	for _, fprintFlow := range fprintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		totalBytes += flowByte
	}
	return totalBytes
}

/**************************************************************************************************
判断某个设备采集的tls
*/

func JudgeAssetCollectTlsInfoRate(assetId string) ([]string, float64) {
	MAX_TLS_RATE := conf.CollectTlsRate
	var ftls float64

	tls := GetAssetCollectTlsInfo(assetId)

	if len(tls) == 0 {
		ftls = 0
	} else {
		ftls = MAX_TLS_RATE
	}

	logger.Logger.Print("%s assetId:%s,tlsInfo:%+v,tlsInfoRate:%f", util.RunFuncName(), assetId, tls, ftls)
	logger.Logger.Info("%s assetId:%s,tlsInfo:%+v,tlsInfoRate:%f", util.RunFuncName(), assetId, tls, ftls)

	return tls, ftls
}
func GetAssetCollectTlsInfo(assetId string) []string {
	tls := []string{}
	fprintFlows := []*model.FprintFlow{}
	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ?", []interface{}{assetId}...)

	for _, fprintFlow := range fprintFlows {
		tlsInfo := fprintFlow.TlsClientInfo
		if tlsInfo != "" {
			if !util.IsExistInSlice(tlsInfo, tls) {
				tls = append(tls, tlsInfo)
			}

		}
	}

	if err != nil {
		return tls
	}

	return tls
}

/**************************************************************************************************
判断某个设备采集的hostname
SELECT * FROM flows WHERE vehicle_id = '';
*/

func JudgeAssetCollectHostNameRate(assetId string) ([]string, float64) {
	MAX_HOSTNAME_RATE := conf.CollectHostRate
	var fhost float64

	hostNames := GetAssetCollectHostName(assetId)
	logger.Logger.Print("%s hostName:%s", util.RunFuncName(), hostNames)
	logger.Logger.Info("%s hostName:%s", util.RunFuncName(), hostNames)

	if len(hostNames) == 0 {
		fhost = 0
	} else {
		fhost = MAX_HOSTNAME_RATE
	}

	logger.Logger.Print("%s assetId:%s,hostNames:%+v,fhost:%f", util.RunFuncName(), assetId, hostNames, fhost)
	logger.Logger.Info("%s assetId:%s,hostNames:%+v,fhost:%f", util.RunFuncName(), assetId, hostNames, fhost)
	return hostNames, fhost
}

func GetAssetCollectHostName(assetId string) []string {
	hostNames := []string{}
	fprintFlows := []*model.FprintFlow{}
	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ?", []interface{}{assetId}...)

	for _, fprintFlow := range fprintFlows {
		hostName := fprintFlow.HostName
		if hostName != "" {
			if !util.IsExistInSlice(hostName, hostNames) {
				hostNames = append(hostNames, hostName)
			}
		}
	}
	if err != nil {
		return hostNames
	}
	return hostNames
}

/**************************************************************************************************
判断某个设备采集的协议种类数
*/
func GetRankAssetCollectProtoFlow(assetId string) map[string]float64 {
	const REMAIN_MIN = 1
	fprotoBytesFloat := map[string]float64{}
	PROTOS := conf.ProtoCount
	fprotosBytesMap := GetAssetCollectProtoFlow(assetId)

	logger.Logger.Print("%s fprotosBytesMap:%v", util.RunFuncName(), fprotosBytesMap)
	logger.Logger.Info("%s fprotosBytesMap:%v", util.RunFuncName(), fprotosBytesMap)

	var protoByteFListData ProtoByteFList
	for protoId, protoByteF := range fprotosBytesMap {
		obj := ProtoByteF{Key: protoId, Value: protoByteF}
		protoByteFListData = append(protoByteFListData, obj)
	}

	sort.Sort(protoByteFListData)

	logger.Logger.Print("%s protoByteFListData:%v", util.RunFuncName(), protoByteFListData)
	logger.Logger.Info("%s protoByteFListData:%v", util.RunFuncName(), protoByteFListData)

	var tmpProtoByteFListData ProtoByteFList

	if len(protoByteFListData) <= int(PROTOS) && len(protoByteFListData) >= REMAIN_MIN {
		tmpProtoByteFListData = protoByteFListData[0:]
	}
	if len(protoByteFListData) > int(PROTOS) {
		tmpProtoByteFListData = protoByteFListData[0:int(PROTOS)]
	}

	for _, v := range tmpProtoByteFListData {
		key := v.Key
		value := v.Value
		value = util.Decimal(value)
		fprotoBytesFloat[key] = value
	}
	return fprotoBytesFloat
}

type ProtoByteFList []ProtoByteF
type ProtoByteF struct {
	Key   string
	Value float64
}

func (list ProtoByteFList) Len() int {
	return len(list)
}
func (list ProtoByteFList) Less(i, j int) bool {
	return list[i].Value > list[j].Value
}
func (list ProtoByteFList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func GetAssetCollectProtoFlow(assetId string) map[string]float64 {
	fprotosMap := map[string]float64{}
	fprintFlows := []*model.FprintFlow{}
	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ?", []interface{}{assetId}...)
	if err != nil {
		return fprotosMap
	}
	//协议->流量
	fprotosBytesMap := map[string]uint64{}
	for _, fpFlow := range fprintFlows {
		protocolStr := protobuf.GetFlowProtocols(int(fpFlow.Protocol))
		upProtocol := fmt.Sprintf("UP_%s", protocolStr)
		downProtocol := fmt.Sprintf("DOWN_%s", protocolStr)
		srcDstBytes := fpFlow.SrcDstBytes //up
		dstSrcBytes := fpFlow.DstSrcBytes //down
		if v, ok := fprotosBytesMap[upProtocol]; ok {
			fprotosBytesMap[upProtocol] = v + srcDstBytes
		} else {
			fprotosBytesMap[upProtocol] = srcDstBytes
		}
		if v, ok := fprotosBytesMap[downProtocol]; ok {
			fprotosBytesMap[downProtocol] = v + dstSrcBytes
		} else {
			fprotosBytesMap[downProtocol] = dstSrcBytes
		}
	}
	//总流量大小
	var totalBytes uint64
	for _, fprintFlow := range fprintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		totalBytes += flowByte
	}
	for p, pb := range fprotosBytesMap {
		pbRate := float64(pb) / float64(totalBytes)

		if v, ok := fprotosMap[p]; ok {
			fprotosMap[p] = pbRate + v
		} else {
			fprotosMap[p] = pbRate
		}
	}

	for p, v := range fprotosMap {
		key := p
		value := v
		value = util.Decimal(value)
		fprotosMap[key] = value
	}

	return fprotosMap
}

func JudgeAssetCollectProtoFlowRate(assetId string) (map[string]float64, float64) {
	PROTOS := conf.ProtoCount
	MAX_PROTOS_RATE := conf.ProtoCountRate

	var fcollectProto float64
	collectProtos := GetAssetCollectProtoFlow(assetId)

	if len(collectProtos) > int(PROTOS) {
		fcollectProto = MAX_PROTOS_RATE
	} else {
		fcollectProto = float64(len(collectProtos)) / float64(PROTOS) * MAX_PROTOS_RATE
	}

	logger.Logger.Print("%s assetId:%s,collectProtos:%+v,fcollectProto:%f", util.RunFuncName(), assetId, collectProtos, fcollectProto)
	logger.Logger.Info("%s assetId:%s,collectProtos:%+v,fcollectProto:%f", util.RunFuncName(), assetId, collectProtos, fcollectProto)

	return collectProtos, fcollectProto
}

/******************************************************************************************
判断某个设备采集时长是否达标
*/

func JudgeAssetCollectTimeRate(assetId string) (uint32, float64) {
	CTIME := conf.CollectTime
	MAX_COLLECT_RATE := conf.CollectTimeRate

	var fcollect_time float64

	collectTime := GetAssetCollectTime(assetId)

	//计算百分比
	if collectTime > CTIME {
		fcollect_time = MAX_COLLECT_RATE
	} else {
		fcollect_time = util.Decimal(float64(float64(collectTime)/float64(CTIME)) * MAX_COLLECT_RATE)
	}

	logger.Logger.Print("%s assetId:%s,collectTime:%d,fcollect_time:%f", util.RunFuncName(), assetId, collectTime, fcollect_time)
	logger.Logger.Info("%s assetId:%s,collectTime:%d,fcollect_time:%f", util.RunFuncName(), assetId, collectTime, fcollect_time)

	return collectTime, fcollect_time
}
func GetAssetCollectTime(assetId string) uint32 {
	var collectTime uint32

	//起始时间
	fprint := &model.Fprint{
		AssetId: assetId,
	}

	fpModelBase := model_base.ModelBaseImpl(fprint)

	err, recordNotFound := fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fprint.AssetId}...)

	if err != nil || recordNotFound {
		return collectTime
	}
	collectTime = fprint.CollectTime

	fmt.Println("GetAssetCollectTime", collectTime)
	return collectTime
}
