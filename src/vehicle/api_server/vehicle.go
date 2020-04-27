package api_server

import (
	"encoding/json"
	"fmt"
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

func EditVehicle(c *gin.Context) {
	vehicleId := c.Param("vehicle_id")
	setTypeP := c.PostForm("type")
	setSwitchP := c.PostForm("switch")

	fmt.Println("EditVehicle:::::::::", vehicleId, setTypeP, setSwitchP)
	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, setTypeP, setSwitchP)
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
	vehicleInfo := &model.VehicleInfo{}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//更新
	vehicleCmd := &emq_cmd.VehicleSetCmd{
		VehicleId: vehicleId,
		TaskType:  int(protobuf.Command_GW_SET),

		Type:   setType,
		Switch: setSwitch,
	}

	topic_publish_handler.GetPublishService().PutMsg2PublicChan(vehicleCmd)

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateVehicleSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}

func GetVehicles(c *gin.Context) {
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

	vehicleInfos := []*model.VehicleInfo{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.VehicleInfo{})

	err := modelBase.GetModelPaginationByCondition(pageIndex, pageSize,
		&total, &vehicleInfos, "",
		[]interface{}{}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehiclesFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"vehicles":   vehicleInfos,
		"totalCount": total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetVehiclesSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func GetVehicle(c *gin.Context) {
	vehicleId := c.Param("vehicle_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), vehicleId)
	}
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	type VehicleResponse struct {
		VehicleId     string
		Ip            string
		Mac           string
		OnlineStatus  bool
		ProtectStatus uint8
	}

	vehicleResponse := VehicleResponse{
		VehicleId:     vehicleId,
		Ip:            vehicleInfo.Ip,
		Mac:           vehicleInfo.Mac,
		OnlineStatus:  vehicleInfo.OnlineStatus,
		ProtectStatus: vehicleInfo.ProtectStatus,
	}

	responseData := map[string]interface{}{
		"vehicle": vehicleResponse,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetVehicleSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

/**
添加
*/

func AddVehicle(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	vehicleInfo := &model.VehicleInfo{}
	err := json.Unmarshal(body, vehicleInfo)

	if err != nil {
		logger.Logger.Error("%s unmarshal vehicle err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Print("%s unmarshal vehicle err:%s", util.RunFuncName(), err.Error())
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if !recordNotFound {
		logger.Logger.Error("%s vehicleId:%s exist", util.RunFuncName(), vehicleInfo.VehicleId)
		logger.Logger.Print("%s vehicleId:%s exist", util.RunFuncName(), vehicleInfo.VehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if err := modelBase.InsertModel(); err != nil {
		logger.Logger.Error("%s add vehicleId:%s err:%s", util.RunFuncName(), vehicleInfo.VehicleId, err.Error())
		logger.Logger.Print("%s add vehicleId:%s err:%s", util.RunFuncName(), vehicleInfo.VehicleId, err.Error())
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"vehicle": vehicleInfo,
	}

	ret := response.StructResponseObj(response.VStatusOK, response.ReqAddVehicleSuccessMsg, responseData)
	c.JSON(http.StatusOK, ret)
}

func DeleVehicle(c *gin.Context) {
	vehicleId := c.Param("vehicle_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s argsTrimsEmpty", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s argsTrimsEmpty", util.RunFuncName(), vehicleId)
		return
	}

	fmt.Println(vehicleId, "jsldfjs")

	vehicleObj := &model.VehicleInfo{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleObj)
	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleObj.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicleId:%s err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicleId:%s err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if recordNotFound {
		logger.Logger.Error("%s vehicleId:%s,record not exist", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicleId:%s,record not exist", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err := modelBase.DeleModelsByCondition("vehicle_id = ?",
		[]interface{}{vehicleObj.VehicleId}...); err != nil {
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqDeleVehicleSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}
