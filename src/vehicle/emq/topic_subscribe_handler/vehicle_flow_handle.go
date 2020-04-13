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

func HandleVehicleFlow(vehicleResult protobuf.GWResult) error {
	vehicleId := vehicleResult.GetGUID()
	//parse
	flowParams := &protobuf.FlowParam{}
	err:=proto.Unmarshal(vehicleResult.GetParam(),flowParams)
	if err != nil {
		logger.Logger.Print("%s unmarshal flowParam err:%s",util.RunFuncName(),err.Error())
		logger.Logger.Error("%s unmarshal flowParam err:%s",util.RunFuncName(),err.Error())
		return fmt.Errorf("%s unmarshal flowParam err:%s",util.RunFuncName(),err.Error())
	}
	for _,flowItem:=  range flowParams.FlowItem{

		flowItemId := flowItem.GetHash()
		flowInfo:=&model.Flow{
			FlowId:flowItemId,
			VehicleId:vehicleId,
		}
		modelBase := model_base.ModelBaseImpl(flowInfo)
		_,recordNotFound :=modelBase.GetModelByCondition(
			"flow_id = ? and vehicle_id = ?",[]interface{}{flowInfo.FlowId,flowInfo.VehicleId}...)
		modelBase.CreateModel(flowItem)

		logger.Logger.Print("%s flowInfo：%+v",util.RunFuncName(),flowInfo)
		if recordNotFound{
			if err:=modelBase.InsertModel();err!=nil{
				logger.Logger.Print("%s insert flowParam err:%s",util.RunFuncName(),err.Error())
				logger.Logger.Error("%s insert flowParam err:%s",util.RunFuncName(),err.Error())
				logger.Logger.Print("%s insert flowInfo:%+v, err：%+v",util.RunFuncName(),flowInfo,err.Error())
				//return fmt.Errorf("%s insert flow err:%s",util.RunFuncName(),err.Error())
				continue
			}
		}else {
			//update
			//更新 排除VehicleId,Name,ProtectStatus,LeaderId
			attrs:= map[string]interface{}{
				"hash":flowInfo.Hash,
				"src_ip":flowInfo.SrcIp,
				"src_port":flowInfo.SrcPort,
				"dst_ip":flowInfo.DstIp,
				"dst_port":flowInfo.DstPort,
				"protocol":flowInfo.Protocol,
				"flow_info":flowInfo.FlowInfo,
				"safe_type":flowInfo.SafeType,
				"safe_info":flowInfo.SafeInfo,
				"start_time":flowInfo.StartTime,
				"last_seen_time":flowInfo.LastSeenTime,
				"src_dst_bytes":flowInfo.SrcDstBytes,
				"dst_src_bytes":flowInfo.DstSrcBytes,
				"stat":flowInfo.Stat,
			}
			if err:=modelBase.UpdateModelsByCondition(attrs,
				"flow_id = ? and vehicle_id = ?",[]interface{}{flowInfo.FlowId,flowInfo.VehicleId}...);err!=nil{
				logger.Logger.Print("%s update flowParam err:%s",util.RunFuncName(),err.Error())
				logger.Logger.Error("%s update flowParam err:%s",util.RunFuncName(),err.Error())

				//return fmt.Errorf("%s update flow err:%s",util.RunFuncName(),err.Error())
				continue
			}
		}

		//会话状态
		if flowInfo.Stat == uint8(protobuf.FlowStat_FST_FINISH){
			logger.Logger.Print("%s flow info %+v",util.RunFuncName(),flowInfo)
			logger.Logger.Info("%s flow info %+v",util.RunFuncName(),flowInfo)
			flowMap := map[string]interface{}{}
			flowMap[flowInfo.VehicleId] = []*model.Flow{flowInfo}
			flow.GetFlowService().SetFlowData(flowMap).WriteFlow()
		}
	}

	return nil
}