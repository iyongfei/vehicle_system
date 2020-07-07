package tdata

import (
	"fmt"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

/**
初始化默认分组
*/
func TdataCheck() error {
	areaGroup := &model.AreaGroup{
		AreaName:       response.UnGroupName,
		AreaCode:       util.RandomString(32),
		ParentAreaCode: "",
		TreeAreaCode:   "",
	}

	modelBase := model_base.ModelBaseImpl(areaGroup)

	_, recordNotFound := modelBase.GetModelByCondition("area_name = ?",
		[]interface{}{areaGroup.AreaName}...)
	if recordNotFound {
		err := modelBase.InsertModel()
		if err != nil {
			return fmt.Errorf("%s insert ungroup err:%s", err)
		}
	}
	//指纹库类别
	cate := &model.Category{
		Name:   response.Vc,
		CateId: util.RandomString(32),
	}

	cateModelBase := model_base.ModelBaseImpl(cate)

	_, cateRecordNotFound := cateModelBase.GetModelByCondition("name = ?",
		[]interface{}{cate.Name}...)
	if cateRecordNotFound {
		err := cateModelBase.InsertModel()
		if err != nil {
			return fmt.Errorf("%s insert cate err:%s", err)
		}
	}

	//初始化资产默认分组
	assetGroup := &model.AssetGroup{
		AreaName:       response.UnGroupName,
		AreaCode:       util.RandomString(32),
		ParentAreaCode: "",
		TreeAreaCode:   "",
	}

	assetGroupModelBase := model_base.ModelBaseImpl(assetGroup)

	_, assetGroupRecordNotFound := assetGroupModelBase.GetModelByCondition("area_name = ?",
		[]interface{}{assetGroup.AreaName}...)
	if assetGroupRecordNotFound {
		err := assetGroupModelBase.InsertModel()
		if err != nil {
			return fmt.Errorf("%s insert asset ungroup err:%s", err)
		}
	}

	return nil
}
