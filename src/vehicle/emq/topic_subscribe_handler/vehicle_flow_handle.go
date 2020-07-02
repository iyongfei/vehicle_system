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

			//输出log
			logger.Logger.Print("%s flows table flow_vehicleId%s,flow_assetId:%s,flowId:%s", util.RunFuncName(), flowInfo.VehicleId, flowInfo.AssetId, flowInfo.FlowId)
			logger.Logger.Info("%s flows table flow_vehicleId%s,flow_assetId:%s,flowId:%s", util.RunFuncName(), flowInfo.VehicleId, flowInfo.AssetId, flowInfo.FlowId)

			if flowRecordNotFound {
				if err := flowModelBase.InsertModel(); err != nil {
					logger.Logger.Print("%s insert flowParam err:%s", util.RunFuncName(), err.Error())
					logger.Logger.Error("%s insert flowParam err:%s", util.RunFuncName(), err.Error())
					continue
				}
			} else {
				//update
				//更新 排除VehicleId,Name,ProtectStatus,LeaderId

				attrs := CreateFlowAttr(flowInfo)

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

	//处理指纹标签flow  todo
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
		//_, totalByRate := model_helper.JudgeAssetCollectByteTotalRate(mac) //总流量
		//_, tlsRate := model_helper.JudgeAssetCollectTlsInfoRate(mac)       //tls
		//_, hostRate := model_helper.JudgeAssetCollectHostNameRate(mac)     //host
		//_, protoRate := model_helper.JudgeAssetCollectProtoFlowRate(mac)   //各协议流量
		//_, collectRate := model_helper.JudgeAssetCollectTimeRate(mac)      //采集时间
		//totalRate := totalByRate + tlsRate + hostRate + collectRate + protoRate
		//logger.Logger.Print("%s,metadata rate asset:%s,totalBytesRate:%f,tlsRate:%f,hostRate:%f,protoRate:%f,collectTimeRate:%f,totalRate:%f",
		//	util.RunFuncName(), mac, totalByRate, tlsRate, hostRate, protoRate, collectRate, totalRate)
		//logger.Logger.Info("%s,metadata rate asset:%s,totalBytesRate:%f,tlsRate:%f,hostRate:%f,protoRate:%f,collectTimeRate:%f,totalRate:%f",
		//	util.RunFuncName(), mac, totalByRate, tlsRate, hostRate, protoRate, collectRate, totalRate)

		fp := &model.Fprint{
			AssetId: mac,
		}
		fpModelBase := model_base.ModelBaseImpl(fp)
		err, _ := fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fp.AssetId}...)
		if err != nil {
			continue
		}
		//指纹完整度分数
		collectTotalRate := fp.CollectTotalRate

		//自动识别
		if util.RrgsTrim(fp.AutoCateId) == "" {
			updateFprintAutoCateId(vehicleId, mac)
		}

		logger.Logger.Print("%s fprint:%+v", util.RunFuncName(), fp)
		logger.Logger.Info("%s fprint:%+v", util.RunFuncName(), fp)

		//采集完毕
		if collectTotalRate >= conf.MinRate {
			if !fp.CollectFinish {
				updateAssetCollectTime(mac)
				updateFprintFinish(vehicleId, mac, true)

				_, _ = fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fp.AssetId}...)
				logger.Logger.Print("%s fprint %+v", util.RunFuncName(), fp)
				logger.Logger.Info("%s fprint %+v", util.RunFuncName(), fp)
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

			logger.Logger.Print("%s fprints table flow_vehicleId%s,flow_assetId:%s,flowId:%s", util.RunFuncName(), fprintFlow.VehicleId, fprintFlow.AssetId, fprintFlow.FlowId)
			logger.Logger.Info("%s fprints table flow_vehicleId%s,flow_assetId:%s,flowId:%s", util.RunFuncName(), fprintFlow.VehicleId, fprintFlow.AssetId, fprintFlow.FlowId)

			if flowRecordNotFound {
				if err := fpflowModelBase.InsertModel(); err != nil {
					logger.Logger.Print("%s insert fingerprint flowParam err:%s", util.RunFuncName(), err.Error())
					logger.Logger.Error("%s insert fingerprint flowParam err:%s", util.RunFuncName(), err.Error())
					continue
				}
			} else {
				//update
				//更新 排除VehicleId,Name,ProtectStatus,LeaderId

				attrs := CreateAttr(fprintFlow)

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
		updateFprint(vehicleId, mac)
		//如果没有记录，并且没有
		updateAssetCollectTime(mac)

		//log
		_, _ = fpModelBase.GetModelByCondition("asset_id = ?", []interface{}{fp.AssetId}...)

		logger.Logger.Print("%s fprint %+v", util.RunFuncName(), fp)
		logger.Logger.Info("%s fprint %+v", util.RunFuncName(), fp)
	}
}

/**
插入指纹资产
*/
func updateFprint(vehicleId string, mac string) {

	protoFlow, fcollectProto := model_helper.JudgeAssetCollectProtoFlowRate(mac)
	protoFlowBys, _ := json.Marshal(protoFlow)
	fprotoFlowStr := string(protoFlowBys)

	totalBytes, ftatalBytesRate := model_helper.JudgeAssetCollectByteTotalRate(mac) //总流量大小
	tlsInfoS, ftls := model_helper.JudgeAssetCollectTlsInfoRate(mac)
	hostNameS, fhost := model_helper.JudgeAssetCollectHostNameRate(mac)
	collectTime, fcollect_time := model_helper.JudgeAssetCollectTimeRate(mac)

	tlsInfoBys, _ := json.Marshal(tlsInfoS)
	tlsInfo := string(tlsInfoBys)

	hostNameBys, _ := json.Marshal(hostNameS)
	hostName := string(hostNameBys)

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
	fprint.CollectProtoRate = fcollectProto

	fprint.CollectBytes = totalBytes
	fprint.CollectBytesRate = ftatalBytesRate

	fprint.CollectTls = tlsInfo
	fprint.CollectTlsRate = ftls

	fprint.CollectHost = hostName
	fprint.CollectHostRate = fhost

	fprint.CollectTime = collectTime
	fprint.CollectTimeRate = fcollect_time
	fprint.CollectTotalRate = fcollectProto + ftatalBytesRate + ftls + fhost + fcollect_time

	fprint.AutoCateId = autoCateId

	logger.Logger.Info("%s updateFprint:%+v", util.RunFuncName(), fprint)
	logger.Logger.Print("%s updateFprint:%+v", util.RunFuncName(), fprint)

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
			"fprint_id":  fprint.FprintId,
			"vehicle_id": fprint.VehicleId,

			"collect_proto_flows": fprint.CollectProtoFlows,
			"collect_proto_rate":  fprint.CollectProtoRate,

			"collect_bytes":      fprint.CollectBytes,
			"collect_bytes_rate": fprint.CollectBytesRate,

			"collect_tls":      fprint.CollectTls,
			"collect_tls_rate": fprint.CollectTlsRate,

			"collect_host":      fprint.CollectHost,
			"collect_host_rate": fprint.CollectHostRate,

			"collect_time":      fprint.CollectTime,
			"collect_time_rate": fprint.CollectTimeRate,

			"collect_total_rate": fprint.CollectTotalRate,

			"auto_cate_id": fprint.AutoCateId,
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
		AssetId: mac,
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
		if fprint.CollectStart == 0 {
			attrs = map[string]interface{}{
				"collect_start": uint64(time.Now().Unix()),
			}
		} else {
			ctime := fprint.CollectTime
			stime := fprint.CollectStart
			now := uint64(time.Now().Unix())
			distanceTime := ctime + uint32(now-stime)

			attrs = map[string]interface{}{
				"collect_start": uint64(time.Now().Unix()),
				"collect_time":  distanceTime,
			}
		}
		if err := fpModelBase.UpdateModelsByCondition(attrs,
			"asset_id = ?", []interface{}{fprint.AssetId}...); err != nil {
			logger.Logger.Print("%s asset:%s,err:%s", util.RunFuncName(), fprint.AssetId, err.Error())
			logger.Logger.Error("%s asset:%s,err:%s", util.RunFuncName(), fprint.AssetId, err.Error())
		}
	}
}

/*
更新资产类别识别
*/
func updateFprintAutoCateId(vehicleId string, mac string) {
	//识别类别
	autoCateId := model_helper.JudgeAssetCate(mac)

	fprint := &model.Fprint{
		AssetId:    mac,
		FprintId:   util.RandomString(32),
		VehicleId:  vehicleId,
		AutoCateId: autoCateId,
	}

	fpModelBase := model_base.ModelBaseImpl(fprint)

	attrs := map[string]interface{}{
		"auto_cate_id": fprint.AutoCateId,
	}
	if err := fpModelBase.UpdateModelsByCondition(attrs, "asset_id = ?", []interface{}{fprint.AssetId}...); err != nil {
		//todo
		logger.Logger.Print("%s update flowParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s update flowParam err:%s", util.RunFuncName(), err.Error())
	}
}

/**
更新采集完毕标志
*/
func updateFprintFinish(vehicleId string, mac string, collectFinish bool) {
	//插入资产指纹信息
	fprint := &model.Fprint{
		AssetId:       mac,
		FprintId:      util.RandomString(32),
		VehicleId:     vehicleId,
		CollectFinish: true,
	}

	fpModelBase := model_base.ModelBaseImpl(fprint)

	attrs := map[string]interface{}{
		"collect_finish": fprint.AutoCateId,
	}
	if err := fpModelBase.UpdateModelsByCondition(attrs, "asset_id = ?", []interface{}{fprint.AssetId}...); err != nil {
		//todo
		logger.Logger.Print("%s update flowParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s update flowParam err:%s", util.RunFuncName(), err.Error())
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

///////////////////////////////

func CreateFlowAttr(fprintFlow *model.Flow) map[string]interface{} {

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
	return attrs
}

func CreateAttr(fprintFlow *model.FprintFlow) map[string]interface{} {

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
	return attrs
}
