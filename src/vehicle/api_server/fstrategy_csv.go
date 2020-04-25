package api_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vehicle_system/src/vehicle/csv"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func GetFStrategyCsv(c *gin.Context) {
	fstrategyId := c.Param("fstrategy_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(fstrategyId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty fstrategyId:%s,argsTrimsEmpty", util.RunFuncName(), fstrategyId)
		logger.Logger.Print("%s argsTrimsEmpty fstrategyId:%s,argsTrimsEmpty", util.RunFuncName(), fstrategyId)
		return
	}

	fstrategy := &model.Fstrategy{
		FstrategyId: fstrategyId,
	}
	fstrategyModelBase := model_base.ModelBaseImpl(fstrategy)

	err, recordNotFound := fstrategyModelBase.GetModelByCondition("fstrategy_id = ?", fstrategy.FstrategyId)
	if err != nil {
		logger.Logger.Error("%s fstrategyId:%s,err:%+v",
			util.RunFuncName(), fstrategyId, err)

		logger.Logger.Print("%s fstrategyId:%s,err:%+v",
			util.RunFuncName(), fstrategyId, err)

		ret := response.StructResponseObj(response.VStatusServerError, response.ReqFstrategyCsvFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if recordNotFound {
		logger.Logger.Error("%s fstrategyId:%s,recordNotFound", util.RunFuncName(), fstrategyId)
		logger.Logger.Print("%s fstrategyId:%s,recordNotFound", util.RunFuncName(), fstrategyId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFstrategyCsvUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//获取csv文件
	csvPath := fstrategy.ScvPath
	fStrategyCsvFolderIndex := strings.Index(csvPath, csv.FStrategyCsvFolder)

	var csvFileName string
	if fStrategyCsvFolderIndex != -1 {
		csvFileName = csvPath[fStrategyCsvFolderIndex:]
	}

	fmt.Println(csvFileName, "csvFileName")
	c.File(csvFileName)
}

/**
上传scv
*/
func UploadFStrategyCsv(c *gin.Context) {
	uploadCsv, err := c.FormFile("upload_csv")
	fmt.Println(uploadCsv, "safly.........")
	if err != nil {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s formfile err:%+v", util.RunFuncName(), err)
		logger.Logger.Print("%s formfile err:%+v", util.RunFuncName(), err)
		return
	}
	dst := uploadCsv.Filename
	if err := c.SaveUploadedFile(uploadCsv, dst); err != nil {
		// ignore
	}
}
