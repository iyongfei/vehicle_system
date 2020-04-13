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


//
//

func AddFlow(c *gin.Context)  {
	//flowId:=c.PostForm("flow_id")
	//vehicleId:=c.PostForm("vehicle_id")

	//argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,flowId)
	//if argsTrimsEmpty{
	//	ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
	//	c.JSON(http.StatusOK,ret)
	//	logger.Logger.Error("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty",util.RunFuncName(),flowId,vehicleId)
	//	logger.Logger.Print("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty",util.RunFuncName(),flowId,vehicleId)
	//}
	//whiteListObj:= &model.Flow{}
	//
	//modelBase := model_base.ModelBaseImpl(whiteListObj)
	//
	//err,recordNotFound:=modelBase.GetModelByCondition("white_list_id = ?",[]interface{}{whiteListId}...)
	//
	//if err!=nil{
	//	logger.Logger.Error("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
	//	logger.Logger.Print("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
	//	ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListFailMsg,"")
	//	c.JSON(http.StatusOK,ret)
	//	return
	//}
	//
	//if recordNotFound{
	//	logger.Logger.Error("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
	//	logger.Logger.Print("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
	//	ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListUnExistMsg,"")
	//	c.JSON(http.StatusOK,ret)
	//	return
	//}
	//
	//retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetWhiteListSuccessMsg,whiteListObj)
	//c.JSON(http.StatusOK,retObj)
}

//
//func EditFlow(c *gin.Context)  {
//	flowId:=c.Param("flow_id")
//	vehicleId:=c.Param("vehicle_id")
//
//	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,flowId)
//	if argsTrimsEmpty{
//		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
//		c.JSON(http.StatusOK,ret)
//		logger.Logger.Error("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty",util.RunFuncName(),flowId,vehicleId)
//		logger.Logger.Print("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty",util.RunFuncName(),flowId,vehicleId)
//	}
//	whiteListObj:= &model.WhiteList{}
//
//	modelBase := model_base.ModelBaseImpl(whiteListObj)
//
//	err,recordNotFound:=modelBase.GetModelByCondition("white_list_id = ?",[]interface{}{whiteListId}...)
//
//	if err!=nil{
//		logger.Logger.Error("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
//		logger.Logger.Print("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
//		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListFailMsg,"")
//		c.JSON(http.StatusOK,ret)
//		return
//	}
//
//	if recordNotFound{
//		logger.Logger.Error("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
//		logger.Logger.Print("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
//		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListUnExistMsg,"")
//		c.JSON(http.StatusOK,ret)
//		return
//	}
//
//	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetWhiteListSuccessMsg,whiteListObj)
//	c.JSON(http.StatusOK,retObj)
//}
//
//
//func DeleFlow(c *gin.Context)  {
//	flowId:=c.Param("flow_id")
//	vehicleId:=c.Param("vehicle_id")
//
//	argsTrimsEmpty:=util.RrgsTrimsEmpty(vehicleId,flowId)
//	if argsTrimsEmpty{
//		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
//		c.JSON(http.StatusOK,ret)
//		logger.Logger.Error("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty",util.RunFuncName(),flowId,vehicleId)
//		logger.Logger.Print("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty",util.RunFuncName(),flowId,vehicleId)
//	}
//	whiteListObj:= &model.WhiteList{}
//
//	modelBase := model_base.ModelBaseImpl(whiteListObj)
//
//	err,recordNotFound:=modelBase.GetModelByCondition("white_list_id = ?",[]interface{}{whiteListId}...)
//
//	if err!=nil{
//		logger.Logger.Error("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
//		logger.Logger.Print("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
//		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListFailMsg,"")
//		c.JSON(http.StatusOK,ret)
//		return
//	}
//
//	if recordNotFound{
//		logger.Logger.Error("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
//		logger.Logger.Print("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
//		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListUnExistMsg,"")
//		c.JSON(http.StatusOK,ret)
//		return
//	}
//
//	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetWhiteListSuccessMsg,whiteListObj)
//	c.JSON(http.StatusOK,retObj)
//}