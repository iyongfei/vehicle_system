package model

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

const (
	CATE_ID      = "cate_id"
	PROTOS       = "protos"
	PROTOS_FLOWS = "protoFlows"
)

func GetFingerPrintProtolRate() {
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

	cidMaps := []map[string]interface{}{}
	for _, cid := range cateIdList {
		cidMap := GetCateFingerPrintProtolRate(cid.CateId)
		fmt.Println("vvcidMapvvv", cidMap)
		cidMaps = append(cidMaps, cidMap)
	}

	for _, v := range cidMaps {
		fmt.Println("vvvvv", v)
	}

}

/**
获取某类别资产信息
*/
func GetCateFingerPrintProtolRate(cateId string) map[string]interface{} {
	//获取该表下某分类的所有资产

	fingerPrint := &FingerPrint{
		CateId: cateId,
	}
	fingerPrintModelBase := model_base.ModelBaseImpl(fingerPrint)

	fingerPrintList := []*FingerPrint{}
	err := fingerPrintModelBase.GetModelListByCondition(&fingerPrintList, "cate_id = ?", []interface{}{fingerPrint.CateId}...)
	if err != nil {
		//todo
	}

	fprotos := map[string]uint64{}

	fprotosFlowBytes := map[string]uint64{}

	for _, fp := range fingerPrintList {
		logger.Logger.Print("%s FingerPrint:%+v", util.RunFuncName(), fp)
		logger.Logger.Info("%s FingerPrint:%+v", util.RunFuncName(), fp)

		protos := fp.Protos
		protosRate := fp.ProtoRate
		logger.Logger.Print("%s FingerPrint protos:%+v", util.RunFuncName(), protos)
		logger.Logger.Info("%s FingerPrint protosRate:%+v", util.RunFuncName(), protosRate)
		//过滤空字段
		if len(protos) == 0 && len(protosRate) == 0 {
			continue
		}
		//数量
		//{"DIAMETER":1,"FREE_205":1,"NETFLOW":1,"RADIUS":1,"RX":1,"SOULSEEK":1,"SSDP":1,"STEAM":1,"SYSLOG":1,"t":15}
		protosMap := map[string]uint64{}
		logger.Logger.Print("%s FingerPrint protosprotos:%+v", util.RunFuncName(), protos)
		protoUnmarshalErr := json.Unmarshal([]byte(protos), &protosMap)
		if protoUnmarshalErr != nil {
			fmt.Println(protoUnmarshalErr, "jsldfj")
			continue
		}
		logger.Logger.Print("%s FingerPrint protosMap:%+v", util.RunFuncName(), protosMap)
		logger.Logger.Info("%s FingerPrint protosMap:%+v", util.RunFuncName(), protosMap)

		//占比
		//{"FREE_205":17173,"NETFLOW":15120,"RADIUS":11934,"SMBV1":10380,"SSDP":10224,"STEALTHNET":12220,"STEAM":13530,"TELEGRAM":11538,"WHOIS_DAS":16362,"t":146951}
		protosFlowMap := map[string]uint64{}
		//protosRate = "`" + protosRate + "`"
		protosRateUnmarshalErr := json.Unmarshal([]byte(protosRate), &protosFlowMap)
		if protosRateUnmarshalErr != nil {
			continue
		}
		logger.Logger.Print("%s FingerPrint: protosFlowMap%+v", util.RunFuncName(), protosFlowMap)
		logger.Logger.Info("%s FingerPrint protosFlowMap:%+v", util.RunFuncName(), protosFlowMap)

		protosFlddowMap := map[string]uint64{}
		vbb := `{"FASTTRACK":1,"IMO":1,"IRC":1,"MYSQL":1,"NETFLIX":1,"RDP":1,"STEAM":1,"TVUPLAYER":1,"TWITCH":1,"t":15}`
		_ = json.Unmarshal([]byte(vbb), &protosFlddowMap)
		fmt.Println("sprotosFlddowMap", protosFlddowMap)

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

	retMap := map[string]interface{}{}

	retMap[CATE_ID] = cateId
	retMap[PROTOS] = fprotos
	retMap[PROTOS_FLOWS] = fprotosFlowBytes

	/**
	{"DIAMETER":1,"FREE_205":1,"NETFLOW":1,"RADIUS":1,"RX":1,"SOULSEEK":1,"SSDP":1,"STEAM":1,"SYSLOG":1,"t":15}
	{"FREE_205":17173,"NETFLOW":15120,"RADIUS":11934,"SMBV1":10380,"SSDP":10224,"STEALTHNET":12220,"STEAM":13530,"TELEGRAM":11538,"WHOIS_DAS":16362,"t":146951}
	*/
	return retMap
}
