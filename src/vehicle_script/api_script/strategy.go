package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var strategyUrls = map[string]string{
	"post_strategy": "http://localhost:7001/api/v1/strategys",
	"get_strategys": "http://localhost:7001/api/v1/strategys",
	"get_strategy":  "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",

	"dele_strategy": "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"edit_strategy": "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",

	"get_strategy_vehicles": "http://localhost:7001/api/v1/strategy_vehicles/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"get_vehicle_results": "http://localhost:7001/api/v1/vehicle_lresults/cuMwUiDA2V8NLNWGznfVI2hP5Zi3PhMJ",
	"get_strategy_vehicle_results": "http://localhost:7001/api/v1/strategy_vehicle_lresults/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
}

func main() {
	addStrategy()
	//getStrategys()
	//getStrategy()

	//deleStrategy()//协议部分没有处理
	//editStrategy()//协议部分没有处理

	//getStrategyVehicle()
	//getVehicleLearningResults()
	//getStrategyVehicleLearningResults()
}
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
func getVehicleLearningResults() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{

	}
	reqUrl := strategyUrls["get_vehicle_results"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
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

func addStrategy() {
	token := tool.GetVehicleToken()
	reqUrl := strategyUrls["post_strategy"]
	queryParams := map[string]interface{}{
		"type":"1",
		"handle_mode":"2",
		"learning_result_ids":"1,3,2,4",
		"vehicle_ids":"tc3ijhYbUI0B2ZiRK6qdlA5QtiXDrfnz,dTtR4sFMYfDJzGAVTv4KWSc9KYLTA64d,v2",
	}

	resp, _ := tool.PostForm(reqUrl, queryParams, token)
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



