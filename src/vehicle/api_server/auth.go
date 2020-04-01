package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func Auth(c *gin.Context)  {

	userName := c.PostForm("user_name")
	password := c.PostForm("password")

	argsTrimsEmpty:=util.RrgsTrimsEmpty(userName,password)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	user:= &model.User{
		UserId:"sdf",
		UserName:userName,
		Password:password,
	}

	re:=model_base.ModelBaseImpl(user)

	re.InsertModel(user)
	//accounts,err := accountHandleModle.GetAllModels()
	//
	//
	//user.InsetUser(user)




	//vhaloClaims := middleware.JWT.VhaloClaims{
	//	UserId:managerModel.UserId,
	//	UserName:managerModel.Account,
	//	PassWord:managerModel.Password,
	//	StandardClaims:jwt.StandardClaims{
	//		ExpiresAt: service.ExpiresAt, // 过期时间 2小时
	//		Issuer:    service.SignKeyStr,              //签名的发行者
	//	},
	//}
	//jwtToken,err := middleware.Jwt.CreateToken(vhaloClaims)



}


func Regist(c *gin.Context)  {

	userName := c.PostForm("user_name")
	password := c.PostForm("password")

	argsTrimsEmpty:=util.RrgsTrimsEmpty(userName,password)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	userId := util.RandomString(32)
	psMd5 := util.Md5(password + response.PasswordSecret)

	user:= &model.User{
		UserId:userId,
		UserName:userName,
		Password:psMd5,
	}

	modelBase:=model_base.ModelBaseImpl(user)

	_,recordNotFound := modelBase.GetModelsByCondition(user,"user_name = ?",user.UserName)
	if !recordNotFound{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqRegistExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if err:=modelBase.InsertModel(user);err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqRegistFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}
	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqRegistSuccessMsg,user)
	c.JSON(http.StatusOK,retObj)

	retMap:=response.StructResponseMap(response.VStatusOK,response.ReqRegistSuccessMsg,user)
	c.JSON(http.StatusOK,retMap)
}