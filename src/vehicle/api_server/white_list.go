package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)


/*
获取一条白名单
 */
func GetWhiteList(c *gin.Context)  {
	whiteListId:=c.Param("white_list_id")
	argsTrimsEmpty:=util.RrgsTrimsEmpty(whiteListId)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty white_list_id:%s",util.RunFuncName(),argsTrimsEmpty)
		logger.Logger.Print("%s argsTrimsEmpty white_list_id:%s",util.RunFuncName(),argsTrimsEmpty)
	}
	whiteListObj:= &model.WhiteList{}

	modelBase := model_base.ModelBaseImpl(whiteListObj)

	err,recordNotFound:=modelBase.GetModelByCondition("white_list_id = ?",[]interface{}{whiteListId}...)

	if err!=nil{
		logger.Logger.Error("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		logger.Logger.Print("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if recordNotFound{
		logger.Logger.Error("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
		logger.Logger.Print("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListUnExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetWhiteListSuccessMsg,whiteListObj)
	c.JSON(http.StatusOK,retObj)
}



/*
获取所有白名单
 */
func GetWhiteLists(c *gin.Context)  {

	whiteListObj:= []*model.WhiteList{}
	err := model_base.ModelBaseImpl(&model.WhiteList{}).
		GetModelListByCondition(&whiteListObj,"",[]interface{}{}...)
	if err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqGetWhiteListSuccessMsg,whiteListObj)
	c.JSON(http.StatusOK,retObj)
}

func AddWhiteList(c *gin.Context)  {
	destIp := c.PostForm("dest_ip")
	url := c.PostForm("url")
	sourceMac := c.PostForm("source_mac")
	sourceIp := c.PostForm("source_ip")

	argsTrimsEmpty:=util.RrgsTrimsEmpty(destIp,url,sourceMac,sourceIp)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty destIp:%s,url:%s,sourceMac:%s,sourceIp:%s",
			util.RunFuncName(),destIp,url,sourceMac,sourceIp)
		logger.Logger.Print("%s argsTrimsEmpty destIp:%s,url:%s,sourceMac:%s,sourceIp:%s",
			util.RunFuncName(),destIp,url,sourceMac,sourceIp)
		return
	}

	whiteList:= &model.WhiteList{
		WhiteListId:util.RandomString(32),
		DestIp:destIp,
		Url:url,
		SourceMac:sourceMac,
		SourceIp:sourceIp,
	}
	modelBase:=model_base.ModelBaseImpl(whiteList)

	if err:=modelBase.InsertModel();err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqAddWhiteListFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}
	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqAddWhiteListSuccessMsg,whiteList)
	c.JSON(http.StatusOK,retObj)
}

/*
编辑白名单
 */
func EditWhiteList(c *gin.Context)  {
	whiteListId:=c.Param("white_list_id")
	destIp:=c.PostForm("dest_ip")
	url:=c.PostForm("url")
	sourceMac:=c.PostForm("source_mac")
	sourceIp:=c.PostForm("source_ip")

	argsTrimsEmpty:=util.RrgsTrimsEmpty(whiteListId)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s white_list_id argsTrimsEmpty",util.RunFuncName())
		logger.Logger.Print("%s white_list_id argsTrimsEmpty",util.RunFuncName())
	}
	//查询是否存在
	whiteListObj:= &model.WhiteList{}
	modelBase := model_base.ModelBaseImpl(whiteListObj)

	err,recordNotFound:= modelBase.GetModelByCondition("white_list_id = ?",[]interface{}{whiteListId}...)

	if err!=nil{
		logger.Logger.Error("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		logger.Logger.Print("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if recordNotFound{
		logger.Logger.Error("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
		logger.Logger.Print("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListUnExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	//更新
	gormMapAttrs :=  map[string]interface{}{
		"dest_ip":destIp,
		"url":url,
		"source_mac":sourceMac,
		"source_ip":sourceIp,
	}

	err=model_base.ModelBaseImpl(&model.WhiteList{}).
		UpdateModelsByCondition(gormMapAttrs,"white_list_id = ?",[]interface{}{whiteListId}...)

	if err!=nil{
		logger.Logger.Error("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		logger.Logger.Print("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqUpdateWhiteListFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqUpdateWhiteListSuccessMsg,"")
	c.JSON(http.StatusOK,retObj)
}
/*
删除白名单
 */
func DeleWhiteList(c *gin.Context)  {

	whiteListId:=c.Param("white_list_id")
	argsTrimsEmpty:=util.RrgsTrimsEmpty(whiteListId)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s white_list_id argsTrimsEmpty",util.RunFuncName())
		logger.Logger.Print("%s white_list_id argsTrimsEmpty",util.RunFuncName())
	}
	//查询是否存在
	whiteListObj:= &model.WhiteList{}
	modelBase := model_base.ModelBaseImpl(whiteListObj)
	err,recordNotFound:=modelBase.GetModelByCondition("white_list_id = ?",[]interface{}{whiteListId}...)

	if err!=nil{
		logger.Logger.Error("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		logger.Logger.Print("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if recordNotFound{
		logger.Logger.Error("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
		logger.Logger.Print("%s white_list_id:%s,recordNotFound",util.RunFuncName(),whiteListId)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqGetWhiteListUnExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	//删除
	err=model_base.ModelBaseImpl(&model.WhiteList{}).
		DeleModelsByCondition("white_list_id = ?",[]interface{}{whiteListId}...)

	if err!=nil{
		logger.Logger.Error("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		logger.Logger.Print("%s white_list_id:%s,err:%s",util.RunFuncName(),whiteListId,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqDeleWhiteListFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqDeleWhiteListSuccessMsg,"")
	c.JSON(http.StatusOK,retObj)
}