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



func GetVehicles(c *gin.Context)  {
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")

	argsTrimsEmpty:=util.RrgsTrimsEmpty(pageSizeP,pageIndexP)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty pageSizeP:%s,pageIndexP:%s",util.RunFuncName(),pageSizeP,pageIndexP)
		logger.Logger.Print("%s argsTrimsEmpty pageSizeP:%s,pageIndexP:%s",util.RunFuncName(),pageSizeP,pageIndexP)
	}

	pageSize, _ := strconv.Atoi(pageSizeP)
	pageIndex, _ := strconv.Atoi(pageIndexP)


	vehicleInfos:= []*model.VehicleInfo{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.VehicleInfo{})

	err := modelBase.GetModelPaginationByCondition(pageIndex,pageSize,
		&total,&vehicleInfos, "",
		[]interface{}{}...)

	if err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetVehiclesFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}


	responseData:= map[string]interface{}{
		"vehicles":vehicleInfos,
		"totalCount":total,
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetVehiclesSuccessMsg,responseData)
	c.JSON(http.StatusOK,retObj)
}



func GetVehicle(c *gin.Context)  {
	vehicleId:=c.Param("vehicle_id")
	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicle_id:%s",util.RunFuncName(),vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicle_id:%s",util.RunFuncName(),vehicleId)
	}
	vehicleInfo:= &model.VehicleInfo{
		VehicleId:vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err,recordNotFound:=modelBase.GetModelByCondition("vehicle_id = ?",[]interface{}{vehicleInfo.VehicleId}...)

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
	responseData:= map[string]interface{}{
		"vehicle":vehicleInfo,
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetVehicleSuccessMsg,responseData)
	c.JSON(http.StatusOK,retObj)
}

/**
添加

type VehicleInfo struct {
	gorm.Model
	VehicleId       string `gorm:"unique"` //小v ID
	Name            string              //小v名称
	Version         string
	StartTime       model_base.UnixTime //启动时间
	FirmwareVersion string
	HardwareModel   string
	Module          string
	SupplyId        string
	UpRouterIp      string

	Ip        string
	Type      uint8
	Mac       string //Mac地址
	TimeStamp uint32 //最近活跃时间戳
	HbTimeout uint32 //最近活跃时间戳

	DeployMode uint8 //部署模式
	FlowIdleTimeSlot uint32

	OnlineStatus  bool   //在线状态
	ProtectStatus uint8  //保护状态										//保护状态
	LeaderId      string //保护状态 // 保护状态
	GroupId string
}

 */

func AddVehicle(c *gin.Context)  {
	body,_:= ioutil.ReadAll(c.Request.Body)

	vehicleInfo := &model.VehicleInfo{

	}
	err:=json.Unmarshal(body,vehicleInfo)


	if err!=nil{
		logger.Logger.Error("%s unmarshal vehicle err:%s",util.RunFuncName(),err.Error())
		logger.Logger.Print("%s unmarshal vehicle err:%s",util.RunFuncName(),err.Error())
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err,recordNotFound:=modelBase.GetModelByCondition("vehicle_id = ?",[]interface{}{vehicleInfo.VehicleId}...)

	if !recordNotFound{
		logger.Logger.Error("%s vehicleId:%s exist",util.RunFuncName(),vehicleInfo.VehicleId)
		logger.Logger.Print("%s vehicleId:%s exist",util.RunFuncName(),vehicleInfo.VehicleId)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetVehicleExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if err:=modelBase.InsertModel();err!=nil{
		logger.Logger.Error("%s add vehicleId:%s err:%s",util.RunFuncName(),vehicleInfo.VehicleId,err.Error())
		logger.Logger.Print("%s add vehicleId:%s err:%s",util.RunFuncName(),vehicleInfo.VehicleId,err.Error())
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqAddVehicleFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	responseData:= map[string]interface{}{
		"vehicle":vehicleInfo,
	}

	ret:=response.StructResponseObj(response.VStatusOK,response.ReqAddVehicleSuccessMsg,responseData)
	c.JSON(http.StatusOK,ret)
}
























