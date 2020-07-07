package api_server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
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

	assetFprintCateJoin, _ := model.GetAssetFprintCateJoin("asset_fprints.asset_id = ?", []interface{}{assetInfo.AssetId}...)
	fprintJoinAsset, _ := model.GetFprintJoinAsset("fprints.asset_id = ?", []interface{}{assetInfo.AssetId}...)

	if fprintJoinAsset.AutoCateName == "" {
		fprintJoinAsset.AutoCateName = response.UnKnow
	}
	AssetJoinFprintJoinCategory := model.AssetJoinFprintJoinCategory{
		Asset:    assetInfo,
		CateId:   assetFprintCateJoin.CateId,
		CateName: assetFprintCateJoin.CateName,

		AutoCateId:   fprintJoinAsset.AutoCateId,
		AutoCateName: fprintJoinAsset.AutoCateName,
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

	assetFprintCateJoin, _ := model.GetAssetFprintCateJoin("asset_fprints.asset_id = ?", []interface{}{assetInfo.AssetId}...)
	fprintJoinAsset, _ := model.GetFprintJoinAsset("fprints.asset_id = ?", []interface{}{assetInfo.AssetId}...)

	if fprintJoinAsset.AutoCateName == "" {
		fprintJoinAsset.AutoCateName = response.UnKnow
	}

	AssetJoinFprintJoinCategory := model.AssetJoinFprintJoinCategory{
		Asset:    assetInfo,
		CateId:   assetFprintCateJoin.CateId,
		CateName: assetFprintCateJoin.CateName,

		AutoCateId:   fprintJoinAsset.AutoCateId,
		AutoCateName: fprintJoinAsset.AutoCateName,
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

	logger.Logger.Info("%s frequest params vehicle_id:%s,fpageSize:%d,fpageIndex:%d,fStartTime%s,fEndTime%s",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)
	logger.Logger.Print("%s frequest params vehicle_id:%s,fpageSize:%d,fpageIndex:%d,fStartTime%s,fEndTime%s",
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

	//assetId列表
	assetIds := []string{}
	for _, v := range assets {
		assetIds = append(assetIds, v.AssetId)
	}
	if len(assetIds) == 0 {
		assetIds = append(assetIds, "")
	}

	//类别识别
	assetFprintCateJoin, _ := model.GetAssetFprintCateListJoin("asset_fprints.asset_id in (?)", []interface{}{assetIds}...)
	fprintJoinAsset, _ := model.GetFprintJoinAssetList("fprints.asset_id in (?)", []interface{}{assetIds}...)

	AssetJoinFprintJoinCategorys := []*model.AssetJoinFprintJoinCategoryTmp{}

	for _, asset := range assets {
		assetJoinFprintJoinCategory := &model.AssetJoinFprintJoinCategoryTmp{}

		model.CreateAssetJoinFprintJoinCategoryTmp(assetJoinFprintJoinCategory, asset)

		assetId := asset.AssetId
		isInassetFprintCateJoin := false
		for _, assetFprint := range assetFprintCateJoin {
			assetFprintAssetId := assetFprint.AssetId
			if assetFprintAssetId == assetId {
				assetJoinFprintJoinCategory.CateId = assetFprint.CateId
				assetJoinFprintJoinCategory.CateName = assetFprint.CateName
				isInassetFprintCateJoin = true
			}
		}
		if !isInassetFprintCateJoin {
			assetJoinFprintJoinCategory.CateId = ""
			assetJoinFprintJoinCategory.CateName = ""
		}
		isInfprintJoinAsset := false
		for _, fprintAsset := range fprintJoinAsset {
			fprintAssetId := fprintAsset.AssetId
			if fprintAssetId == assetId {
				assetJoinFprintJoinCategory.AutoCateId = fprintAsset.AutoCateId
				assetJoinFprintJoinCategory.AutoCateName = fprintAsset.AutoCateName
				isInfprintJoinAsset = true
			}
		}
		if !isInfprintJoinAsset {
			assetJoinFprintJoinCategory.AutoCateId = ""
			assetJoinFprintJoinCategory.AutoCateName = response.UnKnow
		}
		AssetJoinFprintJoinCategorys = append(AssetJoinFprintJoinCategorys, assetJoinFprintJoinCategory)
	}

	responseData := map[string]interface{}{
		"assets":      AssetJoinFprintJoinCategorys,
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
		return
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

	assetFprintCateJoin, _ := model.GetAssetFprintCateJoin("asset_fprints.asset_id = ?", []interface{}{assetInfo.AssetId}...)
	fprintJoinAsset, _ := model.GetFprintJoinAsset("fprints.asset_id = ?", []interface{}{assetInfo.AssetId}...)

	if fprintJoinAsset.AutoCateName == "" {
		fprintJoinAsset.AutoCateName = response.UnKnow
	}

	AssetJoinFprintJoinCategory := model.AssetJoinFprintJoinCategory{
		Asset:    assetInfo,
		CateId:   assetFprintCateJoin.CateId,
		CateName: assetFprintCateJoin.CateName,

		AutoCateId:   fprintJoinAsset.AutoCateId,
		AutoCateName: fprintJoinAsset.AutoCateName,
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

/**
获取资产指纹
legend: {
        data: ['风电电量', '光伏电量', '完整度']
    }

xAxis: ['2017', '2018', '2019', '2020']
series:{
"风电电量":[1,2,3,4]
"光伏电量":[]
"完整度":[]
}
*/
func GetAssetFprint(c *gin.Context) {
	//协议种类
	assetId := c.Param("asset_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(assetId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), assetId)
		logger.Logger.Print("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), assetId)
		return
	}

	////////////////////////////////取出所有的proto////////////////////////////
	fprintFlows := []*model.FprintFlow{}

	fprintFlow := &model.FprintFlow{
		AssetId: assetId,
	}

	fprintFlowModelBase := model_base.ModelBaseImpl(fprintFlow)

	err := fprintFlowModelBase.GetModelListByCondition(&fprintFlows,
		"asset_id = ?", []interface{}{fprintFlow.AssetId}...)

	if err != nil {
		logger.Logger.Error("%s asset_id:%s err:%s", util.RunFuncName(), assetId, err)
		logger.Logger.Print("%s asset_id:%s err:%s", util.RunFuncName(), assetId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	protoMap := map[string]float64{}

	for _, fpFlow := range fprintFlows {
		protocolStr := protobuf.GetFlowProtocols(int(fpFlow.Protocol))
		protoMap[protocolStr] = 1
	}

	if len(protoMap) == 0 {
		//todo
	}
	protoMap[response.FPrintWholeRate] = 1

	////////////////////////////计算时间跨度///////////////////////////////
	firstPrintFlow := &model.FprintFlow{
		AssetId: assetId,
	}
	err = mysql.QueryModelOneRecordByWhereCondition(firstPrintFlow,
		"asset_id = ?", []interface{}{firstPrintFlow.AssetId}...)

	if err != nil {
		//todo
	}
	LstfprintFlow := &model.FprintFlow{
		AssetId: assetId,
	}

	err = mysql.QueryModelLstOneRecordByWhereCondition(LstfprintFlow,
		"asset_id = ?", []interface{}{LstfprintFlow.AssetId}...)
	if err != nil {
		//todo
	}
	distanceTime := util.TimeSpace(firstPrintFlow.CreatedAt, LstfprintFlow.CreatedAt)

	crossPart := distanceTime / 5

	// | | | | | |
	////////////////////////计算每个时间段协议的占比//////////////////////
	timeStart := firstPrintFlow.CreatedAt.Unix()
	timePart1End := timeStart + crossPart

	timePart2End := timePart1End + crossPart
	timePart3End := timePart2End + crossPart
	timePart4End := timePart3End + crossPart
	timePart5End := timePart4End + crossPart

	timePart1PrintFlows := []*model.FprintFlow{}
	timePart2PrintFlows := []*model.FprintFlow{}
	timePart3PrintFlows := []*model.FprintFlow{}
	timePart4PrintFlows := []*model.FprintFlow{}
	timePart5PrintFlows := []*model.FprintFlow{}

	for _, fpFlow := range fprintFlows {
		fpCreatedTime := fpFlow.CreatedAt.Unix()
		if fpCreatedTime > timeStart && fpCreatedTime < timePart1End {
			timePart1PrintFlows = append(timePart1PrintFlows, fpFlow)
		} else if fpCreatedTime > timePart1End && fpCreatedTime < timePart2End {
			timePart2PrintFlows = append(timePart2PrintFlows, fpFlow)
		} else if fpCreatedTime > timePart2End && fpCreatedTime < timePart3End {
			timePart3PrintFlows = append(timePart3PrintFlows, fpFlow)
		} else if fpCreatedTime > timePart3End && fpCreatedTime < timePart4End {
			timePart4PrintFlows = append(timePart4PrintFlows, fpFlow)
		} else if fpCreatedTime > timePart4End && fpCreatedTime < timePart5End {
			timePart5PrintFlows = append(timePart5PrintFlows, fpFlow)
		}
	}

	//1
	var timePart1TotalBys uint64
	timePart1BysRate := map[string]float64{}
	for _, fprintFlow := range timePart1PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		timePart1TotalBys += flowByte
	}
	for _, fprintFlow := range timePart1PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes

		pbRate := float64(flowByte) / float64(timePart1TotalBys)
		pbRate = util.Decimal(pbRate)

		protocolStr := protobuf.GetFlowProtocols(int(fprintFlow.Protocol))
		if v, ok := timePart1BysRate[protocolStr]; ok {
			timePart1BysRate[protocolStr] = pbRate + v
		} else {
			timePart1BysRate[protocolStr] = pbRate
		}

	}

	//2
	var timePart2TotalBys uint64
	timePart2BysRate := map[string]float64{}

	for _, fprintFlow := range timePart2PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		timePart2TotalBys += flowByte

	}
	for _, fprintFlow := range timePart2PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes

		pbRate := float64(flowByte) / float64(timePart1TotalBys)
		pbRate = util.Decimal(pbRate)

		protocolStr := protobuf.GetFlowProtocols(int(fprintFlow.Protocol))
		if v, ok := timePart2BysRate[protocolStr]; ok {
			timePart2BysRate[protocolStr] = pbRate + v
		} else {
			timePart2BysRate[protocolStr] = pbRate
		}

	}

	//3
	var timePart3TotalBys uint64
	for _, fprintFlow := range timePart3PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		timePart3TotalBys += flowByte

	}
	timePart3BysRate := map[string]float64{}
	for _, fprintFlow := range timePart3PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes

		pbRate := float64(flowByte) / float64(timePart1TotalBys)
		pbRate = util.Decimal(pbRate)

		protocolStr := protobuf.GetFlowProtocols(int(fprintFlow.Protocol))
		if v, ok := timePart3BysRate[protocolStr]; ok {
			timePart3BysRate[protocolStr] = pbRate + v
		} else {
			timePart3BysRate[protocolStr] = pbRate
		}

	}
	//4
	var timePart4TotalBys uint64
	for _, fprintFlow := range timePart4PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		timePart4TotalBys += flowByte

	}
	timePart4BysRate := map[string]float64{}
	for _, fprintFlow := range timePart4PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes

		pbRate := float64(flowByte) / float64(timePart1TotalBys)
		pbRate = util.Decimal(pbRate)

		protocolStr := protobuf.GetFlowProtocols(int(fprintFlow.Protocol))
		if v, ok := timePart4BysRate[protocolStr]; ok {
			timePart4BysRate[protocolStr] = pbRate + v
		} else {
			timePart4BysRate[protocolStr] = pbRate
		}

	}
	//5
	var timePart5TotalBys uint64
	for _, fprintFlow := range timePart5PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		timePart5TotalBys += flowByte

	}

	timePart5BysRate := map[string]float64{}
	for _, fprintFlow := range timePart5PrintFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes

		pbRate := float64(flowByte) / float64(timePart1TotalBys)
		pbRate = util.Decimal(pbRate)

		protocolStr := protobuf.GetFlowProtocols(int(fprintFlow.Protocol))
		if v, ok := timePart5BysRate[protocolStr]; ok {
			timePart5BysRate[protocolStr] = pbRate + v
		} else {
			timePart5BysRate[protocolStr] = pbRate
		}

	}

	timeParts := []int64{
		timeStart, timeStart, timePart2End, timePart3End, timePart4End,
	}

	timePartRate := []map[string]float64{
		timePart1BysRate, timePart2BysRate, timePart3BysRate, timePart4BysRate, timePart5BysRate,
	}

	fmt.Println(protoMap)
	fmt.Println(timePart1End, timePart2End, timePart3End, timePart4End, timePart5End)
	fmt.Println(timePart1BysRate, timePart2BysRate, timePart3BysRate, timePart4BysRate, timePart5BysRate)

	responseData := map[string]interface{}{
		"protos":      protoMap,
		"time_parts":  timeParts,
		"proto_ratio": timePartRate,
	}

	ret := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
	c.JSON(http.StatusOK, ret)
}

//https://gallery.echartsjs.com/editor.html?c=x5_1MX1MWt
func GetAssetProtocolRatio(c *gin.Context) {
	const REMAIN_MIN = 5
	assetId := c.Param("asset_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(assetId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), assetId)
		logger.Logger.Print("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), assetId)
		return
	}

	assetFlows := []*model.Flow{}
	err := mysql.QueryModelRecordsByWhereCondition(&assetFlows, "asset_id = ?", []interface{}{assetId}...)
	if err != nil {
		logger.Logger.Error("%s asset_id:%s err:%s", util.RunFuncName(), assetId, err)
		logger.Logger.Print("%s asset_id:%s err:%s", util.RunFuncName(), assetId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//protocol->bytes
	fprotosBytesMap := map[string]uint64{}

	for _, fpFlow := range assetFlows {
		protocolStr := protobuf.GetFlowProtocols(int(fpFlow.Protocol))
		srcDstBytes := fpFlow.SrcDstBytes //up
		dstSrcBytes := fpFlow.DstSrcBytes //down
		flowByte := dstSrcBytes + srcDstBytes

		if v, ok := fprotosBytesMap[protocolStr]; ok {
			fprotosBytesMap[protocolStr] = v + flowByte
		} else {
			fprotosBytesMap[protocolStr] = flowByte
		}
	}

	//总流量大小
	var totalBytes uint64
	for _, fprintFlow := range assetFlows {
		dstSrcBytes := fprintFlow.DstSrcBytes
		srcDstBytes := fprintFlow.SrcDstBytes
		flowByte := dstSrcBytes + srcDstBytes
		totalBytes += flowByte
	}
	fprotosMap := map[string]float64{}

	for p, pb := range fprotosBytesMap {
		pbRate := float64(pb) / float64(totalBytes)
		pbRate = util.Decimal(pbRate)
		if v, ok := fprotosMap[p]; ok {
			fprotosMap[p] = pbRate + v
		} else {
			fprotosMap[p] = pbRate
		}
	}

	fprotoBytesFloat := []map[string]interface{}{}

	var protoByteFListData ProtoByteFList
	for protoId, protoByteF := range fprotosMap {
		obj := ProtoByteF{Key: protoId, Value: protoByteF}
		protoByteFListData = append(protoByteFListData, obj)
	}

	sort.Sort(protoByteFListData)

	for _, v := range protoByteFListData {
		key := v.Key
		value := v.Value

		protoMap := map[string]interface{}{}
		protoMap["name"] = key
		protoMap["value"] = value

		fprotoBytesFloat = append(fprotoBytesFloat, protoMap)
	}

	responseData := map[string]interface{}{
		"asset_flows": fprotoBytesFloat,
	}

	ret := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFlowSuccessMsg, responseData)
	c.JSON(http.StatusOK, ret)
}

type ProtoByteFList []ProtoByteF
type ProtoByteF struct {
	Key   string
	Value float64
}

func (list ProtoByteFList) Len() int {
	return len(list)
}
func (list ProtoByteFList) Less(i, j int) bool {
	return list[i].Value > list[j].Value
}
func (list ProtoByteFList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
