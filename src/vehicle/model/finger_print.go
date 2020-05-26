package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type FingerPrint struct {
	gorm.Model
	FprintId    string
	CateId      string
	VehicleId   string
	DeviceMac   string
	FlowIds     string
	ProtoRate   string
	CollectType uint8
}

func (fingerPrint *FingerPrint) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(fingerPrint, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fingerPrint *FingerPrint) InsertModel() error {
	return mysql.CreateModel(fingerPrint)
}
func (fingerPrint *FingerPrint) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(fingerPrint, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (fingerPrint *FingerPrint) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(fingerPrint, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (fingerPrint *FingerPrint) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(fingerPrint, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fingerPrint *FingerPrint) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fingerPrint *FingerPrint) CreateModel(assetParams ...interface{}) interface{} {
	return fingerPrint
}
