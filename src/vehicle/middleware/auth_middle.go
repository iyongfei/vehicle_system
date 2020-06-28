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
	Vtoken  = "token"
	Vclaims = "claims"
	Vuser   = "user"
)

func AuthMiddle() gin.HandlerFunc {
	return authMiddleHandlerFunc
}

func authMiddleHandlerFunc(c *gin.Context) {
	var token string
	token = c.Request.Header.Get(Vtoken)

	//是否携带token
	if token == "" {
		ret := response.StructResponseObj(response.VStatusUnauthorized, response.AuthTokenLost, "")
		c.JSON(http.StatusOK, ret)
		c.Abort()
		return
	}

	claims, err := service.NewJWT().ParseToken(token)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusUnauthorized, response.AuthTokenResignin, "")
		c.JSON(http.StatusOK, ret)
		c.Abort()
		logger.Logger.Print("token err %+v", err.Error())
		return
	}

	//token是否解析正确
	user := &model.User{
		UserId:   claims.UserId,
		UserName: claims.UserName,
		Password: claims.PassWord,
	}

	modelBase := model_base.ModelBaseImpl(user)

	_, recordNotFound := modelBase.GetModelByCondition(
		"user_name = ? and password = ? and user_id = ?", user.UserName, user.Password, user.UserId)
	if recordNotFound {
		ret := response.StructResponseObj(response.VStatusUnauthorized, response.ValidationErrorUnverifiableStr, "")
		c.JSON(http.StatusOK, ret)
		c.Abort()
		logger.Logger.Error("%s token:%s,verify user_id:%s err", util.RunFuncName(), token, claims.UserId)
		logger.Logger.Print("%s token:%s,verify user_id:%s err", util.RunFuncName(), token, claims.UserId)
		return
	}

	//某个vehicleId是否授权

	////校验vehicle授权
	//var vehicleIdAuths []string
	//_ = mysql.QueryPluckByModelWhere(&model.VehicleAuth{}, "vehicle_id", &vehicleIdAuths,
	//	"", []interface{}{}...)
	//
	//var vehicleIds []string
	//_ = mysql.QueryPluckByModelWhere(&model.VehicleInfo{}, "vehicle_id", &vehicleIds,
	//	"", []interface{}{}...)
	//
	//vehicleIdAllInAuths := false
	//
	//for _, vehicleId := range vehicleIds {
	//
	//	vehicleIdMd5 := util.Md5(vehicleId + response.VehiclePasswordSecret)
	//	exist := util.IsExistInSlice(vehicleIdMd5, vehicleIdAuths)
	//	if exist {
	//		vehicleIdAllInAuths = true
	//	}
	//}
	//
	//if !vehicleIdAllInAuths {
	//	ret := response.StructResponseObj(response.VStatusUnauthorized, response.ValidationVehicleAuthErrorUnverifiableStr, "")
	//	c.JSON(http.StatusOK, ret)
	//	c.Abort()
	//	logger.Logger.Error("%s token:%s,verify user_id:%s err", util.RunFuncName(), token, claims.UserId)
	//	logger.Logger.Print("%s token:%s,verify user_id:%s err", util.RunFuncName(), token, claims.UserId)
	//	return
	//}

	c.Set(Vclaims, claims)
	c.Set(Vuser, user)
	c.Next()
}
