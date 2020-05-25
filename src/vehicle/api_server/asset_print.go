package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func GetPaginationAssetFprints(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")
	startTimeP := c.Query("start_time")
	endTimeP := c.Query("end_time")

	logger.Logger.Info("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)
	logger.Logger.Print("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		return
	}

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
	defaultStartTime := util.GetFewDayAgo(2) //2
	if startTime == 0 {
		fStartTime = defaultStartTime
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

	err := modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
		&total, &vehicleAssetFprints, "fprint_infos.created_at desc", "vehicle_id = ? and fprint_infos.created_at BETWEEN ? AND ?",
		[]interface{}{vehicleId, fStartTime, fEndTime}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"asset_fprints": vehicleAssetFprints,
		"total_count":   total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}

/*
获取资产指纹
*/
const FprintCount = 10

func GetAssetFprints(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")

	logger.Logger.Info("%s request params vehicle_id:%s", util.RunFuncName(), vehicleId)
	logger.Logger.Print("%s request params vehicle_id:%s", util.RunFuncName(), vehicleId)

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		return
	}

	//标签库的个数
	fprintsMacs := []string{}
	_ = mysql.QueryPluckByModelWhere(&model.FingerPrint{}, "device_mac", &fprintsMacs,
		"", []interface{}{}...)

	if len(fprintsMacs) == 0 {
		fprintsMacs = []string{""}
	}

	//临时
	fTemp := []string{}
	for _, v := range fprintsMacs {
		if v != "" {
			fTemp = append(fTemp, v)
		}
	}
	var needInsertFprintCount = FprintCount - len(fTemp)

	vehicleAssetFprints := []*model.FprintInfo{}
	var err error
	var total int
	if needInsertFprintCount > 0 {
		modelBase := model_base.ModelBaseImplPagination(&model.FprintInfo{})

		var fpageSize = needInsertFprintCount
		var fpageIndex = 1

		err = modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
			&total, &vehicleAssetFprints, "fprint_infos.created_at asc",
			"vehicle_id = ? and fprint_infos.device_mac not in (?) and fprint_infos.trade_mark is not null",
			[]interface{}{vehicleId, fprintsMacs}...)
	}

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"asset_fprints": vehicleAssetFprints,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

/**
入网审批
*/

func GetExamineAssetFprints(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")
	startTimeP := c.Query("start_time")
	endTimeP := c.Query("end_time")

	logger.Logger.Info("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)
	logger.Logger.Print("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		return
	}

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
	defaultStartTime := util.GetFewDayAgo(2) //2
	if startTime == 0 {
		fStartTime = defaultStartTime
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
	//查找指纹库所有的mac
	fprintsMacs := []string{}
	_ = mysql.QueryPluckByModelWhere(&model.FingerPrint{}, "device_mac", &fprintsMacs,
		"", []interface{}{}...)

	if len(fprintsMacs) == 0 {
		fprintsMacs = []string{""}
	}
	//
	var totalCount int
	////终端-策略
	vehicleAssetFprints := []*model.FprintInfo{}
	modelBase := model_base.ModelBaseImplPagination(&model.FprintInfo{})

	err := modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
		&totalCount, &vehicleAssetFprints, "fprint_infos.created_at desc",
		"fprint_infos.vehicle_id = ? and fprint_infos.examine_net is null and fprint_infos.created_at BETWEEN ? AND ?",
		[]interface{}{vehicleId, fStartTime, fEndTime}...)

	if len(vehicleAssetFprints) == 0 {
		logger.Logger.Error("%s vehicle_id:%s,vehicleAssetFprints null", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,vehicleAssetFprints null", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,vehicleAssetFprints err:%+v", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,vehicleAssetFprints err:%+v", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"asset_fprints": vehicleAssetFprints,
		"total_count":   totalCount,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

/**
入网审批
*/
func AddExamineAssetFprints(c *gin.Context) {

}
