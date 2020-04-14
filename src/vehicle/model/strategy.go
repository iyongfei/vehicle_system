package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)
///////////////////////////Strategy//////////////////////////////////////
type Strategy struct {
	gorm.Model
	StrategyId string

	Type       uint8 //策略模式
	HandleMode uint8 //处理方式
	Enable     bool  //策略启用状态

	Name      string //策略名称
	Introduce string //策略说明
}

func (strategy *Strategy) InsertModel() error {
	return mysql.CreateModel(strategy)
}
func (strategy *Strategy) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(strategy, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (strategy *Strategy) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(strategy, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (strategy *Strategy) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (strategy *Strategy) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	return nil
}

func (strategy *Strategy) CreateModel(strategyParams ...interface{}) interface{} {
	strategyParam := strategyParams[0].(*protobuf.StrategyParam)

	strategy.Type = uint8(strategyParam.GetDefenseType())
	strategy.HandleMode = uint8(strategyParam.GetHandleMode())
	strategy.Enable = strategyParam.GetEnable()
	return strategy
}
///////////////////////////StrategyVehicle//////////////////////////////////////
type StrategyVehicle struct {
	gorm.Model
	StrategyId string
	VehicleId  string
}

func (strategyVehicle *StrategyVehicle) InsertModel() error {
	return mysql.CreateModel(strategyVehicle)
}
func (strategyVehicle *StrategyVehicle) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(strategyVehicle, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (strategyVehicle *StrategyVehicle) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(strategyVehicle, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (strategyVehicle *StrategyVehicle) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (strategyVehicle *StrategyVehicle) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	return nil
}

func (strategyVehicle *StrategyVehicle) CreateModel(strategyParams ...interface{}) interface{} {
	return strategyVehicle
}

/******************************分组扩展****************************/
type StrategyGroup struct {
	gorm.Model
	StrategyId string
	GroupId    string //终端分组
}

type StrategyGroupLearningResult struct {
	gorm.Model
	GroupId          string
	LearningResultId string //id
}
