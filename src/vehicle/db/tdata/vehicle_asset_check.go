package tdata

import (
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
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

	//初始化设备离线
	assetModelBase := model_base.ModelBaseImpl(&model.Asset{})

	var assetQuery string
	var assetParam []interface{}

	if util.RrgsTrimEmpty(vehicleId) {
		assetQuery = ""
		assetParam = []interface{}{}
	} else {
		assetQuery = "vehicle_id = ?"
		assetParam = []interface{}{vehicleId}
	}

	assetAttrs := map[string]interface{}{
		"online_status": flag,
	}

	if err := assetModelBase.UpdateModelsByCondition(assetAttrs, assetQuery, assetParam...); err != nil {
		logger.Logger.Print("%s unmarshal vehicle asset init online err:%+v", util.RunFuncName(), err)
		logger.Logger.Info("%s unmarshal vehicle asset init online err:%+v", util.RunFuncName(), err)

	}
	return nil
}

func AssetFprintCheck() error {
	attrs := map[string]interface{}{
		"collect_time": gorm.Expr("collect_time + collect_end - collect_start"),
	}
	err := mysql.UpdateModelByMapModel(&model.Fprint{}, attrs, "", []interface{}{}...)
	if err != nil {
		return err
	}

	attrsNull := map[string]interface{}{
		"collect_end":   nil,
		"collect_start": nil,
	}
	err = mysql.UpdateModelByMapModel(&model.Fprint{}, attrsNull, "", []interface{}{}...)
	if err != nil {
		return err
	}

	return nil
}

func VehicleAssetFprintCheck(vehicleId string) error {
	attrs := map[string]interface{}{
		"collect_time": gorm.Expr("collect_time + collect_end - collect_start"),
	}
	err := mysql.UpdateModelByMapModel(&model.Fprint{}, attrs, "vehicle_id = ?", []interface{}{vehicleId}...)
	if err != nil {
		return err
	}

	attrsNull := map[string]interface{}{
		"collect_end":   nil,
		"collect_start": nil,
	}
	err = mysql.UpdateModelByMapModel(&model.Fprint{}, attrsNull, "",
		[]interface{}{}...)
	if err != nil {
		return err
	}

	return nil
}
