package v1

import (
	"github.com/gin-gonic/gin"
	"vehicle_system/src/vehicle/api_server"
	"vehicle_system/src/vehicle/middleware"
)

func V1Router(r *gin.Engine)  {
	apiV1:=r.Group("/api/v1")
	apiV1.Use(middleware.AuthMiddle())
	{
		apiV1.GET("/white_lists/:white_list_id", api_server.GetWhiteList)
		apiV1.GET("/white_lists", api_server.GetWhiteLists)
		apiV1.POST("/white_lists", api_server.AddWhiteList)
		apiV1.PUT("/white_lists/:white_list_id", api_server.EditWhiteList)
		apiV1.DELETE("/white_lists/:white_list_id", api_server.DeleWhiteList)


		apiV1.GET("/threats/:id", api_server.GetEvent)
		apiV1.GET("/threats", api_server.GetEvents)
		apiV1.POST("/threats", api_server.AddEvent)
		apiV1.PUT("/threats/:id", api_server.EditEvent)
		apiV1.DELETE("/threats/:id", api_server.DeleEvent)
	}
}