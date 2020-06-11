package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

/**
获取某个资产的指纹采集信息
*/
func GetAssetPrintInfos(c *gin.Context) {
	assetId := c.Param("asset_id")
	vehicleId := c.Query("vehicle_id")
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")

	logger.Logger.Info("%s request params vehicle_id:%s,assetId:%s,pageSizeP:%s,pageIndexP:%s",
		util.RunFuncName(), vehicleId, assetId, pageSizeP, pageIndexP)
	logger.Logger.Print("%s request params vehicle_id:%s,assetId:%s,pageSizeP:%s,pageIndexP:%s",
		util.RunFuncName(), vehicleId, assetId, pageSizeP, pageIndexP)

	//参数判断
	argsTrimsEmpty := util.RrgsTrimsEmpty(assetId, vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicle_id:%s,assetId:%s", util.RunFuncName(), vehicleId, assetId)
		logger.Logger.Print("%s argsTrimsEmpty vehicle_id:%s,assetId:%s", util.RunFuncName(), vehicleId, assetId)
		return
	}
	fpageSize, _ := strconv.Atoi(pageSizeP)
	fpageIndex, _ := strconv.Atoi(pageIndexP)
	//默认20
	defaultPageSize := 20
	if fpageSize == 0 {
		fpageSize = defaultPageSize
	}
	//默认第一页
	defaultPageIndex := 1
	if fpageIndex == 0 {
		fpageIndex = defaultPageIndex
	}

	//是否存在
	asset := &model.Asset{
		VehicleId: vehicleId,
		AssetId:   assetId,
	}
	assetModelBase := model_base.ModelBaseImpl(asset)

	err, assetRecordNotFound := assetModelBase.GetModelByCondition(
		"asset_id = ? and vehicle_id = ? ", []interface{}{asset.AssetId, asset.VehicleId}...)

	if assetRecordNotFound {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//主动探测，被动探测连表
	fprintInfoJoinActivePassvie, err := model.GetFprintInfoJoinPassvie(
		"fprint_infos.device_mac = ?", []interface{}{asset.AssetId}...)
	if err != nil {
		//todo
	}

	//被动探测采集信息列表
	fprintPassiveInfos := []*model.FprintPassiveInfo{}
	var total int

	fprintPassiveInfoModelBase := model_base.ModelBaseImplPagination(&model.FprintPassiveInfo{})

	err = fprintPassiveInfoModelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
		&total, &fprintPassiveInfos,
		"fprint_passive_infos.created_at desc", "fprint_info_id = ?",
		[]interface{}{fprintInfoJoinActivePassvie.FprintInfoId}...)

	if err != nil {
		//todo
	}

	fprintInfoJoinActivePassvie.FprintPassiveInfos = fprintPassiveInfos
	fprintInfoJoinActivePassvie.TotalCount = total

	responseData := map[string]interface{}{
		"fprint_info": fprintInfoJoinActivePassvie,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}
func GetPaginationPrintInfos(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")
	startTimeP := c.Query("start_time")
	endTimeP := c.Query("end_time")

	logger.Logger.Info("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)
	logger.Logger.Print("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)

	//argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	//if argsTrimsEmpty {
	//	ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
	//	c.JSON(http.StatusOK, ret)
	//	logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
	//	logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
	//	return
	//}

	fpageSize, _ := strconv.Atoi(pageSizeP)
	fpageIndex, _ := strconv.Atoi(pageIndexP)

	var fStartTime time.Time
	var fEndTime time.Time

	startTime, _ := strconv.Atoi(startTimeP)
	endTime, _ := strconv.Atoi(endTimeP)

	//默认20
	defaultPageSize := 20
	if fpageSize == 0 {
		fpageSize = defaultPageSize
	}
	//默认第一页
	defaultPageIndex := 1
	if fpageIndex == 0 {
		fpageIndex = defaultPageIndex
	}
	//默认2天前
	//defaultStartTime := util.GetFewDayAgo(2) //2
	if startTime == 0 {
		//fStartTime = defaultStartTime
		fStartTime = util.StampUnix2Time(int64(0))
	} else {
		fStartTime = util.StampUnix2Time(int64(startTime))
	}

	//默认当前时间
	defaultEndTime := time.Now()
	if endTime == 0 {
		fEndTime = defaultEndTime
	} else {
		fEndTime = util.StampUnix2Time(int64(endTime))
	}

	logger.Logger.Info("%s frequest params vehicle_id:%s,fpageSize:%s,fpageIndex:%s,fStartTime%s,fEndTime%s",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)
	logger.Logger.Print("%s frequest params vehicle_id:%s,fpageSize:%s,fpageIndex:%s,fStartTime%s,fEndTime%s",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)

	vehicleAssetFprints := []*model.FprintInfo{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.FprintInfo{})

	var query string
	var args []interface{}
	vehicleIdTrimsEmpty := util.RrgsTrim(vehicleId)
	if vehicleIdTrimsEmpty == "" {
		query = "fprint_infos.created_at BETWEEN ? AND ?"
		args = []interface{}{fStartTime, fEndTime}
	} else {
		query = "vehicle_id = ? and fprint_infos.created_at BETWEEN ? AND ?"
		args = []interface{}{vehicleId, fStartTime, fEndTime}
	}

	err := modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
		&total, &vehicleAssetFprints, "fprint_infos.created_at desc", query, args...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"fprint_infos": vehicleAssetFprints,
		"total_count":  total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*
获取资产指纹
*/
//const FprintCount = 10
//
//func GetAssetFprints(c *gin.Context) {
//	vehicleId := c.Query("vehicle_id")
//
//	logger.Logger.Info("%s request params vehicle_id:%s", util.RunFuncName(), vehicleId)
//	logger.Logger.Print("%s request params vehicle_id:%s", util.RunFuncName(), vehicleId)
//
//	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
//	if argsTrimsEmpty {
//		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
//		c.JSON(http.StatusOK, ret)
//		logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
//		logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
//		return
//	}
//
//	//标签库的个数
//	fprintsMacs := []string{}
//	_ = mysql.QueryPluckByModelWhere(&model.FingerPrint{}, "device_mac", &fprintsMacs,
//		"", []interface{}{}...)
//
//	if len(fprintsMacs) == 0 {
//		fprintsMacs = []string{""}
//	}
//
//	//临时
//	fTemp := []string{}
//	for _, v := range fprintsMacs {
//		if v != "" {
//			fTemp = append(fTemp, v)
//		}
//	}
//	var needInsertFprintCount = FprintCount - len(fTemp)
//
//	vehicleAssetFprints := []*model.FprintInfo{}
//	var err error
//	var total int
//	if needInsertFprintCount > 0 {
//		modelBase := model_base.ModelBaseImplPagination(&model.FprintInfo{})
//
//		var fpageSize = needInsertFprintCount
//		var fpageIndex = 1
//
//		err = modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
//			&total, &vehicleAssetFprints, "fprint_infos.created_at asc",
//			"vehicle_id = ? and fprint_infos.device_mac not in (?) and fprint_infos.trade_mark is not null",
//			[]interface{}{vehicleId, fprintsMacs}...)
//	}
//
//	if err != nil {
//		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
//		c.JSON(http.StatusOK, ret)
//		return
//	}
//
//	responseData := map[string]interface{}{
//		"asset_fprints": vehicleAssetFprints,
//	}
//
//	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
//	c.JSON(http.StatusOK, retObj)
//}

/**
入网审批
*/
//
//func GetExamineAssetFprints(c *gin.Context) {
//	vehicleId := c.Query("vehicle_id")
//	pageSizeP := c.Query("page_size")
//	pageIndexP := c.Query("page_index")
//	startTimeP := c.Query("start_time")
//	endTimeP := c.Query("end_time")
//
//	logger.Logger.Info("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
//		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)
//	logger.Logger.Print("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
//		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)
//
//	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
//	if argsTrimsEmpty {
//		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
//		c.JSON(http.StatusOK, ret)
//		logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
//		logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
//		return
//	}
//
//	fpageSize, _ := strconv.Atoi(pageSizeP)
//	fpageIndex, _ := strconv.Atoi(pageIndexP)
//
//	var fStartTime time.Time
//	var fEndTime time.Time
//
//	startTime, _ := strconv.Atoi(startTimeP)
//	endTime, _ := strconv.Atoi(endTimeP)
//
//	//默认20
//	defaultPageSize := 20
//	if fpageSize == 0 {
//		fpageSize = defaultPageSize
//	}
//	//默认第一页
//	defaultPageIndex := 1
//	if fpageIndex == 0 {
//		fpageIndex = defaultPageIndex
//	}
//	//默认2天前
//	//defaultStartTime := util.GetFewDayAgo(2) //2
//	if startTime == 0 {
//		fStartTime = util.StampUnix2Time(int64(0))
//		//fStartTime = defaultStartTime
//	} else {
//		fStartTime = util.StampUnix2Time(int64(startTime))
//	}
//
//	//默认当前时间
//	defaultEndTime := time.Now()
//	if endTime == 0 {
//		fEndTime = defaultEndTime
//	} else {
//		fEndTime = util.StampUnix2Time(int64(endTime))
//	}
//
//	logger.Logger.Info("%s frequest params vehicle_id:%s,fpageSize:%s,fpageIndex:%s,fStartTime%s,fEndTime%s",
//		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)
//	logger.Logger.Print("%s frequest params vehicle_id:%s,fpageSize:%s,fpageIndex:%s,fStartTime%s,fEndTime%s",
//		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)
//	//查找指纹库所有的mac
//	fprintsMacs := []string{}
//	_ = mysql.QueryPluckByModelWhere(&model.FingerPrint{}, "device_mac", &fprintsMacs,
//		"", []interface{}{}...)
//
//	if len(fprintsMacs) == 0 {
//		fprintsMacs = []string{""}
//	}
//	//
//	var totalCount int
//	////终端-策略
//	vehicleAssetFprints := []*model.FprintInfo{}
//	modelBase := model_base.ModelBaseImplPagination(&model.FprintInfo{})
//
//	err := modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
//		&totalCount, &vehicleAssetFprints, "fprint_infos.created_at desc",
//		"fprint_infos.vehicle_id = ? and fprint_infos.examine_net is null and fprint_infos.created_at BETWEEN ? AND ?",
//		[]interface{}{vehicleId, fStartTime, fEndTime}...)
//
//	if len(vehicleAssetFprints) == 0 {
//		logger.Logger.Error("%s vehicle_id:%s,vehicleAssetFprints null", util.RunFuncName(), vehicleId)
//		logger.Logger.Print("%s vehicle_id:%s,vehicleAssetFprints null", util.RunFuncName(), vehicleId)
//		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsUnExistMsg, "")
//		c.JSON(http.StatusOK, ret)
//		return
//	}
//	if err != nil {
//		logger.Logger.Error("%s vehicle_id:%s,vehicleAssetFprints err:%+v", util.RunFuncName(), vehicleId, err)
//		logger.Logger.Print("%s vehicle_id:%s,vehicleAssetFprints err:%+v", util.RunFuncName(), vehicleId, err)
//		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
//		c.JSON(http.StatusOK, ret)
//		return
//	}
//
//	responseData := map[string]interface{}{
//		"asset_fprints": vehicleAssetFprints,
//		"total_count":   totalCount,
//	}
//
//	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
//	c.JSON(http.StatusOK, retObj)
//}

/**
入网审批
*/
func AddExamineAssetFprints(c *gin.Context) {
	assetId := c.Param("asset_id")
	vehicleId := c.PostForm("vehicle_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(assetId, vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty assetId:%s", util.RunFuncName(), assetId)
		logger.Logger.Print("%s argsTrimsEmpty assetId:%s", util.RunFuncName(), assetId)
		return
	}
	logger.Logger.Info("%s request params assetId:%s", util.RunFuncName(), assetId)
	logger.Logger.Print("%s request params assetId:%s", util.RunFuncName(), assetId)

	////////资产是否存在////////
	assetInfo := &model.Asset{
		AssetId: assetId,
	}
	modelBase := model_base.ModelBaseImpl(assetInfo)

	err, recordNotFound := modelBase.GetModelByCondition("asset_id = ?", []interface{}{assetInfo.AssetId}...)

	if err != nil {
		logger.Logger.Error("%s asset_id:%s,err:%s", util.RunFuncName(), assetId, err)
		logger.Logger.Print("%s asset_id:%s,err:%s", util.RunFuncName(), assetId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s asset_id:%s,recordNotFound", util.RunFuncName(), assetId)
		logger.Logger.Print("%s asset_id:%s,recordNotFound", util.RunFuncName(), assetId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//////////是否已经识别////////

	assetFprint := &model.FingerPrint{
		DeviceMac: assetId,
	}
	assetFprintInfoModelBase := model_base.ModelBaseImpl(assetFprint)

	err, fprintInfoRecordNotFound := assetFprintInfoModelBase.GetModelByCondition("device_mac = ?", []interface{}{assetFprint.DeviceMac}...)
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s asset_id:%s,%s", util.RunFuncName(), assetId, response.ReqGetAssetFprintsFailMsg)
		logger.Logger.Print("%s asset_id:%s,%s", util.RunFuncName(), assetId, response.ReqGetAssetFprintsFailMsg)
		return
	}
	if !fprintInfoRecordNotFound {
		//连表查询
		assetJoinFprintJoinCategory, _ := model.GetAssetJoinFprintJoinCategory(
			"assets.asset_id = ? and assets.vehicle_id = ?", []interface{}{assetId, vehicleId}...)
		responseData := map[string]interface{}{
			"asset": assetJoinFprintJoinCategory,
		}
		ret := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsUnExistMsg, responseData)
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s asset_id:%s,recordNotFound", util.RunFuncName(), assetId)
		logger.Logger.Print("%s asset_id:%s,recordNotFound", util.RunFuncName(), assetId)
		return
	}

	//////////如果不存在，通过协议种类，占比去分析////////

	//tMark := assetFprintInfo.TradeMark
	//deviceMac := assetFprintInfo.DeviceMac
	//
	////搜索指纹库有没有记录
	//fingerPrint := &model.FingerPrint{
	//	DeviceMac: deviceMac,
	//}
	//fprintModelBase := model_base.ModelBaseImpl(fingerPrint)
	//err, fprintRecordNotFound := fprintModelBase.GetModelByCondition("device_mac = ?", []interface{}{fingerPrint.DeviceMac}...)
	//
	//assetFprintInfo := &model.FprintInfo{
	//	FprintInfoId: assetFprintId,
	//}
	//assetFprintInfoModelBase := model_base.ModelBaseImpl(assetFprintInfo)
	//
	//_, _ = assetFprintInfoModelBase.GetModelByCondition("fprint_info_id = ?", []interface{}{assetFprintInfo.FprintInfoId}...)
	//responseData := map[string]interface{}{
	//	"asset_fprint": assetFprintInfo,
	//}
	//
	//retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
	//c.JSON(http.StatusOK, retObj)
}

/**
注册入网
*/
func AddNetAccessAssetFprints(c *gin.Context) {
	assetFprintId := c.Param("asset_fprint_id")
	netAccessFlag := c.PostForm("access_net_flag")

	argsTrimsEmpty := util.RrgsTrimsEmpty(assetFprintId, netAccessFlag)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty assetFprintId:%s,netAccessFlag%s", util.RunFuncName(), assetFprintId, netAccessFlag)
		logger.Logger.Print("%s argsTrimsEmpty assetFprintId:%s,netAccessFlag%s", util.RunFuncName(), assetFprintId, netAccessFlag)
		return
	}
	logger.Logger.Info("%s argsTrimsEmpty assetFprintId:%s,netAccessFlag%s", util.RunFuncName(), assetFprintId, netAccessFlag)
	logger.Logger.Print("%s argsTrimsEmpty assetFprintId:%s,netAccessFlag%s", util.RunFuncName(), assetFprintId, netAccessFlag)
	fNetAccessFlag := true
	switch netAccessFlag {
	case "true":
		fNetAccessFlag = true
	case "false":
		fNetAccessFlag = false
	}

	//查询是否存在
	assetFprintInfo := &model.FprintInfo{
		FprintInfoId: assetFprintId,
	}
	assetFprintInfoModelBase := model_base.ModelBaseImpl(assetFprintInfo)

	err, fprintInfoRecordNotFound := assetFprintInfoModelBase.GetModelByCondition("fprint_info_id = ?", []interface{}{assetFprintInfo.FprintInfoId}...)
	if fprintInfoRecordNotFound {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//允许入网标识
	attrs := map[string]interface{}{
		"access_net": fNetAccessFlag,
	}
	if err := assetFprintInfoModelBase.UpdateModelsByCondition(attrs, "fprint_info_id = ?",
		[]interface{}{assetFprintId}...); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddFprintsAccessNetFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//查询最新的
	_, _ = assetFprintInfoModelBase.GetModelByCondition("fprint_info_id = ?", []interface{}{assetFprintInfo.FprintInfoId}...)

	responseData := map[string]interface{}{
		"asset_fprint": assetFprintInfo,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddFprintsAccessNetSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}
