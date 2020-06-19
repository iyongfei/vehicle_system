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

/**
初始化授权列表
*/
func TdataVehicleAuthCheck() error {
	vehicleAuths := GetVehicleAuths()
	for _, vehicleId := range vehicleAuths {
		//vgw-1214-pwd-key
		vehicleIdMd5 := util.Md5(vehicleId + response.PasswordSecret)

		vehicleAuth := &model.VehicleAuth{
			VehicleId: vehicleIdMd5,
		}

		modelBase := model_base.ModelBaseImpl(vehicleAuth)

		err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleAuth.VehicleId}...)

		if err != nil {
			continue
		}
		//不存在插入
		if recordNotFound {
			err := modelBase.InsertModel()
			if err != nil {
				continue
			}
		}
	}

	return nil
}

func GetVehicleAuths() []string {
	vehicleAuths := []string{
		"1",
		"2",
		"3",
	}
	return vehicleAuths
}
