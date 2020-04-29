package tdata

import (
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

//初始化离线
func VehicleAssetCheck(vehicleId string, flag bool) error {
	modelBase := model_base.ModelBaseImpl(&model.VehicleInfo{})

	var query string
	var param []interface{}

	if util.RrgsTrimEmpty(vehicleId) {
		query = ""
		param = []interface{}{}
	} else {
		query = "vehicle_id = ?"
		param = []interface{}{vehicleId}
	}

	attrs := map[string]interface{}{
		"online_status": flag,
	}

	if err := modelBase.UpdateModelsByCondition(attrs, query, param...); err != nil {
		logger.Logger.Print("%s unmarshal vehicle init online err:%+v", util.RunFuncName(), err)
		logger.Logger.Info("%s unmarshal vehicle init online err:%+v", util.RunFuncName(), err)

	}
	return nil
}
