package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vehicle_system/src/vehicle/api_server"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/middleware/cors"
	"vehicle_system/src/vehicle/router/v1"
)

func RouterHandler()  {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/auth", api_server.Auth)
	router.POST("/regist", api_server.Regist)

	v1.V1Router(router)

	router.Run(fmt.Sprintf("%s:%d", conf.ServerHost, conf.ServerPort))
}



