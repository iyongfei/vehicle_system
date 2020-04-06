package api_server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/service"
	"vehicle_system/src/vehicle/util"
)

func Auth(c *gin.Context)  {

	userName := c.PostForm("user_name")
	password := c.PostForm("password")

	logger.Logger.Print("%s userName:%s,password:%s",util.RunFuncName(),userName,password)
	logger.Logger.Info("%s userName:%s,password:%s",util.RunFuncName(),userName,password)

	argsTrimsEmpty:=util.RrgsTrimsEmpty(userName,password)
	if argsTrimsEmpty{
		ret:=response.StructResponseObj(response.VStatusBadRequest,response.ReqArgsIllegalMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s argsTrimsEmpty userName:%s,password:%s",util.RunFuncName(),userName,password)
		logger.Logger.Print("%s argsTrimsEmpty userName:%s,password:%s",util.RunFuncName(),userName,password)
		return
	}

	user:= &model.User{
		UserName:userName,
		Password:password,
	}

	modelBase:=model_base.ModelBaseImpl(user)


	err,recordNotFound := modelBase.GetModelByCondition(
		"user_name = ? and password = ?",[]interface{}{user.UserName,user.Password}...)
	if recordNotFound{
		logger.Logger.Error("%s userName:%s recordNotFound,err:%s",util.RunFuncName(),user.UserName,err)
		logger.Logger.Print("%s userName:%s recordNotFound,err:%s",util.RunFuncName(),user.UserName,err)
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqRegistUnAuthMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	vehicleClaims := service.VehicleClaims{
		UserId:user.UserId,
		UserName:user.UserName,
		PassWord:user.Password,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: service.ExpiresAt,
			Issuer:    conf.SignKey,
		},
	}
	jwtToken,err := service.Jwt.CreateToken(vehicleClaims)

	if err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqRegistAuthFailMsg,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s username:%s createToken err:%v",util.RunFuncName(),userName,err)
		logger.Logger.Print("%s username:%s createToken err:%v",util.RunFuncName(),userName,err)
		return
	}

	jwtTokenObj:= struct {
		Token string `json:"token"`
		UserId string `json:"user_id"`
	}{jwtToken,user.UserId}


	ret:=response.StructResponseObj(response.VStatusOK,response.ReqRegistAuthSuccessMsg,jwtTokenObj)
	c.JSON(http.StatusOK,ret)

	logger.Logger.Error("%s username:%s,createToken success",util.RunFuncName(),user.UserName)
	logger.Logger.Print("%s username:%s,createToken success",util.RunFuncName(),user.UserName)
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

	_,recordNotFound := modelBase.GetModelByCondition("user_name = ?",user.UserName)
	if !recordNotFound{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqRegistExistMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}

	if err:=modelBase.InsertModel();err!=nil{
		ret:=response.StructResponseObj(response.VStatusServerError,response.ReqRegistFailMsg,"")
		c.JSON(http.StatusOK,ret)
		return
	}
	retObj:=response.StructResponseObj(response.VStatusOK,response.ReqRegistSuccessMsg,user)
	c.JSON(http.StatusOK,retObj)

	//retMap:=response.StructResponseMap(response.VStatusOK,response.ReqRegistSuccessMsg,user)
	//c.JSON(http.StatusOK,retMap)
}