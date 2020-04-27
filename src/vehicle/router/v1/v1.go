package v1

import (
	"github.com/gin-gonic/gin"
	"vehicle_system/src/vehicle/api_server"
)

func V1Router(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	apiV1.Use( /*middleware.AuthMiddle()*/ )
	{
		//////////////////////////////////////////////会话策略接口//////////////////////////////////////////////
		apiV1.POST("/fstrategys", api_server.AddFStrategy)
		apiV1.DELETE("/fstrategys/:fstrategy_id", api_server.DeleFStrategy)
		apiV1.PUT("/fstrategys/:fstrategy_id", api_server.EditFStrategy)
		apiV1.GET("/fstrategys/:fstrategy_id", api_server.GetFStrategy)
		//////////////////////////////////////////////会话策略下载上传//////////////////////////////////////////////
		apiV1.GET("/fstrategy_csvs/:fstrategy_id", api_server.GetFStrategyCsv)
		apiV1.POST("/fstrategy_csvs/", api_server.UploadFStrategyCsv)
		//////////////////////////////////////////////会话接口//////////////////////////////////////////////
		apiV1.GET("/flows/:flow_id", api_server.GetFlow)
		apiV1.GET("/pagination/flows", api_server.GetPaginationFlows)
		apiV1.DELETE("/flows/:flow_id", api_server.DeleFlow)
		//apiV1.POST("/flows", api_server.AddFlow)
		//apiV1.PUT("/flows/:flow_id", api_server.EditFlow)
		//apiV1.GET("/flows", api_server.GetFlows)

		//////////////////////////////////////////////车载接口//////////////////////////////////////////////
		apiV1.GET("/vehicles/:vehicle_id", api_server.GetVehicle)
		apiV1.POST("/vehicles", api_server.AddVehicle)
		apiV1.PUT("/vehicles/:vehicle_id", api_server.EditVehicle)
		//apiV1.DELETE("/vehicles/:vehicle_id", api_server.DeleVehicle)
		//apiV1.GET("/vehicles", api_server.GetVehicles)
		//////////////////////////////////////////////监控接口//////////////////////////////////////////////
		apiV1.GET("/monitors/:vehicle_id", api_server.GetMonitor)

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

		//资产
		apiV1.GET("/assets", api_server.GetAssets)
		apiV1.GET("/assets/:asset_id", api_server.GetAsset)
		apiV1.POST("/assets", api_server.AddAsset)
		apiV1.PUT("/assets/:asset_id", api_server.EditAsset)
		apiV1.DELETE("/assets/:asset_id", api_server.DeleAsset)

		//////////////////////////////////////////////策略接口//////////////////////////////////////////////
		//添加策略
		apiV1.POST("/strategys", api_server.AddStrategy)
		apiV1.GET("/strategys", api_server.GetStrategys)
		apiV1.GET("/strategys/:strategy_id", api_server.GetStrategy)
		apiV1.DELETE("/strategys/:strategy_id", api_server.DeleStrategy)
		//编辑策略
		apiV1.PUT("/strategys/:strategy_id", api_server.EditStrategy)

		apiV1.GET("/strategy_vehicles/:strategy_id", api_server.GetStrategyVehicle)
		apiV1.GET("/vehicle_lresults/:strategy_vehicle_id", api_server.GetVehicleLearningResults)
		apiV1.GET("/strategy_vehicle_lresults/:strategy_id", api_server.GetStrategyVehicleLearningResults)

		//车载管理信息
		apiV1.PUT("/deployers/:deployer_id", api_server.EditDeployer)

		//端口映射
		apiV1.PUT("/port_maps/:port_map_id", api_server.EditPortMap)
	}
}
