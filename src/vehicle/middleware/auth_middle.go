package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/service"
)

const (
	vtoken = "vtoken"
	vclaims = "vclaims"
)
func AuthMiddle() gin.HandlerFunc {
	return authMiddleHandlerFunc
}

func authMiddleHandlerFunc(c *gin.Context)  {
	var token string
	token = c.Request.Header.Get(vtoken)

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
	c.Set(vclaims,claims)
	c.Next()
}