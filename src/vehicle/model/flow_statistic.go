package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

type FlowStatistic struct {
	gorm.Model
	VehicleId     string
	InterfaceName string
	Receivex      uint64
	Uploadx       uint64
	FlowCount     uint32
	PubFlow       uint32
	NotlocalFlow  uint32
	WhiteCount    uint32
}

//序列化为数字类型
func (flowStatistic *FlowStatistic) MarshalJSON() ([]byte, error) {
	type tempType FlowStatistic
	return json.Marshal(&struct {
		CreatedAt int64
		*tempType
	}{
		CreatedAt: flowStatistic.CreatedAt.Unix(),
		tempType:  (*tempType)(flowStatistic),
	})
}

func (flowStatistic *FlowStatistic) InsertModel() error {
	return mysql.CreateModel(flowStatistic)
}
func (flowStatistic *FlowStatistic) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(flowStatistic, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (flowStatistic *FlowStatistic) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(flowStatistic, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (flowStatistic *FlowStatistic) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(flowStatistic, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (flowStatistic *FlowStatistic) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (flowStatistic *FlowStatistic) CreateModel(vehicleParam ...interface{}) interface{} {
	flowStatisticParams := vehicleParam[0].(*protobuf.FlowStatisticParam)
	flowStatistic.Receivex = flowStatisticParams.GetRx()
	flowStatistic.Uploadx = flowStatisticParams.GetTx()
	flowStatistic.FlowCount = flowStatisticParams.GetFlowCount()
	flowStatistic.PubFlow = flowStatisticParams.GetPub()
	flowStatistic.NotlocalFlow = flowStatisticParams.GetNotlocal()
	flowStatistic.WhiteCount = flowStatisticParams.GetWhite()
	return flowStatistic
}
