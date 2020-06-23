package model_helper

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle/mac"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

/**
获取指纹的流量/协议 标准值
{"DOWN_PPSTREAM":0.2404,"DOWN_RSYNC":0.2531,"UP_NEST_LOG_SINK":0.13311,"UP_PPSTREAM":0.12605,"UP_SSDP":0.12194}
*/
func GetAssetCateStdMark() float64 {
	protoByteMap := GetAssetCateStdProtoFlow()
	fmt.Println("assetCateMark protoByteMap->", protoByteMap)

	var assetCateMark float64
	for _, v := range protoByteMap {
		v2 := v * v * 100
		assetCateMark += v2
	}
	fmt.Println("assetCateMark->", assetCateMark)
	return 0
}

/**
获取某fprint,某字段
*/
func GetAssetCateStdProtoFlow(fp *model.Fprint) map[string]float64 {
	protoByteMap := map[string]float64{}
	////标签列表
	//assetFps := []*model.AssetFprint{}
	//assetFpModelBase := model_base.ModelBaseImpl(&model.AssetFprint{})
	//err := assetFpModelBase.GetModelListByCondition(&assetFps, "", []interface{}{}...)
	//
	//if err != nil {
	//	return protoByteMap
	//}
	//
	//if len(assetFps) <= 0 {
	//	return protoByteMap
	//}
	//
	////目前选取第一个资产标签
	//assetFp := assetFps[0]
	//
	//fp := GetAssetFp(assetFp.AssetId)
	protoFlows := fp.CollectProtoFlows

	_ = json.Unmarshal([]byte(protoFlows), &protoByteMap)
	return protoByteMap
}

/**
获取某fprint,某字段
*/
func GetAssetCateStd() *model.Fprint {

	var fp *model.Fprint
	//标签列表
	assetFps := []*model.AssetFprint{}
	assetFpModelBase := model_base.ModelBaseImpl(&model.AssetFprint{})
	err := assetFpModelBase.GetModelListByCondition(&assetFps, "", []interface{}{}...)

	if err != nil {
		return fp
	}

	if len(assetFps) <= 0 {
		return fp
	}

	//目前选取第一个资产标签
	assetFp := assetFps[0]

	fp = GetAssetFp(assetFp.AssetId)

	return fp
}

/**
获取某fprint
*/
func GetAssetFp(assetId string) *model.Fprint {
	//获取对应的信息
	fp := &model.Fprint{
		AssetId: assetId,
	}

	fpModelBase := model_base.ModelBaseImpl(fp)

	err, recordNotFound := fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fp.AssetId}...)

	if err != nil {
		//todo
	}

	if recordNotFound {
		//todo
	}
	return fp
}

func GetAssetCateMark(assetId string) float64 {
	fp := GetAssetFp(assetId)
	AssetCateStd := GetAssetCateStd()

	fpProtoFlowMap := GetAssetCateStdProtoFlow(fp)
	stdFpProtoFlowMap := GetAssetCateStdProtoFlow(AssetCateStd)

	//基础分
	var stdCateMark float64
	var assetCateMark float64
	for stdFlow, _ := range stdFpProtoFlowMap {

		for fpFlow, value := range fpProtoFlowMap {
			v2 := value * value * 100
			stdCateMark += v2

			if fpFlow == stdFlow {
				v2 := value * value * 100
				assetCateMark += v2
			}
		}
	}
	//得到标准指纹，占比最多的协议
	var weightRate float64

	stdProtoKinds := []string{}
	fpProtoKinds := []string{}

	//主协议
	mainProtoRate := 0.1
	maxKey := ""
	for k, max := range stdFpProtoFlowMap {
		maxKey = k
		for k1, v1 := range stdFpProtoFlowMap {
			stdProtoKinds = append(stdProtoKinds, k1)
			if v1 > max {
				maxKey = k1
				max = v1
			}
		}
		break
	}

	for k, _ := range fpProtoFlowMap {
		if k == maxKey {
			weightRate += mainProtoRate
		}

		if util.IsExistInSlice(k, stdProtoKinds) {
			fpProtoKinds = append(fpProtoKinds, k)
		}
	}
	//协议种类
	protoKindRate := 0.1
	fprotoKindRate := float64(len(fpProtoKinds)) / float64(len(stdProtoKinds)) * protoKindRate
	weightRate += fprotoKindRate

	//相同的mac地址，厂商

	macRate := 0.2
	fpTradeMark := mac.GetOrgByMAC(fp.AssetId)
	assetCateStdTradeMark := mac.GetOrgByMAC(AssetCateStd.AssetId)

	if util.RrgsTrim(assetCateStdTradeMark) == util.RrgsTrim(fpTradeMark) {
		macRate = 0.2

	} else {
		macRate = 0
	}
	weightRate += macRate

	//相同的hostname
	hostNameRate := 0.4
	fpHost := fp.CollectHost
	assetCateStdHost := AssetCateStd.CollectHost

	if util.RrgsTrim(fpHost) == util.RrgsTrim(assetCateStdHost) {
		hostNameRate = 0.4

	} else {
		hostNameRate = 0
	}

	weightRate += hostNameRate

	//相同分类

	//相同的clienname
	clientInfoRate := 0.1
	fpTls := fp.CollectTls
	assetCateStdTls := AssetCateStd.CollectTls
	if util.RrgsTrim(fpTls) == util.RrgsTrim(assetCateStdTls) {
		clientInfoRate = 0.1
	} else {
		clientInfoRate = 0
	}
	weightRate += clientInfoRate

	assetCateMark = assetCateMark * (1 + weightRate)

	stdCateMark = stdCateMark * (1 + 0.5)

	ret := float64(assetCateMark) / float64(stdCateMark)

	return ret
}
