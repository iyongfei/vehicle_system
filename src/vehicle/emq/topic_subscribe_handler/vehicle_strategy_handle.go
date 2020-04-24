package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleStrategy(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	strategyParam := &protobuf.StrategyParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), strategyParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicle strategy err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal vehicle strategy err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal vehicle strategy err:%s", util.RunFuncName(), err.Error())
	}

	logger.Logger.Print("%s unmarshal vehicle strategy:%+v", util.RunFuncName(), strategyParam)
	logger.Logger.Info("%s unmarshal vehicle strategy:%+v", util.RunFuncName(), strategyParam)
	//create
	strategyInfo := &model.Strategy{
		StrategyId: strategyParam.GetStrategyId(),
	}
	modelBase := model_base.ModelBaseImpl(strategyInfo)

	_, recordNotFound := modelBase.GetModelByCondition("strategy_id = ?", strategyInfo.StrategyId)

	modelBase.CreateModel(strategyParam)
	if recordNotFound {
		if err := modelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert vehicle strategy err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		attrs := map[string]interface{}{
			"type":        strategyInfo.Type,
			"handle_mode": strategyInfo.HandleMode,
			"enable":      strategyInfo.Enable,
		}
		if err := modelBase.UpdateModelsByCondition(attrs, "strategy_id = ?", strategyInfo.StrategyId); err != nil {
			return fmt.Errorf("%s update vehicle strategy err:%s", util.RunFuncName(), err.Error())
		}
	}

	//StrategyVehicle
	strategyVehicle := &model.StrategyVehicle{
		StrategyId: strategyInfo.StrategyId,
		VehicleId:  vehicleId,
	}

	strategyVehicleModelBase := model_base.ModelBaseImpl(strategyVehicle)

	_, strategyVehicleRecordNotFound := strategyVehicleModelBase.GetModelByCondition(
		"strategy_id = ? and vehicle_id = ?", []interface{}{strategyVehicle.StrategyId, strategyVehicle.VehicleId}...)
	if strategyVehicleRecordNotFound {
		if err := strategyVehicleModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert vehicle strategyVehicle err:%s", util.RunFuncName(), err.Error())
		}
	}

	return nil
}
