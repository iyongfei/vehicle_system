package api_server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func EditAssetInfo(c *gin.Context) {
	assetId := c.Param("asset_id")
	vehicleId := c.PostForm("vehicle_id")
	name := c.PostForm("name")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, assetId, name)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
	}

	//查询是否存在
	assetInfo := &model.Asset{
		VehicleId: vehicleId,
		AssetId:   assetId,
	}
	modelBase := model_base.ModelBaseImpl(assetInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ? and asset_id = ?",
		[]interface{}{assetInfo.VehicleId, assetInfo.AssetId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//更新名字

	//编辑
	attrs := map[string]interface{}{
		"name": name,
	}
	if err := modelBase.UpdateModelsByCondition(attrs, "vehicle_id = ? and asset_id = ?",
		[]interface{}{assetInfo.VehicleId, assetInfo.AssetId}...); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateAssetFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	_, _ = modelBase.GetModelByCondition("asset_id = ?", []interface{}{assetInfo.AssetId}...)

	assetFprintCateJoin, _ := model.GetAssetFprintCateJoin("asset_id = ?", []interface{}{assetInfo.AssetId}...)

	AssetJoinFprintJoinCategory := model.AssetJoinFprintJoinCategory{
		Asset:    assetInfo,
		CateId:   assetFprintCateJoin.CateId,
		CateName: assetFprintCateJoin.CateName,
	}

	responseContent := map[string]interface{}{}
	responseContent["asset"] = AssetJoinFprintJoinCategory

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetSuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}
func EditAsset(c *gin.Context) {
	assetId := c.Param("asset_id")
	setTypeP := c.PostForm("type")
	setSwitchP := c.PostForm("switch")
	cateId := c.PostForm("cate_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(assetId, setTypeP, setSwitchP)

	//setType
	/**
	var DeviceSetParam_Type_name = map[int32]string{
	0: "DEFAULT",
	1: "PROTECT",
	2: "INTERNET",
	3: "GUEST_ACCESS_DEVICE",
	4: "LANVISIT",}
	*/
	var setTypeYes bool
	types := protobuf.DeviceSetParam_Type_name

	for k, _ := range types {
		kstr := strconv.Itoa(int(k))
		trimSetTypeP := util.RrgsTrim(setTypeP)
		if kstr == trimSetTypeP {
			setTypeYes = true
		}
	}

	//swith
	var setSwitchYes bool
	switchSlice := []string{"true", "false"}

	trimSetSwitchP := util.RrgsTrim(setSwitchP)
	if util.IsExistInSlice(strings.ToLower(trimSetSwitchP), switchSlice) {
		setSwitchYes = true
		trimSetSwitchP = strings.ToLower(trimSetSwitchP)
	}

	if !setTypeYes || !setSwitchYes || argsTrimsEmpty { //类型错误
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
		return

	}
	setType, _ := strconv.Atoi(setTypeP)

	setSwitch := true
	switch trimSetSwitchP {
	case "true":
		setSwitch = true
	case "false":
		setSwitch = false
	}

	//查询是否存在
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
	//判断类别是否存在
	if !util.RrgsTrimEmpty(cateId) {
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

		//贴标签
		assetPrint := &model.AssetFprint{
			AssetFprintId: util.RandomString(32),
			AssetId:       assetId,
			CateId:        cateId,
		}
		assetPrintModelBase := model_base.ModelBaseImpl(assetPrint)

		err, assetPrintRecordNotFound := assetPrintModelBase.GetModelByCondition("asset_id = ?", []interface{}{assetPrint.AssetId}...)

		if assetPrintRecordNotFound {
			//贴标签
			err = assetPrintModelBase.InsertModel()
			if err != nil {
				logger.Logger.Print("%s asset_id%s,get asset err:%+v", util.RunFuncName(), assetPrint.AssetId, err)
				logger.Logger.Error("%s asset_id%s,get asset err:%+v", util.RunFuncName(), assetPrint.AssetId, err)

				ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddAssetFprintsCateFailMsg, "")
				c.JSON(http.StatusOK, ret)
				return
			}
		} else {
			assetPrint.CateId = cateId
			//编辑
			attrs := map[string]interface{}{
				"cate_id": assetPrint.CateId,
			}
			if err := assetPrintModelBase.UpdateModelsByCondition(attrs, "asset_id = ?",
				[]interface{}{assetPrint.AssetId}...); err != nil {
				ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateCategoryFailMsg, "")
				c.JSON(http.StatusOK, ret)
				return
			}
		}
	}

	switch setType {
	case int(protobuf.DeviceSetParam_PROTECT):

		attrs := map[string]interface{}{
			"access_net": setSwitch,
		}
		if err := assetInfo.UpdateModelsByCondition(attrs, "asset_id = ?",
			[]interface{}{assetInfo.AssetId}...); err != nil {
			ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateAssetFailMsg, "")
			c.JSON(http.StatusOK, ret)
			return
		}

		break
	case int(protobuf.DeviceSetParam_INTERNET):
		break
	case int(protobuf.DeviceSetParam_GUEST_ACCESS_DEVICE):
		break
	case int(protobuf.DeviceSetParam_LANVISIT):
		break

	}

	//更新
	//assetCmd := &emq_cmd.AssetSetCmd{
	//	VehicleId: assetInfo.VehicleId,
	//	TaskType:  int(protobuf.Command_DEVICE_SET),
	//
	//	Switch: setSwitch,
	//	Type:   setType,
	//	Mac:    assetId,
	//}
	//
	//topic_publish_handler.GetPublishService().PutMsg2PublicChan(assetCmd)

	_, _ = modelBase.GetModelByCondition("asset_id = ?", []interface{}{assetInfo.AssetId}...)

	assetFprintCateJoin, _ := model.GetAssetFprintCateJoin("asset_id = ?", []interface{}{assetInfo.AssetId}...)

	AssetJoinFprintJoinCategory := model.AssetJoinFprintJoinCategory{
		Asset:    assetInfo,
		CateId:   assetFprintCateJoin.CateId,
		CateName: assetFprintCateJoin.CateName,
	}

	responseContent := map[string]interface{}{}
	responseContent["asset"] = AssetJoinFprintJoinCategory

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateAssetSuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}

/**
获取所有的资产设备
*/
func GetAllAssets(c *gin.Context) {

	var pageIndex = 1
	var pageSize = 1000
	assetInfos := []*model.Asset{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.Asset{})

	err := modelBase.GetModelPaginationByCondition(pageIndex, pageSize,
		&total, &assetInfos, "assets.created_at desc", "",
		[]interface{}{}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"assets":      assetInfos,
		"total_count": total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func GetPaginationAssets(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")
	startTimeP := c.Query("start_time")
	endTimeP := c.Query("end_time")

	logger.Logger.Info("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)
	logger.Logger.Print("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)

	vehicleId = util.RrgsTrim(vehicleId)
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

	assets := []*model.Asset{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.Asset{})

	var sqlQuery string
	var sqlArgs []interface{}

	if vehicleId == "" {
		sqlQuery = "assets.created_at BETWEEN ? AND ?"
		sqlArgs = append(sqlArgs, fStartTime, fEndTime)
	} else {
		sqlQuery = "vehicle_id = ? and assets.created_at BETWEEN ? AND ?"
		sqlArgs = append(sqlArgs, vehicleId, fStartTime, fEndTime)
	}

	err := modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
		&total, &assets, "assets.created_at desc", sqlQuery,
		sqlArgs...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"assets":      assets,
		"total_count": total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func GetAsset(c *gin.Context) {
	assetId := c.Param("asset_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(assetId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), assetId)
		logger.Logger.Print("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), assetId)
	}

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
		logger.Logger.Error("%s assetId:%s,recordNotFound", util.RunFuncName(), assetId)
		logger.Logger.Print("%s assetId:%s,recordNotFound", util.RunFuncName(), assetId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	_, _ = modelBase.GetModelByCondition("asset_id = ?", []interface{}{assetInfo.AssetId}...)

	assetFprintCateJoin, _ := model.GetAssetFprintCateJoin("asset_id = ?", []interface{}{assetInfo.AssetId}...)

	AssetJoinFprintJoinCategory := model.AssetJoinFprintJoinCategory{
		Asset:    assetInfo,
		CateId:   assetFprintCateJoin.CateId,
		CateName: assetFprintCateJoin.CateName,
	}

	responseContent := map[string]interface{}{}
	responseContent["asset"] = AssetJoinFprintJoinCategory

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetSuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)

}

//1、验证（录入）
//2、不验证
// 2-1不非法 在
// 2-2非法 不在白名单
/**
获取所有的资产白名单
*/

/**
添加
*/

func AddAsset(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	assetInfo := &model.Asset{}
	err := json.Unmarshal(body, assetInfo)

	if err != nil {
		logger.Logger.Error("%s unmarshal assetInfo err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Print("%s unmarshal assetInfo err:%s", util.RunFuncName(), err.Error())
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	modelBase := model_base.ModelBaseImpl(assetInfo)

	err, recordNotFound := modelBase.GetModelByCondition("asset_id = ?", []interface{}{assetInfo.AssetId}...)

	if !recordNotFound {
		logger.Logger.Error("%s asset_id:%s exist", util.RunFuncName(), assetInfo.AssetId)
		logger.Logger.Print("%s asset_id:%s exist", util.RunFuncName(), assetInfo.AssetId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if err := modelBase.InsertModel(); err != nil {
		logger.Logger.Error("%s add asset_id:%s err:%s", util.RunFuncName(), assetInfo.AssetId, err.Error())
		logger.Logger.Print("%s add asset_id:%s err:%s", util.RunFuncName(), assetInfo.AssetId, err.Error())
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddAssetFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"asset": assetInfo,
	}

	ret := response.StructResponseObj(response.VStatusOK, response.ReqAddAssetSuccessMsg, responseData)
	c.JSON(http.StatusOK, ret)
}

func DeleAsset(c *gin.Context) {
	assetId := c.Param("asset_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(assetId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty assetId:%s argsTrimsEmpty", util.RunFuncName(), assetId)
		logger.Logger.Print("%s argsTrimsEmpty assetId:%s argsTrimsEmpty", util.RunFuncName(), assetId)
		return
	}

	assetObj := &model.Asset{
		AssetId: assetId,
	}

	modelBase := model_base.ModelBaseImpl(assetObj)
	err, recordNotFound := modelBase.GetModelByCondition("asset_id = ?", []interface{}{assetObj.AssetId}...)

	if err != nil {
		logger.Logger.Error("%s asset_id:%s err:%s", util.RunFuncName(), assetObj.AssetId, err)
		logger.Logger.Print("%s asset_id:%s err:%s", util.RunFuncName(), assetObj.AssetId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleAssetFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if recordNotFound {
		logger.Logger.Error("%s asset_id:%s,record not exist", util.RunFuncName(), assetObj.AssetId)
		logger.Logger.Print("%s asset_id:%s,record not exist", util.RunFuncName(), assetObj.AssetId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err := modelBase.DeleModelsByCondition("asset_id = ?",
		[]interface{}{assetObj.AssetId}...); err != nil {
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqDeleAssetSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}
