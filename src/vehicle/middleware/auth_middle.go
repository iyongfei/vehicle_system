package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/service"
	"vehicle_system/src/vehicle/util"
)

const (
	Vtoken = "token"
	Vclaims = "claims"
	Vuser = "user"
)
func AuthMiddle() gin.HandlerFunc {
	return authMiddleHandlerFunc
}

func authMiddleHandlerFunc(c *gin.Context)  {
	var token string
	token = c.Request.Header.Get(Vtoken)

	if token == "" {
		ret:=response.StructResponseObj(response.VStatusUnauthorized,response.AuthTokenLost,"")
		c.JSON(http.StatusOK,ret)
		c.Abort()
		return
	}

	claims,err := service.Jwt.ParseToken(token)
	if err != nil {
		ret:=response.StructResponseObj(response.VStatusUnauthorized,response.AuthTokenResignin,"")
		c.JSON(http.StatusOK,ret)
		c.Abort()
		return
	}

	user:= &model.User{
		UserId:claims.UserId,
		UserName:claims.UserName,
		Password:claims.PassWord,
	}

	modelBase:=model_base.ModelBaseImpl(user)

	_,recordNotFound := modelBase.GetModelByCondition(
		"user_name = ? and password = ? and user_id = ?",user.UserName,user.Password,user.UserId)
	if recordNotFound{
		ret:=response.StructResponseObj(response.VStatusUnauthorized,response.ValidationErrorUnverifiableStr,"")
		c.JSON(http.StatusOK,ret)
		logger.Logger.Error("%s token:%s,verify user_id:%s err",util.RunFuncName(),token,claims.UserId)
		logger.Logger.Print("%s token:%s,verify user_id:%s err",util.RunFuncName(),token,claims.UserId)
		return
	}

	c.Set(Vclaims,claims)
	c.Set(Vuser,user)
	c.Next()
}