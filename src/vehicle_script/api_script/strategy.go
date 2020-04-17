package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var strategyUrls = map[string]string{
	"post_strategy": "http://localhost:7001/api/v1/strategys",
	"get_strategys": "http://localhost:7001/api/v1/strategys",
	"get_strategy":  "http://localhost:7001/api/v1/strategys/RfL4oNP3VpwsFfqOGNDjuC0FeqTynMqV",
	"dele_strategy": "http://localhost:7001/api/v1/strategys/RfL4oNP3VpwsFfqOGNDjuC0FeqTynMqV",
	"edit_strategy": "http://localhost:7001/api/v1/strategys/xer1bSYURVf7NgSIOwTveBtnvl0dErrH",

	"get_strategy_vehicles": "http://localhost:7001/api/v1/strategy_vehicles/opeuBHjxvP3EW16gD5VXJus7RbrJPNb3",
	"get_strategy_vehicle_results": "http://localhost:7001/api/v1/strategy_vehicle_lresults/vidwjeklflw",
}

func main() {
	//addStrategy()
	//getStrategys()
	//getStrategy()

	//deleStrategy()
	//editStrategy()

	//getStrategyVehicle()
	getStrategyVehicleLearningResults()
}


/**
获取每一条StrategyVehicle信息
*/
func getStrategyVehicleLearningResults() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{

	}
	reqUrl := strategyUrls["get_strategy_vehicle_results"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}


/**
获取每一条StrategyVehicle信息
*/
func getStrategyVehicle() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{

	}
	reqUrl := strategyUrls["get_strategy_vehicles"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}


func editStrategy() {
	token := tool.GetVehicleToken()
	urlReq, _ := strategyUrls["edit_strategy"]

	bodyParams := map[string]interface{}{
		"type":   "2",
		"handle_mode": "1",
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}


func deleStrategy() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
	}

	reqUrl := strategyUrls["dele_strategy"]

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

/**
获取一条策略信息
*/
func getStrategy() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{

	}
	reqUrl := strategyUrls["get_strategy"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
//
///**
//获取所有的车载信息
// */
func getStrategys() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{
		"page_size":  "3",
		"page_index": "1",
	}
	reqUrl := strategyUrls["get_strategys"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}



func addStrategy() {
	token := tool.GetVehicleToken()
	reqUrl := strategyUrls["post_strategy"]
	queryParams := map[string]interface{}{
		"type":"1",
		"handle_mode":"2",
		"learning_result_ids":"1,3,4",
		"vehicle_id":"vidwjeklflw",
	}

	resp, _ := tool.PostForm(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}


