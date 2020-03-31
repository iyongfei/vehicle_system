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
	}
}