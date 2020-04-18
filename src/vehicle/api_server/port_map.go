package api_server

import (
	"github.com/gin-gonic/gin"
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

func EditPortMap(c *gin.Context) {
	portMapId := c.Param("port_map_id")
	vehicleId := c.PostForm("vehicle_id")

	srcPort := c.PostForm("src_port")
	destPort := c.PostForm("dest_port")
	destIp := c.PostForm("dest_ip")
	switchFlag := c.PostForm("switch")
	protocol := c.PostForm("protocol")


	logger.Logger.Info("%s portMapId:%s,vehicleId:%s,srcPort:%s,destPort:%s,destIp:%s,switchFlag:%s,protocol:%s",
		util.RunFuncName(),portMapId,vehicleId,srcPort,destPort,destIp,switchFlag,protocol)
	logger.Logger.Print("%s portMapId:%s,vehicleId:%s,srcPort:%s,destPort:%s,destIp:%s,switchFlag:%s,protocol:%s",
		util.RunFuncName(),portMapId,vehicleId,srcPort,destPort,destIp,switchFlag,protocol)

	argsTrimsEmpty := util.RrgsTrimsEmpty(portMapId, vehicleId, srcPort,destPort,destIp,switchFlag)

	//srcPort，destPort端口限制
	srcPortValid := util.VerifyIpPort(srcPort)
	destPortValid := util.VerifyIpPort(destPort)
	//ip格式
	destIpValid := util.IpFormat(destIp)


	switchFlagValid := util.IsEleExistInSlice(switchFlag,[]interface{}{response.FalseFlag,response.TrueFlag})
	protoValid := util.IsEleExistInSlice(protocol,[]interface{}{
		string(protobuf.PortRedirectSetParam_UNSET),
		string(protobuf.PortRedirectSetParam_UDP),
		string(protobuf.PortRedirectSetParam_TCP),
		string(protobuf.PortRedirectSetParam_ALL),})

	if argsTrimsEmpty ||
		!srcPortValid ||
		!destPortValid||
		!destIpValid  ||
		!switchFlagValid||
		!protoValid {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
	}
	//转换switchFlag
	switchFlagParseBool,_ := strconv.ParseBool(switchFlag)
	//转换protocol
	protocolParseInt,_ := strconv.Atoi(protocol)

	//查询是否存在
	portMapInfo := &model.PortMap{
		PortMapId:portMapId,
		VehicleId:vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(portMapInfo)
	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ? and port_map_id = ?",
		[]interface{}{portMapInfo.VehicleId,portMapInfo.PortMapId}...)

	if err != nil {
		logger.Logger.Error("%s port_map_id:%s,vehicle_id:%s,err:%s",
			util.RunFuncName(), portMapInfo.PortMapId,portMapInfo.VehicleId, err)

		logger.Logger.Print("%s port_map_id:%s,vehicle_id:%s,err:%s",
			util.RunFuncName(), portMapInfo.PortMapId,portMapInfo.VehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetPortMapFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s port_map_id:%s,recordNotFound", util.RunFuncName(), portMapInfo.PortMapId)
		logger.Logger.Print("%s port_map_id:%s,recordNotFound", util.RunFuncName(), portMapInfo.PortMapId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetPortMapUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}


	attrs := map[string]interface{}{
		"src_port": srcPort,
		"dst_port": destPort,
		"dst_ip": destIp,
		"switch": switchFlag,
		"protocol_type": protocolParseInt,
	}
	if err:=modelBase.UpdateModelsByCondition(attrs,"vehicle_id = ? and port_map_id = ?",
		[]interface{}{portMapInfo.VehicleId,portMapInfo.PortMapId}...);err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqUpdatePortMapFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	//发布消息
	//更新
	strategyCmd := &emq_cmd.PortMapSetCmd{
		VehicleId:vehicleId ,
		TaskType:  int(protobuf.Command_PORTREDIRECT_SET),

		DestPort:destPort,
		SrcPort:srcPort,
		Switch:switchFlagParseBool,
		Protocol:protocolParseInt,
		DestIp:destIp,
	}
	topic_publish_handler.GetPublishService().PutMsg2PublicChan(strategyCmd)


	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdatePortMapSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}
