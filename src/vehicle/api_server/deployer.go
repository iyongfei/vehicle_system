package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vehicle_system/src/vehicle/emq/emq_cmd"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/emq/topic_publish_handler"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func EditDeployer(c *gin.Context) {
	deployerId := c.Param("deployer_id")
	vehicleId := c.PostForm("vehicle_id")
	devName := c.PostForm("dev_name")
	name := c.PostForm("name")
	phone := c.PostForm("phone")

	logger.Logger.Info("%s deployer_id:%s,vehicle_id:%s,dev_name:%s,name:%s,phone:%s,", util.RunFuncName(),deployerId,vehicleId,devName,name,phone)
	logger.Logger.Print("%s deployer_id:%s,vehicle_id:%s,dev_name:%s,name:%s,phone:%s,", util.RunFuncName(),deployerId,vehicleId,devName,name,phone)

	//参数都为空
	RrgsTrimsAllEmpty := util.RrgsTrimsAllEmpty(devName, name, phone)
	//只要有一个为空
	RrgsTrimsEmpty := util.RrgsTrimsAllEmpty(deployerId, vehicleId)

	//电话格式
	if RrgsTrimsAllEmpty ||
		RrgsTrimsEmpty ||
		(strings.Trim(phone, " ") != "" && !util.VerifyMobileFormat(phone)){
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmptyrrr", util.RunFuncName())
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		return

	}

	//查询是否存在
	vehicleInfo := &model.VehicleInfo{}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ? and leader_id = ?", []interface{}{vehicleId,deployerId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleBindLeaderFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleBindLeaderUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//更新
	deployerCmd := &emq_cmd.DeployerSetCmd{
		VehicleId:  vehicleId,
		TaskType:   int(protobuf.Command_DEPLOYER_SET),

		Name:       name,
		Phone:      phone,
		DevName:    devName,
	}

	topic_publish_handler.GetPublishService().PutMsg2PublicChan(deployerCmd)

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateVehicleBindLeaderSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}
