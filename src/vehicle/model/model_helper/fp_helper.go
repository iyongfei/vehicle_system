package model_helper

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

func GetAssetCateMark(assetId string) float64 {
	var assetCateMark float64

	fp := &model.Fprint{
		AssetId: assetId,
	}

	fpModelBase := model_base.ModelBaseImpl(fp)

	err, recordNotFound := fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fp.AssetId}...)

	if err != nil || recordNotFound || util.RrgsTrimEmpty(fp.CollectProtoFlows) {
		return assetCateMark
	}

	fmt.Println("fp", fp)
	protoByteMap := map[string]float64{}
	_ = json.Unmarshal([]byte(fp.CollectProtoFlows), &protoByteMap)

	for _, v := range protoByteMap {
		v2 := v * v * 100
		fmt.Println("v2222", v2)
		assetCateMark += v2
	}
	fmt.Println(assetCateMark)
	return 0
}

//{"DOWN_PANDO":0.11356555159611187,"UP_GOOGLE_SERVICES":0.17339875743950448,
// "UP_HTTP_ACTIVESYNC":0.12406947890818859,"UP_PANDO":0.08819194387955,"UP_PPLIVE":0.18420119778354074}
