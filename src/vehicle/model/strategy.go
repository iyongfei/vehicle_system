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
	StrategyId    string

	Type          uint8 //策略模式
	HandleMode    uint8 //处理方式
	Enable        bool  //策略启用状态
}


func (strategy *Strategy) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, query interface{}, args ...interface{})(error){

	err := mysql.QueryModelPaginationByWhereCondition(strategy,pageIndex,pageSize,totalCount,paginModel,query,args...)

	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
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
	err := mysql.HardDeleteModelB(strategy,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (strategy *Strategy) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
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
	err := mysql.HardDeleteModelB(strategyVehicle,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (strategyVehicle *StrategyVehicle) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}

func (strategyVehicle *StrategyVehicle) CreateModel(strategyParams ...interface{}) interface{} {
	return strategyVehicle
}
///////////////////////////StrategyVehicleLearningResult//////////////////////////////////////

type StrategyVehicleLearningResult struct {
	gorm.Model
	VehicleId          string
	LearningResultId string
}


func (strategyVehicleLearningResult *StrategyVehicleLearningResult) InsertModel() error {
	return mysql.CreateModel(strategyVehicleLearningResult)
}
func (strategyVehicleLearningResult *StrategyVehicleLearningResult) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(strategyVehicleLearningResult, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (strategyVehicleLearningResult *StrategyVehicleLearningResult) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(strategyVehicleLearningResult, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (strategyVehicleLearningResult *StrategyVehicleLearningResult) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (strategyVehicleLearningResult *StrategyVehicleLearningResult) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}

func (strategyVehicleLearningResult *StrategyVehicleLearningResult) CreateModel(strategyParams ...interface{}) interface{} {
	return strategyVehicleLearningResult
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
