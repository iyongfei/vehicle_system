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

func HandleVehicleFlowStrategy(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	flowStrategyParam := &protobuf.FlowStrategyParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), flowStrategyParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicle flow strategy err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal vehicle flow strategy err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal vehicle flow strategy err:%s", util.RunFuncName(), err.Error())
	}
	logger.Logger.Print("%s handle_fstrategy:%+v", util.RunFuncName(), *flowStrategyParam)
	logger.Logger.Info("%s handle_fstrategy:%+v", util.RunFuncName(), *flowStrategyParam)

	//create
	flowStrategyInfo := &model.Fstrategy{
		FstrategyId: flowStrategyParam.GetFlowStrategyId(),
	}
	modelBase := model_base.ModelBaseImpl(flowStrategyInfo)

	_, recordNotFound := modelBase.GetModelByCondition("fstrategy_id = ?", flowStrategyInfo.FstrategyId)

	modelBase.CreateModel(flowStrategyParam)
	if recordNotFound {
		if err := modelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert vehicle flow strategy err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		attrs := map[string]interface{}{
			"type":        flowStrategyInfo.Type,
			"handle_mode": flowStrategyInfo.HandleMode,
			//"enable":      flowStrategyInfo.Enable,
		}
		if err := modelBase.UpdateModelsByCondition(attrs, "fstrategy_id = ?", flowStrategyInfo.FstrategyId); err != nil {
			return fmt.Errorf("%s update vehicle flow strategy err:%s", util.RunFuncName(), err.Error())
		}
	}

	//StrategyVehicle
	flowStrategyVehicle := &model.FstrategyVehicle{
		FstrategyId: flowStrategyInfo.FstrategyId,
		VehicleId:   vehicleId,
	}

	flowStrategyVehicleModelBase := model_base.InsertModelImpl(flowStrategyVehicle)
	flowStrategyVehicleGetModelBase := model_base.GetModelImpl(flowStrategyVehicle)

	_, strategyVehicleRecordNotFound := flowStrategyVehicleGetModelBase.GetModelByCondition(
		"fstrategy_id = ? and vehicle_id = ?", []interface{}{flowStrategyVehicle.FstrategyId, flowStrategyVehicle.VehicleId}...)
	if strategyVehicleRecordNotFound {
		if err := flowStrategyVehicleModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert vehicle flowStrategyVehicle err:%s", util.RunFuncName(), err.Error())
		}
	}

	return nil
}
