package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/service/push"
	"vehicle_system/src/vehicle/util"
)

func HandleFlowStatisticInfo(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	flowStatisticParam := &protobuf.FlowStatisticParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), flowStatisticParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal flowStatisticInfoParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal flowStatisticInfoParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal flowStatisticInfoParam err:%s", util.RunFuncName(), err.Error())
	}
	//vehicleId
	logger.Logger.Print("%s unmarshal flowStatisticInfoParam:%+v", util.RunFuncName(), flowStatisticParam)
	logger.Logger.Info("%s unmarshal flowStatisticInfoParam:%+v", util.RunFuncName(), flowStatisticParam)
	//create
	flowStatistic := &model.FlowStatistic{
		VehicleId:     vehicleId,
		InterfaceName: flowStatisticParam.GetInterfaceName(),
	}

	modelBase := model_base.ModelBaseImpl(flowStatistic)
	_, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ? and interface_name = ?",
		[]interface{}{flowStatistic.VehicleId, flowStatistic.InterfaceName}...)
	modelBase.CreateModel(flowStatisticParam)
	if recordNotFound {
		if err := modelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert flowStatisticParam err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		//更新 排除VehicleId,Name,ProtectStatus,LeaderId
		attrs := map[string]interface{}{
			"receivex":      flowStatistic.Receivex,
			"uploadx":       flowStatistic.Uploadx,
			"flow_count":    flowStatistic.FlowCount,
			"pub_flow":      flowStatistic.PubFlow,
			"notlocal_flow": flowStatistic.NotlocalFlow,
			"white_count":   flowStatistic.WhiteCount,
		}
		if err := modelBase.UpdateModelsByCondition(attrs, "vehicle_id = ? and interface_name = ?",
			[]interface{}{flowStatistic.VehicleId, flowStatistic.InterfaceName}...); err != nil {
			return fmt.Errorf("%s update flowStatisticParam err:%s", util.RunFuncName(), err.Error())
		}
	}

	//上报
	//会话状态
	logger.Logger.Print("%s flowStatistic info %+v", util.RunFuncName(), flowStatistic)
	logger.Logger.Info("%s flowStatistic info %+v", util.RunFuncName(), flowStatistic)

	pushActionTypeName := protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]
	pushVehicleid := vehicleId
	pushData := flowStatistic

	fPushData := push.CreatePushData(pushActionTypeName, pushVehicleid, pushData)

	push.GetPushervice().SetPushData(fPushData).Write()

	return nil
}
