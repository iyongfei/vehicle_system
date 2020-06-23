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

func JudgeAssetCollectByteTotalRate(assetId string) float64 {
	collectTotalRate := conf.CollectTotalRate //0.2
	collectTotal := conf.CollectTotal         //字节104857600

	var ftatalBytes float64

	totalBytes := GetAssetCollectByteTotal(assetId)

	logger.Logger.Print("%s totalBytes:%d", util.RunFuncName(), totalBytes)
	logger.Logger.Info("%s totalBytes:%d", util.RunFuncName(), totalBytes)

	if totalBytes > collectTotal {
		ftatalBytes = collectTotalRate
	} else {
		ftatalBytes = float64(float64(totalBytes)/float64(collectTotal)) * collectTotalRate
	}
	logger.Logger.Print("%s ftatalBytes:%f", util.RunFuncName(), ftatalBytes)
	logger.Logger.Info("%s ftatalBytes:%f", util.RunFuncName(), ftatalBytes)

	return ftatalBytes
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

func JudgeAssetCollectTlsInfoRate(assetId string) float64 {
	MAX_TLS_RATE := conf.CollectTlsRate
	var ftls float64

	tlsInfo := GetAssetCollectTlsInfo(assetId)

	logger.Logger.Print("%s tlsInfo:%s", util.RunFuncName(), tlsInfo)
	logger.Logger.Info("%s tlsInfo:%s", util.RunFuncName(), tlsInfo)

	if tlsInfo == "" {
		ftls = 0
	} else {
		ftls = MAX_TLS_RATE
	}
	logger.Logger.Print("%s tlsInfo:%f", util.RunFuncName(), ftls)
	logger.Logger.Info("%s tlsInfo:%f", util.RunFuncName(), ftls)

	return ftls
}
func GetAssetCollectTlsInfo(assetId string) string {

	fprintFlow := &model.FprintFlow{}

	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlow, "asset_id = ? and tls_client_info != ''", []interface{}{assetId}...)

	if err != nil {
		return fprintFlow.TlsClientInfo
	}
	return fprintFlow.TlsClientInfo
}

/**************************************************************************************************
判断某个设备采集的hostname
SELECT * FROM flows WHERE vehicle_id = '';
*/

func JudgeAssetCollectHostNameRate(assetId string) float64 {
	MAX_HOSTNAME_RATE := conf.CollectHostRate
	var fhost float64

	hostName := GetAssetCollectHostName(assetId)
	logger.Logger.Print("%s hostName:%s", util.RunFuncName(), hostName)
	logger.Logger.Info("%s hostName:%s", util.RunFuncName(), hostName)

	if hostName == "" {
		fhost = 0
	} else {
		fhost = MAX_HOSTNAME_RATE
	}
	logger.Logger.Print("%s fhost:%f", util.RunFuncName(), fhost)
	logger.Logger.Info("%s fhost:%f", util.RunFuncName(), fhost)
	return fhost
}

func GetAssetCollectHostName(assetId string) string {
	fprintFlows := &model.FprintFlow{}

	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ? and host_name != ''", []interface{}{assetId}...)

	if err != nil {
		return fprintFlows.HostName
	}

	return fprintFlows.HostName
}

/**************************************************************************************************
判断某个设备采集的协议种类数
*/
func GetRankAssetCollectProtoFlow(assetId string) map[string]float64 {
	const REMAIN_MIN = 5
	fprotoBytesFloat := map[string]float64{}
	PROTOS := conf.ProtoCount
	fprotosBytesMap := GetAssetCollectProtoFlow(assetId)

	var protoByteFListData ProtoByteFList
	for protoId, protoByteF := range fprotosBytesMap {
		obj := ProtoByteF{Key: protoId, Value: protoByteF}
		protoByteFListData = append(protoByteFListData, obj)
	}

	sort.Sort(protoByteFListData)

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

////
func GetAssetCollectProtoFlow(assetId string) map[string]float64 {
	fprotosMap := map[string]float64{}
	fprintFlows := []*model.FprintFlow{}
	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ?", []interface{}{assetId}...)
	if err != nil {
		return fprotosMap
	}
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
		pbRate := float64(float64(pb) / float64(totalBytes))

		if v, ok := fprotosMap[p]; ok {
			fprotosMap[p] = pbRate + v
		} else {
			fprotosMap[p] = pbRate
		}
	}
	return fprotosMap
}

func JudgeAssetCollectProtoFlowRate(assetId string) float64 {
	PROTOS := conf.ProtoCount
	MAX_PROTOS_RATE := conf.ProtoCountRate

	var fcollectProto float64
	collectProtosRate := GetAssetCollectProtoFlow(assetId)

	logger.Logger.Print("%s collectProtosRate:%v", util.RunFuncName(), collectProtosRate)
	logger.Logger.Info("%s collectProtosRate:%v", util.RunFuncName(), collectProtosRate)

	if len(collectProtosRate) > int(PROTOS) {
		fcollectProto = MAX_PROTOS_RATE
	} else {
		fcollectProto = float64(float64(len(collectProtosRate))/float64(PROTOS)) * MAX_PROTOS_RATE
	}

	logger.Logger.Print("%s collectProtosRate:%f", util.RunFuncName(), fcollectProto)
	logger.Logger.Info("%s collectProtosRate:%f", util.RunFuncName(), fcollectProto)
	return fcollectProto
}

/******************************************************************************************
判断某个设备采集时长是否达标
*/

func JudgeAssetCollectTimeRate(assetId string) float64 {
	CTIME := conf.CollectTime
	MAX_COLLECT_RATE := conf.CollectTimeRate

	var fcollect_time float64

	collectTime := GetAssetCollectTime(assetId)

	logger.Logger.Print("%s collectTime:%d", util.RunFuncName(), collectTime)
	logger.Logger.Info("%s collectTime:%d", util.RunFuncName(), collectTime)
	//计算百分比
	if collectTime > CTIME {
		fcollect_time = MAX_COLLECT_RATE
	} else {
		fcollect_time = float64(float64(collectTime)/float64(CTIME)) * MAX_COLLECT_RATE
	}

	logger.Logger.Print("%s fcollect_time:%f", util.RunFuncName(), fcollect_time)
	logger.Logger.Info("%s fcollect_time:%f", util.RunFuncName(), fcollect_time)
	return fcollect_time
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
	endTime := fprint.CollectEnd
	ctime := fprint.CollectTime
	startTime := fprint.CollectStart

	collectTime = uint32(ctime) + uint32(endTime-startTime)

	fmt.Println(util.UnixStamp2Str(int64(endTime)))
	fmt.Println(util.UnixStamp2Str(int64(ctime)))
	fmt.Println(util.UnixStamp2Str(int64(startTime)))
	//2702494284
	return collectTime
}
