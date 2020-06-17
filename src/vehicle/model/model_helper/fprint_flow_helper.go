package model_helper

import (
	"math"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/util"
)

/**
collect_time=300
proto_count=5
collect_total=1048576
是否采集到了hostname
是否采集到了tls_client_info
*/

/**
判断某个设备采集的总流量
*/
const MinTotalRate = 0.6

func JudgeAssetCollectByteTotal(assetId string) float64 {
	const MAX_BYTESTOTAL_RATE = 0.2
	const MAX_BYTESTOTAL = 100 * 1024 * 1024 //字节

	var ftatalBytes float64
	fprintFlows := []*model.FprintFlow{}

	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ?", []interface{}{assetId}...)

	if err != nil {
		return 0
	}

	var totalBytes uint64
	for _, fprintFlow := range fprintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		totalBytes += flowByte
	}
	if totalBytes > MAX_BYTESTOTAL {
		ftatalBytes = MAX_BYTESTOTAL_RATE
	} else {
		ftatalBytes = float64(float64(totalBytes)/float64(MAX_BYTESTOTAL)) * MAX_BYTESTOTAL_RATE
	}

	return ftatalBytes
}

/**
判断某个设备采集的tls
*/

func JudgeAssetCollectTlsInfo(assetId string) float64 {
	const MAX_TLS_RATE = 0.1
	var ftls float64
	fprintFlows := []*model.FprintFlow{}

	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ? and tls_client_info != ''", []interface{}{assetId}...)

	if err != nil {
		return 0
	}

	if len(fprintFlows) == 0 {
		ftls = 0
	} else {
		ftls = MAX_TLS_RATE
	}
	return ftls
}

/**
判断某个设备采集的hostname
SELECT * FROM flows WHERE vehicle_id = '';
*/

func JudgeAssetCollectHostName(assetId string) float64 {
	const MAX_HOSTNAME_RATE = 0.2
	var fhost float64
	fprintFlows := []*model.FprintFlow{}

	err := mysql.QueryModelRecordsByWhereCondition(&fprintFlows, "asset_id = ? and host_name != ''", []interface{}{assetId}...)

	if err != nil {
		return 0
	}

	if len(fprintFlows) == 0 {
		fhost = 0
	} else {
		fhost = MAX_HOSTNAME_RATE
	}

	return fhost
}

/**
判断某个设备采集的协议种类数
*/

func JudgeAssetCollectProtos(assetId string) float64 {
	const PROTOS = 5
	const MAX_PROTOS_RATE = 0.3

	var fProtos float64
	var protocols []string
	_ = mysql.QueryPluckByModelWhere(&model.FprintFlow{}, "protocol", &protocols,
		"asset_id = ?", []interface{}{assetId}...)

	protocolMap := map[string]int{}

	for _, protocol := range protocols {
		if protocolCount, ok := protocolMap[protocol]; ok {
			protocolMap[protocol] = protocolCount + 1
		} else {
			protocolMap[protocol] = 1
		}
	}

	protosCountSlice := []string{}

	for proto, _ := range protocolMap {
		protosCountSlice = append(protosCountSlice, proto)
	}

	if len(protosCountSlice) > PROTOS {
		fProtos = MAX_PROTOS_RATE
	} else {
		fProtos = float64(float64(len(protosCountSlice))/float64(PROTOS)) * MAX_PROTOS_RATE
	}

	return fProtos
}

/**
判断某个设备采集时长是否达标
*/
func JudgeAssetCollectTime(assetId string) float64 {
	const CTIME = 300
	const MAX_COLLECT_RATE = 0.2

	var fcollect float64

	ctime := conf.CollectTime

	if ctime == 0 {
		ctime = CTIME
	}

	//起始时间
	firstFprintFlow := &model.FprintFlow{}

	err := mysql.QueryModelOneRecordByWhereSelectOrderBy(firstFprintFlow, []string{"created_at"},
		"created_at", "asset_id = ?", []interface{}{assetId}...)

	logger.Logger.Print("%s, firt fprint flow assetId:%s,createdAt:%d", util.RunFuncName(), assetId, firstFprintFlow.CreatedAt.Unix())

	if err != nil {
		logger.Logger.Error("%s, first fprint flow assetId:%s,err:%+v", util.RunFuncName(), assetId, err)
		logger.Logger.Print("%s, first fprint flow assetId:%s,err:%+v", util.RunFuncName(), assetId, err)
		return 0
	}
	//截止时间
	lastFprintFlow := &model.FprintFlow{}
	err = mysql.QueryModelOneRecordByWhereSelectOrderBy(lastFprintFlow, []string{"created_at"},
		"created_at desc", "asset_id = ?", []interface{}{assetId}...)
	if err != nil {
		logger.Logger.Error("%s, last fprint flow assetId:%s,err:%+v", util.RunFuncName(), assetId, err)
		logger.Logger.Print("%s, last fprint flow assetId:%s,err:%+v", util.RunFuncName(), assetId, err)
		return 0
	}

	logger.Logger.Print("%s, last fprint flow assetId:%s,createdAt:%d", util.RunFuncName(), assetId, lastFprintFlow.CreatedAt.Unix())

	//计算起始时间绝对值
	firstAbsTime := math.Abs(float64(firstFprintFlow.CreatedAt.Unix()))
	lastAbsTime := math.Abs(float64(lastFprintFlow.CreatedAt.Unix()))

	//差值
	distanceTime := uint32(math.Abs(lastAbsTime - firstAbsTime))

	logger.Logger.Print("%s, last fprint flow assetId:%s,distanceTime:%d", util.RunFuncName(), assetId, distanceTime)

	//计算百分比
	if distanceTime > ctime {
		fcollect = MAX_COLLECT_RATE
	} else {
		fcollect = float64(float64(distanceTime)/float64(ctime)) * MAX_COLLECT_RATE
	}
	return fcollect
}
