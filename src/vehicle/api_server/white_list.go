package api_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/middleware"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)


/*
获取一条白名单
 */
func GetWhiteList(c *gin.Context)  {

	//vehicleClaimsUser,_ := c.Get(middleware.Vuser)




	//argsTrimsEmpty:=util.RrgsTrimsEmpty(userId)
	//if argsTrimsEmpty{
	//	ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
	//	c.JSON(http.StatusOK,ret)
	//	logger.Logger.Error("%s argsTrimsEmpty userName:%s,password:%s",util.RunFuncName(),userName,password)
	//	logger.Logger.Print("%s argsTrimsEmpty userName:%s,password:%s",util.RunFuncName(),userName,password)
	//}





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
	vehicleClaimsUser,_ := c.Get(middleware.Vuser)
	fmt.Printf("vehicleClaimsUser%+v\n",vehicleClaimsUser)

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

	if err:=modelBase.InsertModel(whiteList);err!=nil{
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




}
/*
删除白名单
 */
func DeleWhiteList(c *gin.Context)  {




}