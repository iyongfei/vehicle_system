package v1

import (
	"github.com/gin-gonic/gin"
	"vehicle_system/src/vehicle/api_server"
	"vehicle_system/src/vehicle/middleware"
)

func V1Router(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.AuthMiddle())
	{

		//////////////////////////////////////////////流量接口//////////////////////////////////////////////
		apiV1.GET("/flow_statistics", api_server.FlowStatistics) //todo
		//////////////////////////////////////////////监控接口//////////////////////////////////////////////
		apiV1.GET("/monitors", api_server.GetMonitor) //todo

		//////////////////////////////////////////////会话策略接口//////////////////////////////////////////////
		apiV1.POST("/fstrategys", api_server.AddFStrategy)                  //todo
		apiV1.DELETE("/fstrategys/:fstrategy_id", api_server.DeleFStrategy) //todo
		apiV1.PUT("/fstrategys/:fstrategy_id", api_server.EditFStrategy)    //todo
		apiV1.GET("/fstrategys/:fstrategy_id", api_server.GetFStrategy)     //todo
		//////////////////////////////////////////////会话策略下载上传//////////////////////////////////////////////
		apiV1.GET("/fstrategy_csvs/:fstrategy_id", api_server.GetFStrategyCsv)  //todo
		apiV1.POST("/fstrategy_csvs", api_server.UploadFStrategyCsv)            //todo
		apiV1.PUT("/fstrategy_csvs/:fstrategy_id", api_server.EditFStrategyCsv) //todo
		//////////////////////////////////////////////会话接口//////////////////////////////////////////////
		apiV1.GET("/flows/:flow_id", api_server.GetFlow)              //todo
		apiV1.GET("/pagination/flows", api_server.GetPaginationFlows) //todo
		apiV1.GET("/flow_type_counts", api_server.GetFlowTypeCounts)  //todo

		apiV1.GET("/tflow_dps", api_server.GetTFlowsDps) //todo
		apiV1.GET("/tflows", api_server.GetTFlows)       //todo

		apiV1.GET("/flow_dps", api_server.GetFlowsDps) //todo

		//apiV1.DELETE("/flows/:flow_id", api_server.DeleFlow)
		//apiV1.POST("/flows", api_server.AddFlow)
		//apiV1.PUT("/flows/:flow_id", api_server.EditFlow)
		//apiV1.GET("/flows", api_server.GetFlows)

		//////////////////////////////////////////////指纹库标签//////////////////////////////////////////////
		apiV1.POST("/categorys", api_server.AddCategory)             //todo
		apiV1.GET("/all/categorys", api_server.GetCategorys)         //todo
		apiV1.PUT("/categorys/:cate_id", api_server.EditCategory)    //todo
		apiV1.DELETE("/categorys/:cate_id", api_server.DeleCategory) //todo
		//////////////////////////////////////////////指纹信息接口//////////////////////////////////////////////
		//apiV1.GET("/pagination/asset_fprints", api_server.GetPaginationAssetFprints) //todo
		apiV1.GET("/asset_fprints", api_server.GetAssetFprints) //todo
		//////////////////////////////////////////////指纹库接口//////////////////////////////////////////////
		apiV1.POST("/finger_prints", api_server.AddFprint)               //todo
		apiV1.GET("/all/finger_prints", api_server.GetFprints)           //todo
		apiV1.DELETE("/finger_prints/:fprint_id", api_server.DeleFprint) //todo
		apiV1.PUT("/finger_prints/:fprint_id", api_server.EditFprint)    //todo
		//////////////////////////////////////////////入网审批，允许入网///////////////////////P///////////////////////
		apiV1.GET("/pagination/examine/asset_fprints", api_server.GetExamineAssetFprints)        //todo
		apiV1.POST("/examine/asset_fprints/:asset_fprint_id", api_server.AddExamineAssetFprints) //todo
		//////////////////////////////////////////////车载接口//////////////////////////////////////////////
		apiV1.GET("/vehicles/:vehicle_id", api_server.GetVehicle) //todo
		apiV1.POST("/vehicles", api_server.AddVehicle)
		apiV1.PUT("/vehicles/:vehicle_id", api_server.EditVehicle) //todo
		//apiV1.DELETE("/vehicles/:vehicle_id", api_server.DeleVehicle)
		apiV1.GET("/vehicles", api_server.GetVehicles)

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
