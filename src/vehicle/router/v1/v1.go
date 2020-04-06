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


		apiV1.GET("/threats/:threat_id", api_server.GetThreat)
		apiV1.GET("/threats", api_server.GetThreats)
		apiV1.GET("/pagination/threats", api_server.GetPaginationThreats)
		//apiV1.POST("/threats", api_server.AddThreat)
		//apiV1.PUT("/threats/:threat_id", api_server.EditThreat)
		apiV1.DELETE("/threats/:threat_id", api_server.DeleThreat)


		//需改信息，下发给emq
		apiV1.PUT("/vehicle/:vehicle_id", api_server.EditVehicle)
	}
}