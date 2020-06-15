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

/**
获取指纹库所有指纹的protos平均占比值
*/

func GetFpProtosAverage() map[string]float64 {
	cateRates := map[string]float64{}

	//查询指纹库所有的类别
	type CateId struct {
		CateId string
	}

	cateIdList := []*CateId{}

	execSql := "SELECT cate_id FROM finger_prints GROUP BY cate_id;"
	err := mysql.QueryRawsqlScanStruct(execSql, []interface{}{}, &cateIdList)
	if err != nil {
		return cateRates
	}
	/**
	{
	{"cate_id":["dns","dbs"]},
	{"cate_id":["dns","dbs"]},
	*/
	fprintMaps := map[string][]string{}
	for _, cid := range cateIdList {
		fprintMap := GetCateFingerPrintProtoList(cid.CateId) //map[string][]string
		for cidTmp, cidProtos := range fprintMap {
			fprintMaps[cidTmp] = cidProtos
		}
	}

	//{"cate_id":["dns","dbs"]}
	for tcid, tcidProtos := range fprintMaps {
		fmt.Println("fprintMaps:::::", tcid, tcidProtos)
		var cid = tcid
		var cidProtos = tcidProtos
		//获取每一个类别下的条目["dns","rtp"]
		fmt.Println("cidProtos===", cidProtos)
		var fpProtoRateSlice []float64

		cidFlows := GetCateOneFingerPrintProtolFlows(cid)

		for _, cidFlow := range cidFlows {
			//[SMBV1 STEALTHNET STEAM t DIAMETER NETFLOW RADIUS DIRECTCONNECT SSDP TELEGRAM]
			fmt.Println("cidFlow::::", cidFlow)

			tempCidFlow := []string{}

			for _, macProto := range cidFlow {

				if util.IsExistInSlice(macProto, cidProtos) {
					fmt.Println("macProto===", macProto)
					tempCidFlow = append(tempCidFlow, macProto)
				}
			}

			rate := fmt.Sprintf("%.3f", float64(len(tempCidFlow))/float64(len(cidProtos)))
			frate, _ := strconv.ParseFloat(rate, 64)

			fpProtoRateSlice = append(fpProtoRateSlice, frate)
		}

		var fpProtoRateTotal float64
		for _, fpProtoRate := range fpProtoRateSlice {
			fpProtoRateTotal += fpProtoRate
		}

		ffrate := fmt.Sprintf("%.3f", fpProtoRateTotal/float64(len(fpProtoRateSlice)))
		frate, _ := strconv.ParseFloat(ffrate, 64)
		cateRates[cid] = frate
	}

	return cateRates
}

/**
获取某类别资产信息
*/

func GetCateOneFingerPrintProtolFlows(cateId string) [][]string {
	fprints := []*model.FingerPrint{}

	fingerPrintModelBase := model_base.ModelBaseImpl(&model.FingerPrint{})

	err := fingerPrintModelBase.GetModelListByCondition(&fprints, "cate_id = ?", []interface{}{cateId}...)
	if err != nil {
		//todo
	}

	fprotos := [][]string{}

	for _, fp := range fprints {
		protos := fp.Protos
		//{"DIAMETER":1,"FREE_205":1,"NETFLOW":1,"RADIUS":1,"RX":1,"SOULSEEK":1,"SSDP":1,"STEAM":1,"SYSLOG":1,"t":15}
		var protosMap map[string]uint64
		protoUnmarshalErr := json.Unmarshal([]byte(string(protos)), &protosMap)
		if protoUnmarshalErr != nil {
			continue
		}
		fprotoSlice := []string{}
		for proto, _ := range protosMap {
			fprotoSlice = append(fprotoSlice, proto)
		}

		fprotos = append(fprotos, fprotoSlice)
	}
	return fprotos
}

func GetCateFingerPrintProtoList(cateId string) map[string][]string {
	retMap := map[string][]string{}

	//获取该表下某分类的所有资产指纹
	fingerPrint := &model.FingerPrint{
		CateId: cateId,
	}

	fingerPrintModelBase := model_base.ModelBaseImpl(fingerPrint)

	fingerPrintList := []*model.FingerPrint{}
	err := fingerPrintModelBase.GetModelListByCondition(&fingerPrintList, "cate_id = ?", []interface{}{fingerPrint.CateId}...)
	if err != nil {
		//todo
		return retMap
	}

	cateProtoList := []string{}
	for _, fp := range fingerPrintList {
		protos := fp.Protos
		logger.Logger.Print("%s fingerPrint protos:%+v", util.RunFuncName(), protos)
		logger.Logger.Info("%s fingerPrint protos:%+v", util.RunFuncName(), protos)
		//过滤空字段
		if len(protos) == 0 {
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

		cateProtoMap := map[string]int{}

		for proto, _ := range protosMap {
			cateProtoMap[proto] = 1

		}

		for p, _ := range cateProtoMap {
			cateProtoList = append(cateProtoList, p)
		}
	}
	retMap[cateId] = cateProtoList
	return retMap
}
