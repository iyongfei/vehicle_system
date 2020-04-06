package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vehicle_system/src/vehicle/emq/emq_cmd"
	"vehicle_system/src/vehicle/emq/topic_publish_handler"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
	"vehicle_system/src/vehicle/emq/protobuf"
)

func EditVehicle(c *gin.Context)  {
	vehicleId:=c.Param("vehicle_id")
	setTypeP:=c.PostForm("set_type")
	setSwitchP:=c.PostForm("set_switch")

	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,setTypeP,setSwitchP)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty",util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty",util.RunFuncName())
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
	vehicleInfo:= &model.VehicleInfo{}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err,recordNotFound:= modelBase.GetModelByCondition("vehicle_id = ?",[]interface{}{vehicleId}...)

	if err!=nil{
		logger.Logger.Error("%s vehicle_id:%s,err:%s",util.RunFuncName(),vehicleId,err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s",util.RunFuncName(),vehicleId,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetVehicleFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if recordNotFound{
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound",util.RunFuncName(),vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound",util.RunFuncName(),vehicleId)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetVehicleUnExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	//更新
	vehicleCmd := &emq_cmd.VehicleSetCmd{
		VehicleId:vehicleId,
		Type:setType,
		TaskType:int(protobuf.Command_GW_SET),

		Switch:setSwitch,
		CmdId:int(protobuf.Command_GW_SET),
	}

	topic_publish_handler.GetPublishService().PutMsg2PublicChan(vehicleCmd)

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqUpdateWhiteListSuccessMsg,"")
	c.JSON(http.StatusOK,retObj)
}