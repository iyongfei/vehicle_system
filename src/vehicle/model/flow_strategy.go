package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

/***************************************会话策略***************************************/
type FlowStrategy struct {
	gorm.Model
	FlowStrategyId string

	Type uint8 //策略模式
	HandleMode uint8 //处理方式
	Enable  bool//策略启用状态

	Name string//策略名称
	Introduce string//策略说明
}



func (flowStrategy *FlowStrategy) InsertModel() error {
	return mysql.CreateModel(flowStrategy)
}
func (flowStrategy *FlowStrategy) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(flowStrategy, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (flowStrategy *FlowStrategy) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(flowStrategy, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (flowStrategy *FlowStrategy) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (flowStrategy *FlowStrategy) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	return nil
}

func (flowStrategy *FlowStrategy) CreateModel(flowStrategyParams ...interface{}) interface{} {
	//flowStrategyParam := flowStrategyParams[0].(*protobuf.FlowStrategyParam)
	//
	//strategy.Type = uint8(strategyParam.GetDefenseType())
	//strategy.HandleMode = uint8(strategyParam.GetHandleMode())
	//strategy.Enable = strategyParam.GetEnable()
	return flowStrategy
}


/***********************************会话策略item**************************************/
type FlowStrategyItem struct {
	gorm.Model
	FlowStrategyItemId string
	DstIp uint32
	DstPort uint32
}

/**
表关联
 */
type FlowStrategyRelateItem struct {
	gorm.Model
	FlowStrategyId string
	FlowStrategyItemId string
}

type FlowStrategyVehicle struct {
	gorm.Model
	FlowStrategyId string
	VehicleId  string
}

/******************************分组扩展****************************/
//分组扩展
type FlowStrategyGroup struct {
	gorm.Model
	FlowStrategyId string
	GroupId    string //终端分组
}

