package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type VehicleLeader struct {
	gorm.Model//创建时间在
	LeaderId  string
	Name  string

	Phone string//手机
}


func (leader *VehicleLeader) InsertModel() error {
	return mysql.CreateModel(leader)
}
func (leader *VehicleLeader) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(leader, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}

func (leader *VehicleLeader) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(leader, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (leader *VehicleLeader) DeleModelsByCondition(query interface{}, args ...interface{}) error {


	return nil
}
func (leader *VehicleLeader) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	return nil
}

func (leader *VehicleLeader) CreateModel(flowParam ...interface{}) interface{} {

	return leader
}


