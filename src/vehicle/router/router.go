package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"vehicle_system/src/vehicle/api_server"
	"vehicle_system/src/vehicle/conf"
	_ "vehicle_system/src/vehicle/docs"
	"vehicle_system/src/vehicle/middleware/cors"
	"vehicle_system/src/vehicle/router/v1"
)

func RouterHandler() {
	router := gin.Default()
	router.StaticFS("fstrategy_csv", http.Dir("fstrategy_csv"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.Default())
	router.POST("/t_flow", api_server.TFlow)
	router.POST("/auth", api_server.Auth)     //todo
	router.POST("/regist", api_server.Regist) //todo

	v1.V1Router(router)
	router.Run(fmt.Sprintf("%s:%d", conf.ServerHost, conf.ServerPort))
}
