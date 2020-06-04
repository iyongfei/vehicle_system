package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"vehicle_system/src/vehicle/csv"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

/**
添加白名单资产
*/
func AddWhiteAsset(c *gin.Context) {
	assetIds := c.PostForm("asset_ids")
	argsTrimsEmpty := util.RrgsTrimsEmpty(assetIds)

	//先删除所有的白名单
	whiteAssetDeleModelBase := model_base.ModelBaseImpl(&model.WhiteAsset{})
	err := whiteAssetDeleModelBase.DeleModelsByCondition("", []interface{}{}...)
	if err != nil {
		retObj := response.StructResponseObj(response.VStatusOK, response.ReqDeleAssetFailMsg, nil)
		c.JSON(http.StatusOK, retObj)
		return
	}
	//空白名单
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusOK, response.ReqUpdateWhiteListSuccessMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Print("%s assetIds:%s", util.RunFuncName(), assetIds)
		logger.Logger.Error("%s assetIds:%s", util.RunFuncName(), assetIds)
		return
	}
	logger.Logger.Print("%s assetIds:%s", util.RunFuncName(), assetIds)
	logger.Logger.Info("%s assetIds:%s", util.RunFuncName(), assetIds)

	assetIdSlice := strings.Split(assetIds, ",")
	//过滤需要添加的资产mac列表
	var insertAssetIds []string
	for _, asset_id := range assetIdSlice {
		fAssetId := util.RrgsTrim(asset_id)
		if fAssetId != "" {
			insertAssetIds = append(insertAssetIds, fAssetId)
		}
	}

	///增加指纹库
	var insertWhiteAssets []*model.WhiteAsset
	var insertWhiteAssetIds []string
	for _, asset_id := range insertAssetIds {
		fAssetId := util.RrgsTrim(asset_id)
		if fAssetId != "" {
			//插入
			whiteAssetInfo := &model.WhiteAsset{
				DeviceMac: fAssetId,
				AccessNet: true,
			}

			fprintInfoModelBase := model_base.ModelBaseImpl(whiteAssetInfo)

			err, recordNotFound := fprintInfoModelBase.GetModelByCondition("device_mac = ?", []interface{}{whiteAssetInfo.DeviceMac}...)

			if err != nil {
				logger.Logger.Error("%s asset_id:%s err:%s", util.RunFuncName(), whiteAssetInfo.DeviceMac, err)
				logger.Logger.Print("%s asset_id:%s err:%s", util.RunFuncName(), whiteAssetInfo.DeviceMac, err)
				continue
			}
			if recordNotFound {
				if err := fprintInfoModelBase.InsertModel(); err != nil {
					logger.Logger.Error("%s add asset_id:%s err:%s", util.RunFuncName(), whiteAssetInfo.DeviceMac, err.Error())
					logger.Logger.Print("%s add asset_id:%s err:%s", util.RunFuncName(), whiteAssetInfo.DeviceMac, err.Error())
					continue
				} else {
					insertWhiteAssetIds = append(insertWhiteAssetIds, whiteAssetInfo.DeviceMac)
					insertWhiteAssets = append(insertWhiteAssets, whiteAssetInfo)
				}
			}
		}
	}

	//修改资产状态
	assetModel := &model.Asset{
		AccessNet: true,
	}

	assetModelBase := model_base.ModelBaseImpl(assetModel)

	attrs := map[string]interface{}{
		"access_net": assetModel.AccessNet,
	}
	if err := assetModelBase.UpdateModelsByCondition(attrs, "asset_id in (?)", insertWhiteAssetIds); err != nil {

	}

	//获取插入的然后返回

	responseData := map[string]interface{}{
		"white_assets": insertWhiteAssets,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddWhiteListSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

/**
上传资产白名单
*/
func UploadWhiteAsset(c *gin.Context) {
	uploadCsv, err := c.FormFile("upload_csv")

	//文件获取失败
	if err != nil {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s upload fstrategy csv formfile err:%+v", util.RunFuncName(), err)
		logger.Logger.Print("%s upload fstrategy csv formfile err:%+v", util.RunFuncName(), err)
		return
	}
	logger.Logger.Print("%s vehicle_id:%s,uploadCsv:%s,", util.RunFuncName(), uploadCsv.Filename)
	logger.Logger.Info("%s vehicle_id:%s,uploadCsv:%s", util.RunFuncName(), uploadCsv.Filename)

	uploadFileName := uploadCsv.Filename

	logger.Logger.Info("%s fileName:%s, err:%+v", util.RunFuncName(), uploadFileName, err)
	logger.Logger.Print("%s fileName:%s, err:%+v", util.RunFuncName(), uploadFileName, err)
	//创建文件
	tempCsvName := util.RandomString(8)
	tempCsvFileFolderPath, _ := csv.CreateCsvFolder()
	tempCsvPathName := tempCsvFileFolderPath + "/" + tempCsvName

	if err := c.SaveUploadedFile(uploadCsv, tempCsvPathName); err != nil {
	}

	//解析
	csvReaderModel := csv.CreateCsvReader(tempCsvPathName)
	assetIdSlice, _ := csvReaderModel.ParseAddAssetPrintsCsvFile()

	//先删除所有的白名单
	whiteAssetDeleModelBase := model_base.ModelBaseImpl(&model.WhiteAsset{})
	err = whiteAssetDeleModelBase.DeleModelsByCondition("", []interface{}{}...)
	if err != nil {
		retObj := response.StructResponseObj(response.VStatusOK, response.ReqDeleAssetFailMsg, nil)
		c.JSON(http.StatusOK, retObj)
		return
	}

	if len(assetIdSlice) == 0 {
		//删除文件
		if csv.IsExists(tempCsvPathName) {
			os.Remove(tempCsvPathName)
		}

		ret := response.StructResponseObj(response.VStatusOK, response.ReqDeleAssetSuccessMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Print("%s assetIds:%+v", util.RunFuncName(), assetIdSlice)
		logger.Logger.Error("%s assetIds:%+v", util.RunFuncName(), assetIdSlice)
		return
	}

	logger.Logger.Print("%s assetIds:%s", util.RunFuncName(), assetIdSlice)
	logger.Logger.Info("%s assetIds:%s", util.RunFuncName(), assetIdSlice)

	//过滤需要添加的资产mac列表
	var insertAssetIds []string
	for _, asset_id := range assetIdSlice {
		fAssetId := util.RrgsTrim(asset_id)
		if fAssetId != "" {
			insertAssetIds = append(insertAssetIds, fAssetId)
		}
	}

	///增加指纹库
	var insertFprints []*model.WhiteAsset
	var insertFprintIds []string
	for _, asset_id := range insertAssetIds {
		fAssetId := util.RrgsTrim(asset_id)
		if fAssetId != "" {
			//插入
			fprintInfo := &model.WhiteAsset{
				DeviceMac: fAssetId,
				AccessNet: true,
			}

			fprintInfoModelBase := model_base.ModelBaseImpl(fprintInfo)

			err, recordNotFound := fprintInfoModelBase.GetModelByCondition("device_mac = ?", []interface{}{fprintInfo.DeviceMac}...)

			if err != nil {
				logger.Logger.Error("%s asset_id:%s err:%s", util.RunFuncName(), fprintInfo.DeviceMac, err)
				logger.Logger.Print("%s asset_id:%s err:%s", util.RunFuncName(), fprintInfo.DeviceMac, err)
				continue
			}
			if recordNotFound {
				if err := fprintInfoModelBase.InsertModel(); err != nil {
					logger.Logger.Error("%s add asset_id:%s err:%s", util.RunFuncName(), fprintInfo.DeviceMac, err.Error())
					logger.Logger.Print("%s add asset_id:%s err:%s", util.RunFuncName(), fprintInfo.DeviceMac, err.Error())
					continue
				} else {
					insertFprintIds = append(insertFprintIds, fprintInfo.DeviceMac)
					insertFprints = append(insertFprints, fprintInfo)
				}
			}
		}
	}

	//修改资产状态
	assetModel := &model.Asset{
		ProtectStatus: true,
	}

	assetModelBase := model_base.ModelBaseImpl(assetModel)

	attrs := map[string]interface{}{
		"access_net": assetModel.ProtectStatus,
	}
	if err := assetModelBase.UpdateModelsByCondition(attrs, "asset_id in (?)", insertFprintIds); err != nil {

	}

	if csv.IsExists(tempCsvPathName) {
		os.Remove(tempCsvPathName)
	}

	//获取插入的然后返回

	responseData := map[string]interface{}{
		"white_assets": insertFprints,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddAssetSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}

/**
查看所有的白名单
*/
func GetWhiteAssetMacs(c *gin.Context) {
	var deviceMacs []string
	err := mysql.QueryPluckByModelWhere(&model.WhiteAsset{}, "device_mac", &deviceMacs,
		"", []interface{}{}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetAssetFprintsFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	responseData := map[string]interface{}{
		"white_asset_macs": deviceMacs,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetAssetFprintsSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}
