package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var fstrategyUrls = map[string]string{
	"post_fstrategy": "http://localhost:7001/api/v1/fstrategys",
	"get_fstrategys": "http://localhost:7001/api/v1/fstrategys",
	"get_fstrategy":  "http://localhost:7001/api/v1/fstrategys/VTkP0Qka6QHkjoy9OE5R079lz7zEa5o1",
	"dele_fstrategy": "http://localhost:7001/api/v1/fstrategys/C4UOrF6oiJoQINfF5gjrQEhZ1CvLCEvY",

	"edit_strategy": "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",

	"get_strategy_vehicles": "http://localhost:7001/api/v1/strategy_vehicles/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"get_vehicle_results": "http://localhost:7001/api/v1/vehicle_lresults/cuMwUiDA2V8NLNWGznfVI2hP5Zi3PhMJ",
	"get_strategy_vehicle_results": "http://localhost:7001/api/v1/strategy_vehicle_lresults/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
}
//apiV1.GET("/fstrategys", api_server.GetFStrategys)
//apiV1.GET("/fstrategys/:fstrategy_id", api_server.GetFStrategy)
//apiV1.DELETE("/fstrategys/:fstrategy_id", api_server.DeleFStrategy)
//apiV1.PUT("/fstrategys/:fstrategy_id", api_server.EditFStrategy)
//apiV1.GET("/fstrategy_vehicle_items/:fstrategy_vehicle_id", api_server.GetVehicleFStrategyItem)
func main() {
	//addFStrategy()
	//getFStrategys()
	//getFStrategy()

	deleFStrategy()
	//editStrategy()//

	//getStrategyVehicle()
	//getVehicleLearningResults()
	//getStrategyVehicleLearningResults()
}

func deleFStrategy() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
	}

	reqUrl := fstrategyUrls["dele_fstrategy"]

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}


/**
获取一条会话策略信息
*/
func getFStrategy() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{

	}
	reqUrl := fstrategyUrls["get_fstrategy"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}


/**
获取所有的会话策略
*/
func getFStrategys() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{
		"page_size":  "3",
		"page_index": "1",
	}
	reqUrl := fstrategyUrls["get_fstrategys"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func addFStrategy() {
	token := tool.GetVehicleToken()
	reqUrl := fstrategyUrls["post_fstrategy"]
	queryParams := map[string]interface{}{
		"vehicle_ids":"TDav,TDavlll,12",
		"type":"1",
		"handle_mode":"2",
		"dips":"192.168.1.1,192.168.1.1",
		"dst_ports":"123,234,344,123",
	}

	resp, _ := tool.PostForm(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
//
//
//
//func getStrategyVehicleLearningResults() {
//	token := tool.GetVehicleToken()
//	queryParams := map[string]interface{}{
//
//	}
//	reqUrl := strategyUrls["get_strategy_vehicle_results"]
//	resp, _ := tool.Get(reqUrl, queryParams, token)
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}
//
//
//func editStrategy() {
//	token := tool.GetVehicleToken()
//	urlReq, _ := strategyUrls["edit_strategy"]
//
//	bodyParams := map[string]interface{}{
//		"type":   "22",
//		"handle_mode": "11",
//		"vehicle_id": "dTtR4sFMYfDJzGAVTv4KWSc9KYLTA64d",
//	}
//	resp, _ := tool.PutForm(urlReq, bodyParams, token)
//
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}
//
//
///**
//获取每一条StrategyVehicle信息
//*/
//func getVehicleLearningResults() {
//	token := tool.GetVehicleToken()
//	queryParams := map[string]interface{}{
//
//	}
//	reqUrl := strategyUrls["get_vehicle_results"]
//	resp, _ := tool.Get(reqUrl, queryParams, token)
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}

//
//
///**
//获取每一条StrategyVehicle信息
//*/
//func getStrategyVehicle() {
//	token := tool.GetVehicleToken()
//	queryParams := map[string]interface{}{
//
//	}
//	reqUrl := strategyUrls["get_strategy_vehicles"]
//	resp, _ := tool.Get(reqUrl, queryParams, token)
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}



