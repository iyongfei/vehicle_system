package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/service/flow"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleFlow(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	flowParams := &protobuf.FlowParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), flowParams)
	if err != nil {
		logger.Logger.Print("%s unmarshal flowParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal flowParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal flowParam err:%s", util.RunFuncName(), err.Error())
	}

	var sendFlows []*model.Flow

	for _, flowItem := range flowParams.FlowItem {

		flowItemId := flowItem.GetHash()
		flowInfo := &model.Flow{
			FlowId:    flowItemId,
			VehicleId: vehicleId,
		}
		modelBase := model_base.ModelBaseImpl(flowInfo)
		_, recordNotFound := modelBase.GetModelByCondition(
			"flow_id = ? and vehicle_id = ?", []interface{}{flowInfo.FlowId, flowInfo.VehicleId}...)
		modelBase.CreateModel(flowItem)

		if recordNotFound {
			if err := modelBase.InsertModel(); err != nil {
				logger.Logger.Print("%s insert flowParam err:%s", util.RunFuncName(), err.Error())
				logger.Logger.Error("%s insert flowParam err:%s", util.RunFuncName(), err.Error())
				//return fmt.Errorf("%s insert flow err:%s",util.RunFuncName(),err.Error())
				continue
			}
		} else {
			//update
			//更新 排除VehicleId,Name,ProtectStatus,LeaderId
			attrs := map[string]interface{}{
				"hash":           flowInfo.Hash,
				"src_ip":         flowInfo.SrcIp,
				"src_port":       flowInfo.SrcPort,
				"dst_ip":         flowInfo.DstIp,
				"dst_port":       flowInfo.DstPort,
				"protocol":       flowInfo.Protocol,
				"flow_info":      flowInfo.FlowInfo,
				"safe_type":      flowInfo.SafeType,
				"safe_info":      flowInfo.SafeInfo,
				"start_time":     flowInfo.StartTime,
				"last_seen_time": flowInfo.LastSeenTime,
				"src_dst_bytes":  flowInfo.SrcDstBytes,
				"dst_src_bytes":  flowInfo.DstSrcBytes,
				"stat":           flowInfo.Stat,
			}
			if err := modelBase.UpdateModelsByCondition(attrs,
				"flow_id = ? and vehicle_id = ?", []interface{}{flowInfo.FlowId, flowInfo.VehicleId}...); err != nil {
				logger.Logger.Print("%s update flowParam err:%s", util.RunFuncName(), err.Error())
				logger.Logger.Error("%s update flowParam err:%s", util.RunFuncName(), err.Error())
				continue
			}
		}
		sendFlows = append(sendFlows, flowInfo)
	}

	//删除临时表
	tFlow := &model.TempFlow{}
	tFlowModelBase := model_base.ModelBaseImpl(tFlow)
	if err := tFlowModelBase.DeleModelsByCondition("vehicle_id = ?",
		[]interface{}{vehicleId}...); err != nil {
		logger.Logger.Print("%s dele temp vehicle flow err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s dele temp vehicle flow err:%s", util.RunFuncName(), err.Error())

	}
	for _, tFlowItem := range flowParams.FlowItem {

		tFlowItemId := tFlowItem.GetHash()
		tFlowInfo := &model.TempFlow{
			FlowId:    tFlowItemId,
			VehicleId: vehicleId,
		}
		tFlowModelBase := model_base.ModelBaseImpl(tFlowInfo)
		_, recordNotFound := tFlowModelBase.GetModelByCondition(
			"flow_id = ? and vehicle_id = ?", []interface{}{tFlowInfo.FlowId, tFlowInfo.VehicleId}...)
		tFlowModelBase.CreateModel(tFlowItem)

		if recordNotFound {
			if err := tFlowModelBase.InsertModel(); err != nil {
				logger.Logger.Print("%s insert temp flow err:%s", util.RunFuncName(), err.Error())
				logger.Logger.Error("%s insert temp flow err:%s", util.RunFuncName(), err.Error())
				//return fmt.Errorf("%s insert flow err:%s",util.RunFuncName(),err.Error())
				continue
			}
		}
	}

	//会话状态
	logger.Logger.Print("%s flow info %+v", util.RunFuncName(), sendFlows)
	logger.Logger.Info("%s flow info %+v", util.RunFuncName(), sendFlows)

	pushActionTypeName := protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]
	pushVehicleid := vehicleId
	pushData := sendFlows

	fPushData := flow.CreatePushData(pushActionTypeName, pushVehicleid, pushData)

	flow.GetFlowService().SetFlowData(fPushData).WriteFlow()

	return nil
}
