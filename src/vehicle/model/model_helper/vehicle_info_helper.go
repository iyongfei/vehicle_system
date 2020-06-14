package model_helper

import (
	"fmt"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

//初始化strategy
func HandleVehicleStrategyInitAction(vehicleId string) error {

	//分配AutomatedLearningResult
	automatedLearningResult := &model.AutomatedLearningResult{
		LearningResultId: util.RandomString(32),
		OriginId:         vehicleId,
		OriginType:       response.OriginTypeSelf,
	}

	automatedLResultModelBase := model_base.ModelBaseImpl(automatedLearningResult)

	err, automatedLResultRecordNotFound := automatedLResultModelBase.GetModelByCondition(
		"origin_id = ? and origin_type = ?", automatedLearningResult.OriginId, automatedLearningResult.OriginType)

	if err != nil {
		return fmt.Errorf("%s get learning result vehicleId:%s,err:%s", util.RunFuncName(), vehicleId)
	}

	if automatedLResultRecordNotFound {
		if err := automatedLResultModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert learning result err:%s", util.RunFuncName(), automatedLearningResult.OriginId, err.Error())
		}
	}

	//strategy table
	strategy := &model.Strategy{
		StrategyId: vehicleId,
		Type:       uint8(protobuf.StrategyParam_WHITEMODE),
		HandleMode: uint8(protobuf.StrategyParam_WARNING),
		Enable:     true,
	}
	strategyModelBase := model_base.ModelBaseImpl(strategy)

	err, strategyRecordNotFound := strategyModelBase.GetModelByCondition(
		"strategy_id = ?", strategy.StrategyId)

	if err != nil {
		return fmt.Errorf("%s get strategy vehicleId:%s,err:%s", util.RunFuncName(), vehicleId)
	}

	if strategyRecordNotFound {
		if err := strategyModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert strategy_id:%s, err:%s", util.RunFuncName(), strategy.StrategyId, err.Error())
		}
	}

	//strategyVehicle table
	strategyVehicle := &model.StrategyVehicle{
		StrategyVehicleId: util.RandomString(32),
		StrategyId:        strategy.StrategyId,
		VehicleId:         vehicleId,
	}

	strategyVehicleModelBase := model_base.ModelBaseImpl(strategyVehicle)

	err, strategyVehicleRecordNotFound := strategyVehicleModelBase.GetModelByCondition(
		"strategy_id = ? and vehicle_id = ?", strategyVehicle.StrategyId, strategyVehicle.VehicleId)

	if err != nil {
		return fmt.Errorf("%s get strategyVehicle vehicleId:%s,err:%s", util.RunFuncName(), vehicleId)
	}

	if strategyVehicleRecordNotFound {
		if err := strategyVehicleModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert strategyVehicle strategy_id:%s,vehicle_id:%s, err:%s",
				util.RunFuncName(), strategyVehicle.StrategyId, strategyVehicle.VehicleId, err.Error())
		}
	}

	//learningResultIds table
	strategyVehicleLearningResult := &model.StrategyVehicleLearningResult{
		StrategyVehicleId: strategyVehicle.StrategyVehicleId,
		LearningResultId:  automatedLearningResult.LearningResultId,
	}

	strategyVehicleLearningResultModelBase := model_base.ModelBaseImpl(strategyVehicleLearningResult)

	err, strategyVehicleLearningResultRecordNotFound := strategyVehicleLearningResultModelBase.GetModelByCondition(
		"strategy_vehicle_id = ? and learning_result_id = ?",
		strategyVehicleLearningResult.StrategyVehicleId, strategyVehicleLearningResult.LearningResultId)

	if err != nil {
		return fmt.Errorf("%s get strategyVehicleLearningResult strategy_vehicleId:%s,err:%s",
			util.RunFuncName(), strategyVehicleLearningResult.StrategyVehicleId, err)
	}

	if strategyVehicleLearningResultRecordNotFound {
		if err := strategyVehicleLearningResultModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert strategyVehicleLearningResult strategy_vehicleId:%s, err:%s",
				util.RunFuncName(), strategyVehicleLearningResult.StrategyVehicleId, err.Error())
		}
	}

	return nil
}
