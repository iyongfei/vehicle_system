package topic_subscribe_handler

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/model/model_helper"
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
					"flow_id = ? and vehicle_id = ? and asset_id = ?",
					[]interface{}{flowInfo.FlowId, flowInfo.VehicleId, flowInfo.AssetId}...); err != nil {
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

	//处理临时表
	deleTmpFlows(vehicleId, flowParams)

	//处理指纹标签flow
	handleFprintFlows(vehicleId, flowParams)

	pushActionTypeName := protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]
	pushVehicleid := vehicleId
	pushData := sendAssetFlows
	fPushData := push.CreatePushData(pushActionTypeName, pushVehicleid, pushData)
	push.GetPushervice().SetPushData(fPushData).Write()

	return nil
}

/**
处理指纹库flows
*/
func handleFprintFlows(vehicleId string, flowParams *protobuf.FlowParam) {
	for _, macItems := range flowParams.MacItems {
		mac := macItems.GetMac()
		//判断资产是否达标指纹完整度
		totalByRate := model_helper.JudgeAssetCollectByteTotalRate(mac) //总流量
		tlsRate := model_helper.JudgeAssetCollectTlsInfoRate(mac)       //tls
		hostRate := model_helper.JudgeAssetCollectHostNameRate(mac)     //host
		protoRate := model_helper.JudgeAssetCollectProtoFlowRate(mac)   //各协议流量
		collectRate := model_helper.JudgeAssetCollectTimeRate(mac)      //采集时间

		totalRate := totalByRate + tlsRate + hostRate + collectRate + protoRate
		logger.Logger.Print("%s,asset:%s,totalByRate:%f,tlsRate:%f,hostRate:%f,protoRate:%f,collectRate:%f,totalRate:%f",
			util.RunFuncName(), mac, totalByRate, tlsRate, hostRate, protoRate, collectRate, totalRate)
		logger.Logger.Info("%s,asset:%s,totalByRate:%f,tlsRate:%f,hostRate:%f,protoRate:%f,collectRate:%f,totalRate:%f",
			util.RunFuncName(), mac, totalByRate, tlsRate, hostRate, protoRate, collectRate, totalRate)

		//无记录，更新
		fp := &model.Fprint{
			AssetId: mac,
		}
		fpModelBase := model_base.ModelBaseImpl(fp)
		err, _ := fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fp.AssetId}...)
		if err != nil {
			continue
		}

		//识别

		if totalRate >= conf.MinRate {
			if !fp.CollectFinish {
				updateAssetCollectTime(mac)
				updateFprint(vehicleId, mac, true)
			}

			if util.RrgsTrim(fp.AutoCateId) == "" {
				updateFprintAutoCateId(vehicleId, mac)
			}

			continue
		}

		macFlows := macItems.GetFlowItem()
		for _, flowItem := range macFlows {
			flowItemId := flowItem.GetHash()
			fprintFlow := &model.FprintFlow{
				FlowId:    flowItemId,
				AssetId:   mac,
				VehicleId: vehicleId,
			}
			fpflowModelBase := model_base.ModelBaseImpl(fprintFlow)

			_, flowRecordNotFound := fpflowModelBase.GetModelByCondition(
				"flow_id = ? and vehicle_id = ? and asset_id = ?",
				[]interface{}{fprintFlow.FlowId, fprintFlow.VehicleId, fprintFlow.AssetId}...)

			fpflowModelBase.CreateModel(flowItem)

			if flowRecordNotFound {
				if err := fpflowModelBase.InsertModel(); err != nil {
					logger.Logger.Print("%s insert fingerprint flowParam err:%s", util.RunFuncName(), err.Error())
					logger.Logger.Error("%s insert fingerprint flowParam err:%s", util.RunFuncName(), err.Error())
					continue
				}
			} else {
				//update
				//更新 排除VehicleId,Name,ProtectStatus,LeaderId
				attrs := map[string]interface{}{
					"hash":           fprintFlow.Hash,
					"src_ip":         fprintFlow.SrcIp,
					"src_port":       fprintFlow.SrcPort,
					"dst_ip":         fprintFlow.DstIp,
					"dst_port":       fprintFlow.DstPort,
					"protocol":       fprintFlow.Protocol,
					"flow_info":      fprintFlow.FlowInfo,
					"safe_type":      fprintFlow.SafeType,
					"safe_info":      fprintFlow.SafeInfo,
					"start_time":     fprintFlow.StartTime,
					"last_seen_time": fprintFlow.LastSeenTime,
					"src_dst_bytes":  fprintFlow.SrcDstBytes,
					"dst_src_bytes":  fprintFlow.DstSrcBytes,
					"stat":           fprintFlow.Stat,
					//add
					"src_dst_packets":      fprintFlow.SrcDstPackets,
					"dst_src_packets":      fprintFlow.DstSrcPackets,
					"host_name":            fprintFlow.HostName,
					"has_passive":          fprintFlow.HasPassive,
					"iat_flow_avg":         fprintFlow.IatFlowAvg,
					"iat_flow_stddev":      fprintFlow.IatFlowStddev,
					"data_ratio":           fprintFlow.DataRatio,
					"str_data_ratio":       fprintFlow.StrDataRatio,
					"pktlen_c_to_s_avg":    fprintFlow.PktlenCToSAvg,
					"pktlen_c_to_s_stddev": fprintFlow.PktlenCToSStddev,
					"pktlen_s_to_c_avg":    fprintFlow.PktlenSToCAvg,
					"pktlen_s_to_c_stddev": fprintFlow.PktlenSToCStddev,
					"tls_client_info":      fprintFlow.TlsClientInfo,
					"ja3c":                 fprintFlow.Ja3c,
				}
				if err := fpflowModelBase.UpdateModelsByCondition(attrs,
					"flow_id = ? and vehicle_id = ? and asset_id = ?",
					[]interface{}{fprintFlow.FlowId, fprintFlow.VehicleId, fprintFlow.AssetId}...); err != nil {
					logger.Logger.Print("%s update flowParam err:%s", util.RunFuncName(), err.Error())
					logger.Logger.Error("%s update flowParam err:%s", util.RunFuncName(), err.Error())
					continue
				}
			}
		}

		//更新指纹库
		updateFprint(vehicleId, mac, false)
		//如果没有记录，并且没有
		updateAssetCollectTime(mac)
	}
}

func updateFprintAutoCateId(vehicleId string, mac string) {
	//识别类别
	autoCateId := model_helper.JudgeAssetCate(mac)

	fprint := &model.Fprint{
		AssetId:   mac,
		FprintId:  util.RandomString(32),
		VehicleId: vehicleId,
	}

	fpModelBase := model_base.ModelBaseImpl(fprint)

	err, recordNotFound := fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fprint.AssetId}...)

	fprint.AutoCateId = autoCateId
	if err != nil {
		//todo
	}

	if recordNotFound {
		err := fprint.InsertModel()
		if err != nil {
			//todo
		}
	} else {
		attrs := map[string]interface{}{
			"auto_cate_id": fprint.AutoCateId,
		}
		if err := fpModelBase.UpdateModelsByCondition(attrs, "asset_id = ?", []interface{}{fprint.AssetId}...); err != nil {
			//todo
			logger.Logger.Print("%s update flowParam err:%s", util.RunFuncName(), err.Error())
			logger.Logger.Error("%s update flowParam err:%s", util.RunFuncName(), err.Error())
		}
	}
}

/**
插入指纹资产
*/
func updateFprint(vehicleId string, mac string, finishFlg bool) {
	//插入资产指纹信息
	protoFlow := model_helper.GetRankAssetCollectProtoFlow(mac)
	protoFlowBys, _ := json.Marshal(protoFlow)
	fprotoFlowStr := string(protoFlowBys)

	totalBytes := model_helper.GetAssetCollectByteTotal(mac) //总流量大小
	tlsInfo := model_helper.GetAssetCollectTlsInfo(mac)
	hostName := model_helper.GetAssetCollectHostName(mac)
	collectTime := model_helper.GetAssetCollectTime(mac)

	//识别类别
	autoCateId := model_helper.JudgeAssetCate(mac)

	fprint := &model.Fprint{
		AssetId:   mac,
		FprintId:  util.RandomString(32),
		VehicleId: vehicleId,
	}

	fpModelBase := model_base.ModelBaseImpl(fprint)

	err, recordNotFound := fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fprint.AssetId}...)

	fprint.CollectProtoFlows = fprotoFlowStr
	fprint.CollectHost = hostName
	fprint.CollectTls = tlsInfo
	fprint.CollectBytes = totalBytes

	if finishFlg {
		fprint.CollectTime = collectTime
	}

	fprint.CollectFinish = finishFlg
	fprint.AutoCateId = autoCateId

	if err != nil {
		//todo
	}
	if recordNotFound {
		err := fprint.InsertModel()
		if err != nil {
			//todo
		}
	} else {
		attrs := map[string]interface{}{
			"fprint_id":           fprint.FprintId,
			"vehicle_id":          fprint.VehicleId,
			"collect_proto_flows": fprint.CollectProtoFlows,
			"collect_host":        fprint.CollectHost,
			"collect_tls":         fprint.CollectTls,
			"collect_bytes":       fprint.CollectBytes,
			"collect_time":        fprint.CollectTime,
			"collect_finish":      fprint.CollectFinish,
			"auto_cate_id":        fprint.AutoCateId,
		}
		if err := fpModelBase.UpdateModelsByCondition(attrs, "asset_id = ?", []interface{}{fprint.AssetId}...); err != nil {
			//todo
			logger.Logger.Print("%s update flowParam err:%s", util.RunFuncName(), err.Error())
			logger.Logger.Error("%s update flowParam err:%s", util.RunFuncName(), err.Error())
		}
	}
}

/*
更新采集时间
*/
func updateAssetCollectTime(mac string) {
	fprint := &model.Fprint{
		AssetId:      mac,
		CollectStart: uint64(time.Now().Unix()),
		CollectEnd:   uint64(time.Now().Unix()),
	}

	fpModelBase := model_base.ModelBaseImpl(fprint)

	err, recordNotFound := fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fprint.AssetId}...)
	if err != nil {
		//todo
		logger.Logger.Print("%s asset:%s,err:%s", util.RunFuncName(), fprint.AssetId, err.Error())
		logger.Logger.Error("%s asset:%s,err:%s", util.RunFuncName(), fprint.AssetId, err.Error())
	}

	if recordNotFound {
		insertErr := fpModelBase.InsertModel()

		if insertErr != nil {
			//todo
			logger.Logger.Print("%s asset:%s,err:%s", util.RunFuncName(), fprint.AssetId, insertErr.Error())
			logger.Logger.Error("%s asset:%s,err:%s", util.RunFuncName(), fprint.AssetId, insertErr.Error())
		}
	} else {
		var attrs map[string]interface{}
		if fprint.CollectStart == 0 && fprint.CollectEnd == 0 {
			attrs = map[string]interface{}{
				"collect_start": uint64(time.Now().Unix()),
				"collect_end":   uint64(time.Now().Unix()),
			}
		} else {
			attrs = map[string]interface{}{
				"collect_end": uint64(time.Now().Unix()),
			}
		}
		if err := fpModelBase.UpdateModelsByCondition(attrs,
			"asset_id = ?", []interface{}{fprint.AssetId}...); err != nil {
			logger.Logger.Print("%s asset:%s,err:%s", util.RunFuncName(), fprint.AssetId, err.Error())
			logger.Logger.Error("%s asset:%s,err:%s", util.RunFuncName(), fprint.AssetId, err.Error())
		}
	}
}

/**
处理临时表
*/
func deleTmpFlows(vehicleId string, flowParams *protobuf.FlowParam) {

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

}
