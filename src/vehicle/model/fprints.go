package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type Fprint struct {
	gorm.Model

	FprintId string

	VehicleId string
	AssetId   string

	//CollectProtoRate  float64 //流量占比
	CollectProtoFlows string //流量占比
	CollectHost       string //hostname
	CollectTls        string //tls
	CollectBytes      uint64 //采集流量
	CollectTime       uint32 //采集时间

	CollectStart  uint64
	CollectEnd    uint64
	CollectFinish bool
	AutoCateId    string
}

//序列化为数字类型
func (tmp *Fprint) MarshalJSON() ([]byte, error) {
	type Type Fprint
	return json.Marshal(&struct {
		CreatedAt int64
		*Type
	}{
		CreatedAt: tmp.CreatedAt.Unix(),
		Type:      (*Type)(tmp),
	})
}

func (f *Fprint) InsertModel() error {
	return mysql.CreateModel(f)
}

func (f *Fprint) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(f, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}

func (f *Fprint) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(f, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (f *Fprint) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(f, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (f *Fprint) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (flow *Fprint) CreateModel(flowParam ...interface{}) interface{} {
	return flow
}

func (flow *Fprint) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(flow, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
