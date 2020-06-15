package model_helper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

const (
	CATE_ID      = "cate_id"
	PROTOS       = "protos"
	PROTOS_FLOWS = "proto_flows"
)

func GetFingerPrintProtolFlowsRate() []map[string]map[string]float64 {
	//查询指纹库所有的类别
	type CateId struct {
		CateId string
	}

	cateIdList := []*CateId{}

	execSql := "SELECT cate_id FROM finger_prints GROUP BY cate_id;"
	err := mysql.QueryRawsqlScanStruct(execSql, []interface{}{}, &cateIdList)
	if err != nil {
		//todo
	}
	cidMaps := []map[string]map[string]uint64{}
	for _, cid := range cateIdList {
		cidMap := GetCateFingerPrintProtolFlow(cid.CateId)
		cidMaps = append(cidMaps, cidMap)
	}
	fCidMaps := []map[string]map[string]float64{}

	for _, cidMap := range cidMaps {
		fCidMap := map[string]map[string]float64{
			CATE_ID:      map[string]float64{},
			PROTOS:       map[string]float64{},
			PROTOS_FLOWS: map[string]float64{},
		}
		//1
		cid := cidMap[CATE_ID]
		for cidTmp, _ := range cid {
			fCidMap[CATE_ID][cidTmp] = 1
		}

		//2
		protoTmp := cidMap[PROTOS]
		protoTotalCount := protoTmp[TOTAL]
		if cidMapProtoV, ok := cidMap[PROTOS]; ok {
			for p, pcount := range cidMapProtoV {
				rate := fmt.Sprintf("%.3f", float64(pcount)/float64(protoTotalCount))
				frate, _ := strconv.ParseFloat(rate, 64)
				fCidMap[PROTOS][p] = frate
			}
		}
		//3
		protoFlowsTmp := cidMap[PROTOS_FLOWS]
		protoFlowsTotalCount := protoFlowsTmp[TOTAL]

		if cidMapProtoFlowsV, ok := cidMap[PROTOS_FLOWS]; ok {
			for p, pcount := range cidMapProtoFlowsV {
				rate := fmt.Sprintf("%.3f", float64(pcount)/float64(protoFlowsTotalCount))
				frate, _ := strconv.ParseFloat(rate, 64)
				fCidMap[PROTOS_FLOWS][p] = frate
			}
		}
		//添加
		fCidMaps = append(fCidMaps, fCidMap)
	}

	for _, fcidMap := range fCidMaps {
		for _, v := range fcidMap {
			delete(v, TOTAL)
		}
	}
	return fCidMaps
}

///**
//获取某类别资产所有信息
//*/
func GetCateFingerPrintProtolFlow(cateId string) map[string]map[string]uint64 {
	//获取该表下某分类的所有资产

	fingerPrint := &model.FingerPrint{
		CateId: cateId,
	}
	fingerPrintModelBase := model_base.ModelBaseImpl(fingerPrint)

	fingerPrintList := []*model.FingerPrint{}
	err := fingerPrintModelBase.GetModelListByCondition(&fingerPrintList, "cate_id = ?", []interface{}{fingerPrint.CateId}...)
	if err != nil {
		//todo
	}

	fprotos := map[string]uint64{}

	fprotosFlowBytes := map[string]uint64{}

	for _, fp := range fingerPrintList {

		protos := fp.Protos
		protosRate := fp.ProtoRate
		logger.Logger.Print("%s fingerPrint protos:%+v,protosRate%+v", util.RunFuncName(), protos, protosRate)
		logger.Logger.Info("%s fingerPrint protos:%+v,protosRate%+v", util.RunFuncName(), protos, protosRate)
		//过滤空字段
		if len(protos) == 0 && len(protosRate) == 0 {
			continue
		}
		//数量
		var protosMap map[string]uint64
		protoUnmarshalErr := json.Unmarshal([]byte(string(protos)), &protosMap)
		if protoUnmarshalErr != nil {
			logger.Logger.Print("%s fingerPrint unmarshal protos err:%+v", util.RunFuncName(), protoUnmarshalErr)
			logger.Logger.Error("%s fingerPrint unmarshal protos err:%+v", util.RunFuncName(), protoUnmarshalErr)
			continue
		}
		logger.Logger.Print("%s fingerPrint unmarshal protos:%+v", util.RunFuncName(), protosMap)
		logger.Logger.Info("%s fingerPrint unmarshal protos:%+v", util.RunFuncName(), protosMap)

		//占比
		protosFlowMap := map[string]uint64{}
		protosRateUnmarshalErr := json.Unmarshal([]byte(protosRate), &protosFlowMap)
		if protosRateUnmarshalErr != nil {
			logger.Logger.Print("%s fingerPrint unmarshal proto_flowmap err:%+v", util.RunFuncName(), protosRateUnmarshalErr)
			logger.Logger.Error("%s fingerPrint unmarshal proto_flowmap err:%+v", util.RunFuncName(), protosRateUnmarshalErr)

			continue
		}
		logger.Logger.Print("%s fingerPrint unmarshal proto_flowmap:%+v", util.RunFuncName(), protosMap)
		logger.Logger.Info("%s fingerPrint unmarshal proto_flowmap:%+v", util.RunFuncName(), protosMap)

		//去除键t

		for proto, protoCount := range protosMap {
			if v, ok := fprotos[proto]; ok {
				fprotos[proto] = v + protoCount
			} else {
				fprotos[proto] = protoCount
			}
		}

		for proto, flow := range protosFlowMap {
			if v, ok := fprotosFlowBytes[proto]; ok {
				fprotosFlowBytes[proto] = v + flow
			} else {
				fprotosFlowBytes[proto] = flow
			}
		}
	}
	//fprotos := map[string]uint64{}

	fcateId := map[string]uint64{}
	fcateId[cateId] = 1

	retMap := map[string]map[string]uint64{}

	retMap[CATE_ID] = fcateId
	retMap[PROTOS] = fprotos
	retMap[PROTOS_FLOWS] = fprotosFlowBytes

	/**
	{"DIAMETER":1,"FREE_205":1,"NETFLOW":1,"RADIUS":1,"RX":1,"SOULSEEK":1,"SSDP":1,"STEAM":1,"SYSLOG":1,"t":15}
	{"FREE_205":17173,"NETFLOW":15120,"RADIUS":11934,"SMBV1":10380,"SSDP":10224,"STEALTHNET":12220,"STEAM":13530,"TELEGRAM":11538,"WHOIS_DAS":16362,"t":146951}
	*/
	return retMap
}
