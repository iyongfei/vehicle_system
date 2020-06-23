package model_helper

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/mac"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

/**
获取指纹的流量/协议 标准值
{"DOWN_PPSTREAM":0.2404,"DOWN_RSYNC":0.2531,"UP_NEST_LOG_SINK":0.13311,"UP_PPSTREAM":0.12605,"UP_SSDP":0.12194}
*/
func GetFprintProtoFlow(fp *model.Fprint) map[string]float64 {
	protoByteMap := map[string]float64{}
	protoFlows := fp.CollectProtoFlows
	_ = json.Unmarshal([]byte(protoFlows), &protoByteMap)
	return protoByteMap
}

/**
获取某fprint
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
	const Hundred = 100

	mainProtoWeight := conf.MainProtoWeight
	protosKindWeight := conf.ProtosKindWeight
	hostnameWeight := conf.HostnameWeight
	macWeight := conf.MacWeight
	//typeWeight := conf.TypeWeight
	tlsWeight := conf.TlsWeight
	MinRateWeight := conf.MinRateWeight

	//获取需要识别属性的资产
	fp := GetAssetFp(assetId)
	//或者指纹库资产
	AssetCateStd := GetAssetCateStd()

	//1.基础分
	fpProtoFlowMap := GetFprintProtoFlow(fp)
	stdFpProtoFlowMap := GetFprintProtoFlow(AssetCateStd)

	logger.Logger.Print("%s fpProtoFlowMap:%+v", util.RunFuncName(), fpProtoFlowMap)
	logger.Logger.Info("%s fpProtoFlowMap:%+v", util.RunFuncName(), fpProtoFlowMap)

	logger.Logger.Print("%s stdFpProtoFlowMap:%+v", util.RunFuncName(), stdFpProtoFlowMap)
	logger.Logger.Info("%s stdFpProtoFlowMap:%+v", util.RunFuncName(), stdFpProtoFlowMap)

	var stdCateMark float64
	var assetCateMark float64
	for stdFlow, stdValue := range stdFpProtoFlowMap {
		stdValue2 := stdValue * stdValue * Hundred
		stdCateMark += stdValue2

		for fpFlow, fpValue := range fpProtoFlowMap {

			if fpFlow == stdFlow {
				fpValue2 := fpValue * stdValue * Hundred
				fmt.Println("value:", fpValue2)
				assetCateMark += fpValue2
			}
		}
	}

	logger.Logger.Print("%s stdCateMark:%f", util.RunFuncName(), stdCateMark)
	logger.Logger.Info("%s stdCateMark:%f", util.RunFuncName(), stdCateMark)

	logger.Logger.Print("%s assetCateMark:%f", util.RunFuncName(), assetCateMark)
	logger.Logger.Info("%s assetCateMark:%f", util.RunFuncName(), assetCateMark)

	//得到标准指纹，占比最多的协议
	var weightRate float64

	stdProtoKinds := []string{}
	fpProtoKinds := []string{}

	//主协议
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
			weightRate += mainProtoWeight
		}

		if util.IsExistInSlice(k, stdProtoKinds) {
			fpProtoKinds = append(fpProtoKinds, k)
		}
	}
	logger.Logger.Print("%s stdProtoKinds:%v,fpProtoKinds:%v", util.RunFuncName(), stdProtoKinds, fpProtoKinds)
	logger.Logger.Info("%s stdProtoKinds:%v,fpProtoKinds:%v", util.RunFuncName(), stdProtoKinds, fpProtoKinds)

	//协议种类
	fprotoKindRate := float64(len(fpProtoKinds)) / float64(len(stdProtoKinds)) * protosKindWeight
	logger.Logger.Print("%s fprotoKindRate:%f", util.RunFuncName(), fprotoKindRate)
	logger.Logger.Info("%s fprotoKindRate:%f", util.RunFuncName(), fprotoKindRate)
	weightRate += fprotoKindRate

	//相同的mac地址，厂商
	fpTradeMark := mac.GetOrgByMAC(fp.AssetId)
	assetCateStdTradeMark := mac.GetOrgByMAC(AssetCateStd.AssetId)

	logger.Logger.Print("%s assetCateStdTradeMark:%s", util.RunFuncName(), assetCateStdTradeMark)
	logger.Logger.Info("%s assetCateStdTradeMark:%s", util.RunFuncName(), assetCateStdTradeMark)

	if util.RrgsTrim(assetCateStdTradeMark) == util.RrgsTrim(fpTradeMark) {
		weightRate += macWeight
	}

	//相同的hostname
	fpHost := fp.CollectHost
	assetCateStdHost := AssetCateStd.CollectHost

	if util.RrgsTrim(fpHost) == util.RrgsTrim(assetCateStdHost) {
		weightRate += hostnameWeight
	}
	logger.Logger.Print("%s assetCateStdHost:%s", util.RunFuncName(), assetCateStdHost)
	logger.Logger.Info("%s assetCateStdHost:%s", util.RunFuncName(), assetCateStdHost)

	//相同分类

	//相同的clienname
	fpTls := fp.CollectTls
	assetCateStdTls := AssetCateStd.CollectTls
	if util.RrgsTrim(fpTls) == util.RrgsTrim(assetCateStdTls) {
		weightRate += tlsWeight
	}
	assetCateMark = assetCateMark * (1 + weightRate)

	stdCateMark = stdCateMark * (1 + MinRateWeight)

	ret := float64(assetCateMark) / float64(stdCateMark)

	return ret
}
