package model_helper

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
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

func GetAssetCateStdProtoFlow() map[string]float64 {
	protoByteMap := map[string]float64{}
	//标签列表
	assetFps := []*model.AssetFprint{}
	assetFpModelBase := model_base.ModelBaseImpl(&model.AssetFprint{})
	err := assetFpModelBase.GetModelListByCondition(&assetFps, "", []interface{}{}...)

	if err != nil {
		return protoByteMap
	}

	if len(assetFps) <= 0 {
		return protoByteMap
	}

	//目前选取第一个资产标签
	assetFp := assetFps[0]

	fp := GetAssetFp(assetFp.AssetId)

	_ = json.Unmarshal([]byte(fp.CollectProtoFlows), &protoByteMap)
	return protoByteMap
}

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
	//fp := GetAssetFp(assetId)
	//stdProtoRate := GetAssetCateStdProtoFlow()

	return 0
}
