package model_helper

import (
	"encoding/json"
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
func GetFprintProtoFlow(fp model.Fprint) map[string]float64 {
	protoByteMap := map[string]float64{}
	protoFlows := fp.CollectProtoFlows
	err := json.Unmarshal([]byte(protoFlows), &protoByteMap)
	if err != nil {
		return protoByteMap
	}
	return protoByteMap
}

/**
获取某fprint
*/
func GetAssetCateStd() []model.Fprint {
	var fplist []model.Fprint
	//标签列表
	assetFps := []*model.AssetFprint{}
	assetFpModelBase := model_base.ModelBaseImpl(&model.AssetFprint{})
	err := assetFpModelBase.GetModelListByCondition(&assetFps, "", []interface{}{}...)
	if err != nil {
		return fplist
	}
	if len(assetFps) <= 0 {
		return fplist
	}
	for _, assetFp := range assetFps {
		//todo
		fprint := GetAssetFp(assetFp.AssetId)
		fplist = append(fplist, fprint)
	}
	return fplist
}

/**
获取某fprint
*/
func GetAssetFp(assetId string) model.Fprint {
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
	return *fp
}

/**
main_proto_weight = 0.1
protos_kind_weight = 0.1
hostname_weight = 0.4
mac_weight = 0.2
type_weight = 0.1
tls_weight = 0.1

min_rate_weight=0.5
*/
func GetAssetCateMark(assetId string) map[string]float64 {
	const Hundred = 100

	mainProtoWeight := conf.MainProtoWeight
	protosKindWeight := conf.ProtosKindWeight
	hostnameWeight := conf.HostnameWeight
	tlsWeight := conf.TlsWeight
	macWeight := conf.MacWeight
	typeWeight := conf.TypeWeight

	//MinRateWeight := conf.MinRateWeight

	//获取需要识别属性的资产
	fp := GetAssetFp(assetId)
	//或者指纹库资产
	assetCateList := GetAssetCateStd()

	assetMarkMap := map[string]float64{}

	for _, assetCate := range assetCateList {
		//1.基础分{"DOWN_PPSTREAM":0.2404,"DOWN_RSYNC":0.2531,"UP_NEST_LOG_SINK":0.13311,"UP_PPSTREAM":0.12605,"UP_SSDP":0.12194}
		fpProtoFlowMap := GetFprintProtoFlow(fp)
		//贴标签的资产分数
		stdFpProtoFlowMap := GetFprintProtoFlow(assetCate)

		var stdCateMark float64
		var assetCateMark float64
		for stdFlow, stdValue := range stdFpProtoFlowMap {
			stdValue2 := stdValue * stdValue * Hundred
			stdCateMark += stdValue2

			for fpFlow, fpValue := range fpProtoFlowMap {

				if fpFlow == stdFlow {
					fpValue2 := fpValue * stdValue * Hundred
					assetCateMark += fpValue2
				}
			}
		}

		logger.Logger.Print("%s stdAssetId:%s,stdCateMark:%f,,,assetId:%s,assetCateMark:%f", util.RunFuncName(), assetCate.AssetId, stdCateMark, assetId, assetCateMark)
		logger.Logger.Info("%s stdAssetId:%s,stdCateMark:%f,,,assetId:%s,assetCateMark:%f", util.RunFuncName(), assetCate.AssetId, stdCateMark, assetId, assetCateMark)

		//2、得到标准指纹，占比最多的协议
		var weightRate float64

		var fprotoKindRate float64   //协议种类的权重
		var fmainProtoWeight float64 //协议种类占比最多的权重
		var fmacWeight float64       //mac权重
		var fhostnameWeight float64  //host权重
		var ftlsWeight float64       //tls权重
		var ftypeWeight float64      //tls权重

		stdProtoKinds := []string{}
		fpProtoKinds := []string{}

		//3、主协议是否相同
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
				fmainProtoWeight = mainProtoWeight
				weightRate += fmainProtoWeight
			}

			if util.IsExistInSlice(k, stdProtoKinds) {
				fpProtoKinds = append(fpProtoKinds, k)
			}
		}
		//协议种类占比
		if len(stdProtoKinds) == 0 {
			fprotoKindRate = 0
			weightRate += fprotoKindRate
		} else {
			fprotoKindRate = float64(len(fpProtoKinds)) / float64(len(stdProtoKinds)) * protosKindWeight
			weightRate += fprotoKindRate
		}

		logger.Logger.Print("%s stdAssetId:%s,stdCateProtoKinds:%+v,,,assetId:%s,fpProtoKinds:%v,fmainProtoWeight:%f,fprotoKindRate:%f",
			util.RunFuncName(), assetCate.AssetId, stdFpProtoFlowMap, assetId, fpProtoFlowMap, fmainProtoWeight, fprotoKindRate)

		logger.Logger.Info("%s stdAssetId:%s,stdCateProtoKinds:%+v,,,assetId:%s,fpProtoKinds:%v,fmainProtoWeight:%f,fprotoKindRate:%f",
			util.RunFuncName(), assetCate.AssetId, stdFpProtoFlowMap, assetId, fpProtoFlowMap, fmainProtoWeight, fprotoKindRate)

		//相同的mac地址，厂商
		fpTradeMark := mac.GetOrgByMAC(fp.AssetId)
		assetCateStdTradeMark := mac.GetOrgByMAC(assetCate.AssetId)

		if util.RrgsTrim(assetCateStdTradeMark) == util.RrgsTrim(fpTradeMark) {
			fmacWeight = macWeight
			weightRate += fmacWeight
		}

		logger.Logger.Print("%s stdAssetId:%s,assetCateStdTradeMark:%s,,,assetId:%s,fpTradeMark:%s,fmacWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetCateStdTradeMark, assetId, fpTradeMark, fmacWeight)

		logger.Logger.Info("%s stdAssetId:%s,assetCateStdTradeMark:%s,,,assetId:%s,fpTradeMark:%s,fmacWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetCateStdTradeMark, assetId, fpTradeMark, fmacWeight)

		//相同的hostname
		fpHost := fp.CollectHost
		assetCateStdHost := assetCate.CollectHost

		var fpHostslice []string
		_ = json.Unmarshal([]byte(fpHost), &fpHostslice)

		var assetCateStdHostslice []string
		_ = json.Unmarshal([]byte(assetCateStdHost), &assetCateStdHostslice)

		hostCommonMap := []string{}
		for _, stdhost := range assetCateStdHostslice {
			for _, host := range fpHostslice {
				if host == stdhost {
					hostCommonMap = append(hostCommonMap, host)
				}
			}
		}

		if len(assetCateStdHostslice) == 0 {
			fhostnameWeight = 0
			if len(hostCommonMap) == 0 {
				fhostnameWeight = hostnameWeight
			}
			weightRate += fhostnameWeight
		} else {
			hostRateF := float64(len(hostCommonMap)) / float64(len(assetCateStdHostslice))
			fhostnameWeight = hostRateF * hostnameWeight
			weightRate += fhostnameWeight
		}

		logger.Logger.Info("%s stdAssetId:%s,assetCateStdHostslice:%+v,,,assetId:%s,fpHostslice:%+v,fhostnameWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetCateStdHostslice, assetId, fpHostslice, fhostnameWeight)

		logger.Logger.Print("%s stdAssetId:%s,assetCateStdHostslice:%+v,,,assetId:%s,fpHostslice:%+v,fhostnameWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetCateStdHostslice, assetId, fpHostslice, fhostnameWeight)
		//相同分类
		fpCategorys := fp.Categorys
		assetCateStdCategorys := assetCate.Categorys

		var fpCategorySlice []uint32
		_ = json.Unmarshal([]byte(fpCategorys), &fpCategorySlice)

		var assetCateStdCategorySlice []uint32
		_ = json.Unmarshal([]byte(assetCateStdCategorys), &assetCateStdCategorySlice)

		commonCategoryMap := []uint32{}
		for _, category := range assetCateStdCategorySlice {
			for _, fc := range fpCategorySlice {
				if fc == category {
					commonCategoryMap = append(commonCategoryMap, fc)
				}
			}
		}

		if len(assetCateStdCategorys) == 0 {
			ftypeWeight = 0
			if len(commonCategoryMap) == 0 {
				ftypeWeight = typeWeight
			}
			weightRate += ftypeWeight
		} else {
			categoryRateF := float64(len(commonCategoryMap)) / float64(len(assetCateStdCategorySlice))
			ftypeWeight = categoryRateF * typeWeight
			weightRate += ftypeWeight
		}

		logger.Logger.Info("%s stdAssetId:%s,assetCateStdCategorySlice:%+v,,,assetId:%s,fpCategorySlice:%+v,fcategoryWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetCateStdCategorySlice, assetId, fpCategorySlice, ftypeWeight)

		logger.Logger.Print("%s stdAssetId:%s,assetCateStdCategorySlice:%+v,,,assetId:%s,fpCategorySlice:%+v,fcategoryWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetCateStdCategorySlice, assetId, fpCategorySlice, ftypeWeight)

		//相同的clienname
		fpTls := fp.CollectTls
		assetCateStdTls := assetCate.CollectTls

		var fpTlslice []string
		_ = json.Unmarshal([]byte(fpTls), &fpTlslice)

		var assetCateStdTlslice []string
		_ = json.Unmarshal([]byte(assetCateStdTls), &assetCateStdTlslice)

		commonMap := []string{}
		for _, stdtls := range assetCateStdTlslice {
			for _, tl := range fpTlslice {
				if tl == stdtls {
					commonMap = append(commonMap, tl)
				}
			}
		}

		if len(assetCateStdTlslice) == 0 {
			ftlsWeight = 0
			if len(commonMap) == 0 {
				fhostnameWeight = hostnameWeight
			}
			weightRate += ftlsWeight
		} else {
			tlsRateF := float64(len(commonMap)) / float64(len(assetCateStdTlslice))
			ftlsWeight = tlsRateF * tlsWeight
			weightRate += ftlsWeight
		}

		logger.Logger.Info("%s stdAssetId:%s,assetCateStdTlslice:%+v,,,assetId:%s,fpTlslice:%+v,ftlsWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetCateStdTlslice, assetId, fpTlslice, ftlsWeight)

		logger.Logger.Print("%s stdAssetId:%s,assetCateStdTlslice:%+v,,,assetId:%s,fpTlslice:%+v,ftlsWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetCateStdTlslice, assetId, fpTlslice, ftlsWeight)

		//相同分类

		assetCateMark = assetCateMark * (1 + weightRate)

		stdCateMark = stdCateMark * (1 + 1)

		ret := float64(assetCateMark) / float64(stdCateMark)

		logger.Logger.Print("%s stdAssetId:%s,assetId:%s,fprotoKindRate:%f,fmainProtoWeight:%f,fmacWeight:%f,fhostnameWeight:%f,ftlsWeight:%f,ftotalWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetId, fprotoKindRate, fmainProtoWeight, fmacWeight, fhostnameWeight, ftlsWeight, ret)

		logger.Logger.Info("%s stdAssetId:%s,assetId:%s,fprotoKindRate:%f,fmainProtoWeight:%f,fmacWeight:%f,fhostnameWeight:%f,ftlsWeight:%f,ftotalWeight:%f",
			util.RunFuncName(), assetCate.AssetId, assetId, fprotoKindRate, fmainProtoWeight, fmacWeight, fhostnameWeight, ftlsWeight, ret)

		assetMarkMap[assetCate.AssetId] = ret
	}
	return assetMarkMap
}

/******************************************************************************************
资产类别识别
*/
func JudgeAssetCate(assetId string) (string, float64) {
	MinRateWeight := conf.MinRateWeight //0.5

	var cateId string

	//map[string]float64
	assetCateMarkMap := GetAssetCateMark(assetId)

	logger.Logger.Print("%s assetId:%s,assetCateMarkMap:%+v", util.RunFuncName(), assetId, assetCateMarkMap)
	logger.Logger.Info("%s assetId:%s,assetCateMarkMap:%+v", util.RunFuncName(), assetId, assetCateMarkMap)

	//寻找最大值
	maxAssetIdKey := ""
	for assetId, value := range assetCateMarkMap {
		maxAssetIdKey = assetId
		for tmpAssetId, tmpValue := range assetCateMarkMap {
			if tmpValue > value {
				value = tmpValue
				maxAssetIdKey = tmpAssetId

			}
		}
		break
	}

	//判断是否5成
	maxAssetIdValue := assetCateMarkMap[maxAssetIdKey]
	if maxAssetIdValue < MinRateWeight {
		maxAssetIdKey = ""
		maxAssetIdValue = 0
	}

	assetPrint := &model.AssetFprint{
		AssetId: maxAssetIdKey,
	}

	assetPrintModelBase := model_base.ModelBaseImpl(assetPrint)

	err, recordNotFound := assetPrintModelBase.GetModelByCondition("asset_id = ?", []interface{}{assetPrint.AssetId}...)

	if err != nil {
		//todo
	}
	if recordNotFound {
		//todo
	}

	cateId = assetPrint.CateId

	return cateId, maxAssetIdValue
}
