package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type VehicleAuth struct {
	gorm.Model
	VehicleId string //关联的小v ID
}

func (vehicleAuth *VehicleAuth) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(vehicleAuth, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (vehicleAuth *VehicleAuth) InsertModel() error {
	return mysql.CreateModel(vehicleAuth)
}
func (vehicleAuth *VehicleAuth) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(vehicleAuth, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (vehicleAuth *VehicleAuth) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(vehicleAuth, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (vehicleAuth *VehicleAuth) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(vehicleAuth, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (vehicleAuth *VehicleAuth) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	return nil
}

func (vehicleAuth *VehicleAuth) CreateModel(assetParams ...interface{}) interface{} {
	return vehicleAuth
}
