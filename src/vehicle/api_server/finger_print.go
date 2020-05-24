package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

/*
获取资产指纹
*/
func GetAssetFprints(c *gin.Context) {
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

	var totalCount int
	//终端-策略
	vehicleAssetFprints, err := model.GetPaginAssetFprints(fpageIndex, fpageSize, &totalCount,
		"fprint_detect_infos.vehicle_id = ? and fprint_detect_infos.trade_mark IS NOT null and fprint_detect_infos.created_at BETWEEN ? AND ?", []interface{}{vehicleId, fStartTime, fEndTime}...)

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
添加指纹库
*/

func AddFprint(c *gin.Context) {
	assetIds := c.PostForm("asset_ids")
	cateId := c.PostForm("cate_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(assetIds, cateId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s assetIds:%s,cateId%s", util.RunFuncName(), assetIds, cateId)
		logger.Logger.Error("%s assetIds:%s,cateId%s", util.RunFuncName(), assetIds, cateId)
		return
	}

	logger.Logger.Print("%s assetIds:%s,cateId%s", util.RunFuncName(), assetIds, cateId)
	logger.Logger.Info("%s assetIds:%s,cateId%s", util.RunFuncName(), assetIds, cateId)

	assetIdSlice := strings.Split(assetIds, ",")
	//todo
	//asset_ids没有过滤

	var insertFprintIds []string
	for _, assetId := range assetIdSlice {
		fingerPrint := &model.FingerPrint{
			FprintId:  util.RandomString(32),
			CateId:    cateId,
			VehicleId: "",
			DeviceMac: assetId,
		}
		fingerPrintModelBase := model_base.ModelBaseImpl(fingerPrint)

		err, fingerPrintRecordNotFound := fingerPrintModelBase.GetModelByCondition("device_mac = ? and cate_id = ?",
			[]interface{}{fingerPrint.DeviceMac, fingerPrint.CateId}...)

		if fingerPrintRecordNotFound {
			if err = fingerPrintModelBase.InsertModel(); err != nil {
				continue
			} else {
				insertFprintIds = append(insertFprintIds, fingerPrint.DeviceMac)
			}
		} else {
			//todo
		}
	}

	fingerPrintInsertList := []*model.FingerPrint{}
	fingerPrintModelBase := model_base.ModelBaseImpl(&model.FingerPrint{})
	err := fingerPrintModelBase.GetModelListByCondition(&fingerPrintInsertList,
		"device_mac in (?)", []interface{}{insertFprintIds}...)

	if err != nil {
		retObj := response.StructResponseObj(response.VStatusServerError, response.ReqGetFprintsFailMsg, "")
		c.JSON(http.StatusOK, retObj)
		return
	}
	responseContent := map[string]interface{}{}
	responseContent["fprints"] = fingerPrintInsertList

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFprintsSuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}

/**
查询所有指纹库
*/

func GetFprints(c *gin.Context) {

	cateModelBase := model_base.ModelBaseImpl(&model.Category{})
	cates := []*model.Category{}
	err := cateModelBase.GetModelListByCondition(&cates, "", []interface{}{}...)
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseContent := map[string]interface{}{}
	responseContent["categorys"] = cates

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqCategoryListSuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}

/**
删除指纹库
*/

func DeleFprint(c *gin.Context) {
	cateId := c.Param("cate_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(cateId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId)
		logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId)
		return
	}

	logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId)
	logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId)

}

/**
编辑指纹库
*/

func EditFprint(c *gin.Context) {
	cateId := c.Param("cate_id")
	cateName := c.PostForm("cate_name")

	argsTrimsEmpty := util.RrgsTrimsEmpty(cateId, cateName)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
		logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
		return
	}

	logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
	logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)

	cate := &model.Category{
		CateId: cateId,
	}
	cateModelBase := model_base.ModelBaseImpl(cate)

	err, cateRecordNotFound := cateModelBase.GetModelByCondition("cate_id = ?", []interface{}{cate.CateId}...)
	if cateRecordNotFound {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryNotExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//编辑
	attrs := map[string]interface{}{
		"name": cateName,
	}
	if err := cateModelBase.UpdateModelsByCondition(attrs, "cate_id = ?",
		[]interface{}{cate.CateId}...); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateCategoryFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//ReqUpdateCategorySuccessMsg
	cateUpdated := &model.Category{
		CateId: cateId,
	}
	cateUpdatedModelBase := model_base.ModelBaseImpl(cateUpdated)
	_, _ = cateUpdatedModelBase.GetModelByCondition("cate_id = ?", []interface{}{cateUpdated.CateId}...)

	responseContent := map[string]interface{}{}
	responseContent["category"] = cateUpdated

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateCategorySuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}
