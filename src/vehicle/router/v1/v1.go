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




		//flow
		apiV1.GET("/flows/:flow_id", api_server.GetFlow)
		apiV1.GET("/flows", api_server.GetFlows)
		apiV1.GET("/pagination/flows", api_server.GetPaginationFlows)
		apiV1.POST("/flows", api_server.AddFlow)
		apiV1.PUT("/flows/:flow_id", api_server.EditFlow)
		apiV1.DELETE("/flows/:flow_id", api_server.DeleFlow)

		//车载
		apiV1.GET("/vehicles", api_server.GetVehicles)
		apiV1.GET("/vehicles/:vehicle_id", api_server.GetVehicle)
		apiV1.POST("/vehicles", api_server.AddVehicle)
		apiV1.PUT("/vehicles/:vehicle_id", api_server.EditVehicle)
		apiV1.DELETE("/vehicles/:vehicle_id", api_server.DeleVehicle)

		//资产
		apiV1.GET("/assets", api_server.GetAssets)
		apiV1.GET("/assets/:asset_id", api_server.GetAsset)
		apiV1.POST("/assets", api_server.AddAsset)
		apiV1.PUT("/assets/:asset_id", api_server.EditAsset)
		apiV1.DELETE("/assets/:asset_id", api_server.DeleAsset)

		//策略
		apiV1.POST("/strategys", api_server.AddStrategy)
		apiV1.GET("/strategys", api_server.GetStrategys)
		apiV1.GET("/strategys/:strategy_id", api_server.GetStrategy)
		apiV1.DELETE("/strategys/:strategy_id", api_server.DeleStrategy)
		apiV1.PUT("/strategys/:strategy_id", api_server.EditStrategy)

		apiV1.GET("/strategy_vehicles/:strategy_id", api_server.GetStrategyVehicle)
		apiV1.GET("/vehicle_lresults/:vehicle_id", api_server.GetVehicleLearningResults)
		apiV1.GET("/strategy_vehicle_lresults/:strategy_id", api_server.GetStrategyVehicleLearningResults)

	}
}




















