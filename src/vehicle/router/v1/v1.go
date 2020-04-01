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
		apiV1.GET("/events/:id", api_server.GetEvent)
		apiV1.GET("/events", api_server.GetEvents)
		apiV1.POST("/events", api_server.AddEvent)
		apiV1.PUT("/events/:id", api_server.EditEvent)
		apiV1.DELETE("/events/:id", api_server.DeleEvent)

		apiV1.GET("/white_lists/:id", api_server.GetWhiteList)
		apiV1.GET("/white_lists", api_server.GetWhiteLists)
		apiV1.POST("/white_lists", api_server.AddWhiteList)
		apiV1.PUT("/white_lists/:id", api_server.EditWhiteList)
		apiV1.DELETE("/white_lists/:id", api_server.DeleWhiteList)
	}
}