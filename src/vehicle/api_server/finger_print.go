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

	//更新资产指纹信息

	fingerPrintInsertList := []*model.FingerPrint{}
	fingerPrintModelBase := model_base.ModelBaseImpl(&model.FingerPrint{})
	err := fingerPrintModelBase.GetModelListByCondition(&fingerPrintInsertList,
		"device_mac in (?)", []interface{}{insertFprintIds}...)

	if err != nil {
		retObj := response.StructResponseObj(response.VStatusServerError, response.ReqAddFprintsFailMsg, "")
		c.JSON(http.StatusOK, retObj)
		return
	}
	responseContent := map[string]interface{}{}
	responseContent["fprints"] = fingerPrintInsertList

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddFprintsSuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}

/**
查询所有指纹库
*/

func GetFprints(c *gin.Context) {
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

	logger.Logger.Info("%s frequest params vehicle_id:%s,fpageSize:%d,fpageIndex:%d,fStartTime%s,fEndTime%s",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)
	logger.Logger.Print("%s frequest params vehicle_id:%s,fpageSize:%d,fpageIndex:%d,fStartTime%s,fEndTime%s",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)

	//////////////////
	fprints := []*model.FingerPrint{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.FingerPrint{})

	err := modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
		&total, &fprints, "finger_prints.created_at desc", "vehicle_id = ? and finger_prints.created_at BETWEEN ? AND ?",
		[]interface{}{vehicleId, fStartTime, fEndTime}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	responseContent := map[string]interface{}{}
	responseContent["fprints"] = fprints

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFprintsSuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}

/**
删除指纹库
*/

func DeleFprint(c *gin.Context) {
	fprintId := c.Param("fprint_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(fprintId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s fprintId%s", util.RunFuncName(), fprintId)
		logger.Logger.Error("%s fprintId%s", util.RunFuncName(), fprintId)
		return
	}

	logger.Logger.Print("%s fprintId%s", util.RunFuncName(), fprintId)
	logger.Logger.Error("%s fprintId%s", util.RunFuncName(), fprintId)

	fprint := &model.FingerPrint{
		FprintId: fprintId,
	}

	fprintModelBase := model_base.ModelBaseImpl(fprint)

	err, fprintRecordNotFound := fprintModelBase.GetModelByCondition("fprint_id = ?", []interface{}{fprint.FprintId}...)
	if fprintRecordNotFound {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFprintsUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	err = fprintModelBase.DeleModelsByCondition("fprint_id = ?", []interface{}{fprint.FprintId}...)

	if err != nil {
		logger.Logger.Error("%s fprintId:%s err:%s", util.RunFuncName(), fprintId, err)
		logger.Logger.Print("%s fprintId:%s err:%s", util.RunFuncName(), fprintId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqDeleFprintsSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}

/**
编辑指纹库
*/

func EditFprint(c *gin.Context) {
	//cateId := c.Param("cate_id")
	//cateName := c.PostForm("cate_name")
	//
	//argsTrimsEmpty := util.RrgsTrimsEmpty(cateId, cateName)
	//if argsTrimsEmpty {
	//	ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
	//	c.JSON(http.StatusOK, ret)
	//
	//	logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
	//	logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
	//	return
	//}
	//
	//logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
	//logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
	//
	//cate := &model.Category{
	//	CateId: cateId,
	//}
	//cateModelBase := model_base.ModelBaseImpl(cate)
	//
	//err, cateRecordNotFound := cateModelBase.GetModelByCondition("cate_id = ?", []interface{}{cate.CateId}...)
	//if cateRecordNotFound {
	//	ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryNotExistMsg, "")
	//	c.JSON(http.StatusOK, ret)
	//	return
	//}
	//if err != nil {
	//	ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryFailMsg, "")
	//	c.JSON(http.StatusOK, ret)
	//	return
	//}
	////编辑
	//attrs := map[string]interface{}{
	//	"name": cateName,
	//}
	//if err := cateModelBase.UpdateModelsByCondition(attrs, "cate_id = ?",
	//	[]interface{}{cate.CateId}...); err != nil {
	//	ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateCategoryFailMsg, "")
	//	c.JSON(http.StatusOK, ret)
	//	return
	//}
	////ReqUpdateCategorySuccessMsg
	//cateUpdated := &model.Category{
	//	CateId: cateId,
	//}
	//cateUpdatedModelBase := model_base.ModelBaseImpl(cateUpdated)
	//_, _ = cateUpdatedModelBase.GetModelByCondition("cate_id = ?", []interface{}{cateUpdated.CateId}...)
	//
	//responseContent := map[string]interface{}{}
	//responseContent["category"] = cateUpdated
	//
	//retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateCategorySuccessMsg, responseContent)
	//c.JSON(http.StatusOK, retObj)
}
