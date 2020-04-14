package api_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

/*
获取一条消息流
 */
func GetFlow(c *gin.Context)  {
	flowId:=c.Param("flow_id")
	vehicleId:=c.Query("vehicle_id")

	fmt.Println("getflow.........")
	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,flowId)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty",util.RunFuncName(),flowId,vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty",util.RunFuncName(),flowId,vehicleId)
	}
	flowObj:= &model.Flow{}

	modelBase := model_base.ModelBaseImpl(flowObj)

	err,recordNotFound:=modelBase.GetModelByCondition("vehicle_id = ? and flow_id = ?",[]interface{}{vehicleId,flowId}...)

	if err!=nil{
		logger.Logger.Error("%s flowId:%s,err:%s",util.RunFuncName(),flowId,err)
		logger.Logger.Print("%s flowId:%s,err:%s",util.RunFuncName(),flowId,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetFlowFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if recordNotFound{
		logger.Logger.Error("%s flowId:%s,recordNotFound",util.RunFuncName(),flowId)
		logger.Logger.Print("%s flowId:%s,recordNotFound",util.RunFuncName(),flowId)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetFlowUnExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetFlowSuccessMsg,flowObj)
	c.JSON(http.StatusOK,retObj)
}

func GetFlows(c *gin.Context)  {
	vehicleId:=c.Query("vehicle_id")
	fmt.Println("vehicleId.....",vehicleId)

	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s argsTrimsEmpty",util.RunFuncName(),vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s argsTrimsEmpty",util.RunFuncName(),vehicleId)
	}

	flows:= []*model.Flow{}
	err := model_base.ModelBaseImpl(&model.Flow{}).
		GetModelListByCondition(&flows,"vehicle_id = ?",[]interface{}{vehicleId}...)
	if err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetFlowFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetFlowSuccessMsg,flows)
	c.JSON(http.StatusOK,retObj)
}




/*
获取所有消息会话
GetModelPaginationByCondition
 */
func GetPaginationFlows(c *gin.Context)  {
	vehicleId := c.Query("vehicle_id")
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")

	logger.Logger.Info("%s request params vehicle_id:%s,page_size:%s,page_index:%s",util.RunFuncName(),vehicleId,pageSizeP,pageIndexP)
	logger.Logger.Print("%s request params vehicle_id:%s,page_size:%s,page_index:%s",util.RunFuncName(),vehicleId,pageSizeP,pageIndexP)

	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,pageSizeP,pageIndexP)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty threatId:%s",util.RunFuncName(),argsTrimsEmpty)
		logger.Logger.Print("%s argsTrimsEmpty threatId:%s",util.RunFuncName(),argsTrimsEmpty)
	}

	pageSize, _ := strconv.Atoi(pageSizeP)
	pageIndex, _ := strconv.Atoi(pageIndexP)


	flows:= []*model.Flow{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.Flow{})

	err := modelBase.GetModelPaginationByCondition(pageIndex,pageSize,
		&total,&flows, "vehicle_id = ?",
		[]interface{}{vehicleId}...)

	if err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetFlowFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetFlowSuccessMsg,flows)
	c.JSON(http.StatusOK,retObj)
}

func AddFlow(c *gin.Context)  {
	vehicleId:=c.PostForm("vehicle_id")
	hashP:=c.PostForm("hash")
	srcIpArp:=c.PostForm("src_ip")
	dstIpArg:=c.PostForm("dst_ip")

	logger.Logger.Print("%s vehicleId:%s,hash:%s,srcIpArp:%s,dstIpArg:%s",util.RunFuncName(),vehicleId,hashP,srcIpArp,dstIpArg)


	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,hashP,srcIpArp,dstIpArg)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty",util.RunFuncName(),vehicleId,hashP)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty",util.RunFuncName(),vehicleId,hashP)
	}

	hash,_:= strconv.Atoi(hashP)
	srcIp,_:= strconv.Atoi(srcIpArp)
	dstIp,_:= strconv.Atoi(dstIpArg)

	flowObj:= &model.Flow{
		VehicleId:vehicleId,
		Hash:uint32(hash),
		SrcIp:uint32(srcIp),
		DstIp:uint32(dstIp),
		FlowId:uint32(hash),
	}
	modelBase := model_base.ModelBaseImpl(flowObj)
	err,recordNotFound:=modelBase.GetModelByCondition("vehicle_id = ? and hash = ?",[]interface{}{vehicleId,hash}...)

	if err!=nil{
		logger.Logger.Error("%s vehicleId:%s,hash:%d,get flow info err:%s",util.RunFuncName(),vehicleId,hash,err)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,get flow info err:%s",util.RunFuncName(),vehicleId,hash,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqAddFlowFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}
	if !recordNotFound{
		logger.Logger.Error("%s vehicleId:%s,hash:%d,record exist",util.RunFuncName(),vehicleId,hash)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,record exist",util.RunFuncName(),vehicleId,hash)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetFlowExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if err:=modelBase.InsertModel();err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqAddFlowFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}
	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqAddFlowSuccessMsg,flowObj)
	c.JSON(http.StatusOK,retObj)
}

func EditFlow(c *gin.Context)  {
	hashP:=c.Param("flow_id")
	vehicleId:=c.PostForm("vehicle_id")
	srcIpArp:=c.PostForm("src_ip")
	dstIpArg:=c.PostForm("dst_ip")

	logger.Logger.Print("%s vehicleId:%s,hash:%s,srcIpArp:%s,dstIpArg:%s",util.RunFuncName(),vehicleId,hashP,srcIpArp,dstIpArg)


	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,hashP,srcIpArp,dstIpArg)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty",util.RunFuncName(),vehicleId,hashP)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty",util.RunFuncName(),vehicleId,hashP)
	}

	hash,_:= strconv.Atoi(hashP)
	srcIp,_:= strconv.Atoi(srcIpArp)
	dstIp,_:= strconv.Atoi(dstIpArg)

	flowObj:= &model.Flow{
		VehicleId:vehicleId,
		Hash:uint32(hash),
	}

	modelBase := model_base.ModelBaseImpl(flowObj)
	err,recordNotFound:=modelBase.GetModelByCondition("vehicle_id = ? and hash = ?",[]interface{}{vehicleId,hash}...)

	if err!=nil{
		logger.Logger.Error("%s vehicleId:%s,hash:%d,get flow info err:%s",util.RunFuncName(),vehicleId,hash,err)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,get flow info err:%s",util.RunFuncName(),vehicleId,hash,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqEditFlowFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}
	if recordNotFound{
		logger.Logger.Error("%s vehicleId:%s,hash:%d,record not exist",util.RunFuncName(),vehicleId,hash)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,record not exist",util.RunFuncName(),vehicleId,hash)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetFlowUnExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	//赋值
	flowObj.SrcIp = uint32(srcIp)
	flowObj.DstIp = uint32(dstIp)

	attrs := map[string]interface{}{
		"src_ip": flowObj.SrcIp,
		"dst_ip": flowObj.DstIp,
	}
	if err:=modelBase.UpdateModelsByCondition(attrs,"vehicle_id = ? and hash = ?",
		[]interface{}{flowObj.VehicleId,flowObj.Hash}...);err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqEditFlowFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqEditFlowSuccessMsg,flowObj)
	c.JSON(http.StatusOK,retObj)
}


func DeleFlow(c *gin.Context)  {
	hashP:=c.Param("flow_id")
	vehicleId:=c.Query("vehicle_id")

	fmt.Println(hashP,vehicleId,"param::::::::::")

	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,hashP)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty",util.RunFuncName(),vehicleId,hashP)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty",util.RunFuncName(),vehicleId,hashP)
		return
	}

	hash,_:= strconv.Atoi(hashP)

	flowObj:= &model.Flow{
		VehicleId:vehicleId,
		Hash:uint32(hash),
	}

	modelBase := model_base.ModelBaseImpl(flowObj)
	err,recordNotFound:=modelBase.GetModelByCondition("vehicle_id = ? and hash = ?",[]interface{}{vehicleId,hash}...)

	if err!=nil{
		logger.Logger.Error("%s vehicleId:%s,hash:%d,get flow info err:%s",util.RunFuncName(),vehicleId,hash,err)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,get flow info err:%s",util.RunFuncName(),vehicleId,hash,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqDeleFlowFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}
	if recordNotFound{
		logger.Logger.Error("%s vehicleId:%s,hash:%d,record not exist",util.RunFuncName(),vehicleId,hash)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,record not exist",util.RunFuncName(),vehicleId,hash)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetFlowUnExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}
	if err:=modelBase.DeleModelsByCondition("vehicle_id = ? and hash = ?",
		[]interface{}{flowObj.VehicleId,flowObj.Hash}...);err!=nil{
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetWhiteListSuccessMsg,flowObj)
	c.JSON(http.StatusOK,retObj)
}