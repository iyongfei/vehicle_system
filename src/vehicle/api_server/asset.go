package api_server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"vehicle_system/src/vehicle/emq/emq_cmd"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/emq/topic_publish_handler"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func EditAsset(c *gin.Context) {
	assetId := c.Param("asset_id")
	setTypeP := c.PostForm("type")
	setSwitchP := c.PostForm("switch")

	argsTrimsEmpty := util.RrgsTrimsEmpty(assetId, setTypeP, setSwitchP)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
	}

	setType, _ := strconv.Atoi(setTypeP)
	setSwitch := true
	switch setSwitchP {
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

	//更新
	assetCmd := &emq_cmd.AssetSetCmd{
		VehicleId: assetInfo.VehicleId,
		TaskType:  int(protobuf.Command_DEVICE_SET),

		Switch: setSwitch,
		Type:   setType,
		Mac:    assetId,
	}

	topic_publish_handler.GetPublishService().PutMsg2PublicChan(assetCmd)

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateAssetSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)

}

func GetAssets(c *gin.Context) {
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")

	argsTrimsEmpty := util.RrgsTrimsEmpty(pageSizeP, pageIndexP)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty pageSizeP:%s,pageIndexP:%s", util.RunFuncName(), pageSizeP, pageIndexP)
		logger.Logger.Print("%s argsTrimsEmpty pageSizeP:%s,pageIndexP:%s", util.RunFuncName(), pageSizeP, pageIndexP)
	}

	pageSize, _ := strconv.Atoi(pageSizeP)
	pageIndex, _ := strconv.Atoi(pageIndexP)

	assetInfos := []*model.Asset{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.Asset{})

	err := modelBase.GetModelPaginationByCondition(pageIndex, pageSize,
		&total, &assetInfos, "", "",
		[]interface{}{}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"vehicles":   assetInfos,
		"totalCount": total,
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
	responseData := map[string]interface{}{
		"asset": assetInfo,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetVehicleSuccessMsg, responseData)
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
