package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)
///////////////////////////FlowStrategy//////////////////////////////////////
type Fstrategy struct {
	gorm.Model
	FstrategyId    string

	Type          uint8 //策略模式
	HandleMode    uint8 //处理方式
	Enable        bool  //策略启用状态
}


func (flowStrategy *Fstrategy) InsertModel() error {
	return mysql.CreateModel(flowStrategy)
}
func (flowStrategy *Fstrategy) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(flowStrategy, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (flowStrategy *Fstrategy) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(flowStrategy, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (flowStrategy *Fstrategy) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(flowStrategy,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (flowStrategy *Fstrategy) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (flowStrategy *Fstrategy) CreateModel(flowStrategyParams ...interface{}) interface{} {

	strategyParam := flowStrategyParams[0].(*protobuf.FlowStrategyParam)

	flowStrategy.Type = uint8(strategyParam.GetDefenseType())
	flowStrategy.HandleMode = uint8(strategyParam.GetHandleMode())
	flowStrategy.Enable = strategyParam.GetEnable()
	return flowStrategy
}
func (flowStrategy *Fstrategy) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, query interface{}, args ...interface{})(error){

	err := mysql.QueryModelPaginationByWhereCondition(flowStrategy,pageIndex,pageSize,totalCount,paginModel,query,args...)

	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}

///////////////////////////FlowStrategyVehicle//////////////////////////////////////
type FstrategyVehicle struct {
	gorm.Model
	FstrategyVehicleId string
	FstrategyId        string
	VehicleId  		      string
}
func (flowStrategyVehicle *FstrategyVehicle) InsertModel() error {
	return mysql.CreateModel(flowStrategyVehicle)
}
func (flowStrategyVehicle *FstrategyVehicle) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(flowStrategyVehicle, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (flowStrategyVehicle *FstrategyVehicle) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(flowStrategyVehicle, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (flowStrategyVehicle *FstrategyVehicle) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(flowStrategyVehicle,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (flowStrategyVehicle *FstrategyVehicle) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (flowStrategyVehicle *FstrategyVehicle) CreateModel(strategyParams ...interface{}) interface{} {
	return flowStrategyVehicle
}
///////////////////////////VehicleLearningResult//////////////////////////////////////

type FstrategyVehicleItem struct {
	gorm.Model
	FstrategyVehicleId  string
	FstrategyItemId   string
}
func (flowStrategyVehicleItem *FstrategyVehicleItem) InsertModel() error {
	return mysql.CreateModel(flowStrategyVehicleItem)
}
func (flowStrategyVehicleItem *FstrategyVehicleItem) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(flowStrategyVehicleItem, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (flowStrategyVehicleItem *FstrategyVehicleItem) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(flowStrategyVehicleItem, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (flowStrategyVehicleItem *FstrategyVehicleItem) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (flowStrategyVehicleItem *FstrategyVehicleItem) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (flowStrategyVehicleItem *FstrategyVehicleItem) CreateModel(strategyParams ...interface{}) interface{} {
	return flowStrategyVehicleItem
}

///////////////////////////flow_strategy_items//////////////////////////////////////

type FstrategyItem struct {
	gorm.Model
	FstrategyItemId string
	VehicleId  string
	DstIp   string
	DstPort   uint32
}
func (flowStrategyItem *FstrategyItem) InsertModel() error {
	return mysql.CreateModel(flowStrategyItem)
}
func (flowStrategyItem *FstrategyItem) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(flowStrategyItem, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (flowStrategyItem *FstrategyItem) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(flowStrategyItem, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (flowStrategyItem *FstrategyItem) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (flowStrategyItem *FstrategyItem) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (flowStrategyItem *FstrategyItem) CreateModel(strategyParams ...interface{}) interface{} {
	return flowStrategyItem
}


/**
获取一条策略信息
SELECT strategies.*,strategy_vehicles.vehicle_id ,strategy_vehicle_learning_results.learning_result_id FROM strategies
inner JOIN strategy_vehicles ON strategies.strategy_id = strategy_vehicles.strategy_id
inner JOIN strategy_vehicle_learning_results ON strategy_vehicles.vehicle_id = strategy_vehicle_learning_results.vehicle_id
 */

type FlowStrategyVehicleItemJoin struct {
	gorm.Model
	FlowStrategyId     		string

	Type               		uint8
	HandleMode         		uint8
	Enable             		bool

	VehicleId 				string

	FlowStrategyVehicleId  	string  //join
	FlowStrategyItemId 		string //join
}


func GetFlowStrategyVehicleItems(query string,args ...interface{}) ([]*StrategyVehicleLearningResultJoin,error) {
	vgorm,err := mysql.GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return nil,fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	strategyVehicleLearningResultJoins := []*StrategyVehicleLearningResultJoin{}
	err = vgorm.Debug().
		Table("strategies").
		Select("strategies.*,strategy_vehicles.vehicle_id ,strategy_vehicle_learning_results.learning_result_id").
		Where(query,args...).
		Joins("inner join strategy_vehicles ON strategies.strategy_id = strategy_vehicles.strategy_id").
		Joins("inner JOIN strategy_vehicle_learning_results ON strategy_vehicles.strategy_vehicle_id = strategy_vehicle_learning_results.strategy_vehicle_id")	.
		Scan(&strategyVehicleLearningResultJoins).
		Error
	return strategyVehicleLearningResultJoins,err
}



func GetVehicleFlowStrategy(query string,args ...interface{}) (*StrategyVehicleLearningResultJoin,error) {
	vgorm,err := mysql.GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return nil,fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	strategyVehicleLearningResultJoins := &StrategyVehicleLearningResultJoin{}
	err = vgorm.Debug().
		Table("strategies").
		Select("strategies.*,strategy_vehicles.vehicle_id").
		Where(query,args...).
		Joins("inner join strategy_vehicles ON strategies.strategy_id = strategy_vehicles.strategy_id").
		Scan(strategyVehicleLearningResultJoins).
		Error
	return strategyVehicleLearningResultJoins,err
}




func GetVehicleAllFlowStrategys(query string,args ...interface{}) ([]*StrategyVehicleLearningResultJoin,error) {
	vgorm,err := mysql.GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return nil,fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	strategyVehicleLearningResultJoins := []*StrategyVehicleLearningResultJoin{}
	err = vgorm.Debug().
		Table("strategies").
		Select("strategies.*,strategy_vehicles.vehicle_id").
		Where(query,args...).
		Joins("inner join strategy_vehicles ON strategies.strategy_id = strategy_vehicles.strategy_id").
		Order("strategies.created_at desc").
		Scan(&strategyVehicleLearningResultJoins).
		Error
	return strategyVehicleLearningResultJoins,err
}

