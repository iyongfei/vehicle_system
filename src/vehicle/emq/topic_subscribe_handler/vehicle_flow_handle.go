package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/service/push"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleFlow(vehicleResult protobuf.GWResult, vehicleId string) error {
	//初始化资产默认分组
	assetGroup := &model.AreaGroup{
		AreaName:       response.UnGroupName,
		AreaCode:       util.RandomString(32),
		ParentAreaCode: "",
		TreeAreaCode:   "",
	}

	assetGroupModelBase := model_base.ModelBaseImpl(assetGroup)

	_, assetGroupRecordNotFound := assetGroupModelBase.GetModelByCondition("area_name = ?",
		[]interface{}{assetGroup.AreaName}...)

	if assetGroupRecordNotFound {
		err := assetGroupModelBase.InsertModel()
		if err != nil {
			return fmt.Errorf("%s insert asset ungroup err:%s", err)
		}
	}

	//parse
	flowParams := &protobuf.FlowParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), flowParams)
	if err != nil {
		logger.Logger.Print("%s unmarshal flowParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal flowParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal flowParam err:%s", util.RunFuncName(), err.Error())
	}

	var sendAssetFlows []map[string]interface{}

	for _, macItems := range flowParams.MacItems {
		mac := macItems.GetMac()

		macFlows := macItems.GetFlowItem()

		//创建或者获取asset
		asset := &model.Asset{
			VehicleId:  vehicleId,
			AssetId:    mac,
			AssetGroup: assetGroup.AreaCode,
		}
		assetModelBase := model_base.ModelBaseImpl(asset)
		_, recordNotFound := assetModelBase.GetModelByCondition("asset_id = ?", asset.AssetId)
		if recordNotFound {
			exist := checkoutAssetPrintInfos(asset.AssetId)
			asset.AccessNet = exist

			err := assetModelBase.InsertModel()
			if err != nil {
				continue
			}
		}

		AssetFlows := []interface{}{}
		//为asset插入flow
		for _, flowItem := range macFlows {
			flowItemId := flowItem.GetHash()
			flowInfo := &model.Flow{
				FlowId:    flowItemId,
				AssetId:   mac,
				VehicleId: vehicleId,
			}
			flowModelBase := model_base.ModelBaseImpl(flowInfo)
			_, flowRecordNotFound := flowModelBase.GetModelByCondition(
				"flow_id = ? and vehicle_id = ? and asset_id = ?",
				[]interface{}{flowInfo.FlowId, flowInfo.VehicleId, flowInfo.AssetId}...)
			flowModelBase.CreateModel(flowItem)

			if flowRecordNotFound {
				if err := flowModelBase.InsertModel(); err != nil {
					logger.Logger.Print("%s insert flowParam err:%s", util.RunFuncName(), err.Error())
					logger.Logger.Error("%s insert flowParam err:%s", util.RunFuncName(), err.Error())
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
					//add
					"src_dst_packets":      flowInfo.SrcDstPackets,
					"dst_src_packets":      flowInfo.DstSrcPackets,
					"host_name":            flowInfo.HostName,
					"has_passive":          flowInfo.HasPassive,
					"iat_flow_avg":         flowInfo.IatFlowAvg,
					"iat_flow_stddev":      flowInfo.IatFlowStddev,
					"data_ratio":           flowInfo.DataRatio,
					"str_data_ratio":       flowInfo.StrDataRatio,
					"pktlen_c_to_s_avg":    flowInfo.PktlenCToSAvg,
					"pktlen_c_to_s_stddev": flowInfo.PktlenCToSStddev,
					"pktlen_s_to_c_avg":    flowInfo.PktlenSToCAvg,
					"pktlen_s_to_c_stddev": flowInfo.PktlenSToCStddev,
					"tls_client_info":      flowInfo.TlsClientInfo,
					"ja3c":                 flowInfo.Ja3c,
				}
				if err := flowModelBase.UpdateModelsByCondition(attrs,
					"flow_id = ? and vehicle_id = ?",
					[]interface{}{flowInfo.FlowId, flowInfo.VehicleId}...); err != nil {
					logger.Logger.Print("%s update flowParam err:%s", util.RunFuncName(), err.Error())
					logger.Logger.Error("%s update flowParam err:%s", util.RunFuncName(), err.Error())
					continue
				}
			}
			AssetFlows = append(AssetFlows, flowInfo)
		}

		sendAssetFlow := map[string]interface{}{}
		sendAssetFlow["asset_id"] = mac
		sendAssetFlow["asset_flows"] = AssetFlows
		sendAssetFlows = append(sendAssetFlows, sendAssetFlow)
	}

	//删除临时表
	tFlow := &model.TempFlow{}
	tFlowModelBase := model_base.ModelBaseImpl(tFlow)
	if err := tFlowModelBase.DeleModelsByCondition("vehicle_id = ?",
		[]interface{}{vehicleId}...); err != nil {
		logger.Logger.Print("%s dele temp vehicle flow err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s dele temp vehicle flow err:%s", util.RunFuncName(), err.Error())

	}
	//插入临时表
	for _, macItems := range flowParams.MacItems {
		mac := macItems.GetMac()
		macFlows := macItems.GetFlowItem()

		for _, flowItem := range macFlows {
			flowItemId := flowItem.GetHash()
			tFlowInfo := &model.TempFlow{
				FlowId:    flowItemId,
				AssetId:   mac,
				VehicleId: vehicleId,
			}
			flowModelBase := model_base.ModelBaseImpl(tFlowInfo)
			_, flowRecordNotFound := flowModelBase.GetModelByCondition(
				"flow_id = ? and vehicle_id = ? and asset_id = ?",
				[]interface{}{tFlowInfo.FlowId, tFlowInfo.VehicleId, tFlowInfo.AssetId}...)
			flowModelBase.CreateModel(flowItem)
			if flowRecordNotFound {
				if err := flowModelBase.InsertModel(); err != nil {
					logger.Logger.Print("%s insert flowParam err:%s", util.RunFuncName(), err.Error())
					logger.Logger.Error("%s insert flowParam err:%s", util.RunFuncName(), err.Error())
					continue
				}
			}
		}
	}

	pushActionTypeName := protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]
	pushVehicleid := vehicleId
	pushData := sendAssetFlows
	fPushData := push.CreatePushData(pushActionTypeName, pushVehicleid, pushData)
	push.GetPushervice().SetPushData(fPushData).Write()

	return nil
}
