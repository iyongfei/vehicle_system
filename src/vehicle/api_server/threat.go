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
获取一条事件
 */
func GetThreat(c *gin.Context)  {
	threatId:=c.Param("threat_id")
	argsTrimsEmpty:=util.RrgsTrimsEmpty(threatId)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty threatId:%s",util.RunFuncName(),argsTrimsEmpty)
		logger.Logger.Print("%s argsTrimsEmpty threatId:%s",util.RunFuncName(),argsTrimsEmpty)
	}
	threat:= &model.Threat{}

	modelBase := model_base.ModelBaseImpl(threat)

	err,recordNotFound:=modelBase.GetModelByCondition("threat_id = ?",[]interface{}{threatId}...)

	if err!=nil{
		logger.Logger.Error("%s threat_id:%s,err:%s",util.RunFuncName(),threatId,err)
		logger.Logger.Print("%s threat_id:%s,err:%s",util.RunFuncName(),threatId,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetThreatInfoFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if recordNotFound{
		logger.Logger.Error("%s threat_id:%s,recordNotFound",util.RunFuncName(),threatId)
		logger.Logger.Print("%s threat_id:%s,recordNotFound",util.RunFuncName(),threatId)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetThreatInfoUnExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetThreatInfoSuccessMsg,threat)
	c.JSON(http.StatusOK,retObj)
}


/*
获取所有事件
 */
func GetThreats(c *gin.Context)  {
	//whiteListObj:= []*model.WhiteList{}
	threats:= []*model.Threat{}
	err := model_base.ModelBaseImpl(&model.Threat{}).
		GetModelListByCondition(&threats,"",[]interface{}{}...)
	if err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetThreatInfoFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetThreatInfoSuccessMsg,threats)
	c.JSON(http.StatusOK,retObj)
}


/*
获取所有事件
GetModelPaginationByCondition
 */
func GetPaginationThreats(c *gin.Context)  {

	threatType, _ := strconv.Atoi(c.Query("type"))
	status, _ := strconv.Atoi(c.Query("status"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageIndex, _ := strconv.Atoi(c.Query("page_index"))

	fmt.Println("GetPaginationThreats",threatType,status,pageSize,pageIndex)

	//argsTrimsEmpty:=util.RrgsTrimsEmpty(threatId)
	//if argsTrimsEmpty{
	//	ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
	//	c.JSON(http.StatusOK,ret)
	//	logger.Logger.Error("%s argsTrimsEmpty threatId:%s",util.RunFuncName(),argsTrimsEmpty)
	//	logger.Logger.Print("%s argsTrimsEmpty threatId:%s",util.RunFuncName(),argsTrimsEmpty)
	//}


	threats:= []*model.Threat{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.Threat{})

	err := modelBase.GetModelPaginationByCondition(pageIndex,pageSize,
		&total,&threats, "type = ? and status = ?",
		[]interface{}{threatType,status}...)

	if err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetThreatInfoFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetThreatInfoSuccessMsg,threats)
	c.JSON(http.StatusOK,retObj)
}
/*
添加事件
 */

func AddThreat(c *gin.Context)  {

}

/*
编辑事件
 */
func EditThreat(c *gin.Context)  {




}
/*
删除事件
 */
func DeleThreat(c *gin.Context)  {




}