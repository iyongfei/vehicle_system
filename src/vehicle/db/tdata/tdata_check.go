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
	return nil
}
