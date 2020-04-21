package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"vehicle_system/src/vehicle_script/tool"
)

var fstrategyUrls = map[string]string{
	"post_fstrategy": "http://localhost:7001/api/v1/fstrategys",
	"get_fstrategys": "http://localhost:7001/api/v1/fstrategys",
	"get_fstrategy":  "http://localhost:7001/api/v1/fstrategys/VTkP0Qka6QHkjoy9OE5R079lz7zEa5o1",
	"dele_fstrategy": "http://localhost:7001/api/v1/fstrategys/",

	"edit_fstrategy": "http://localhost:7001/api/v1/fstrategys/",

	"get_strategy_vehicles":        "http://localhost:7001/api/v1/strategy_vehicles/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"get_vehicle_results":          "http://localhost:7001/api/v1/vehicle_lresults/cuMwUiDA2V8NLNWGznfVI2hP5Zi3PhMJ",
	"get_strategy_vehicle_results": "http://localhost:7001/api/v1/strategy_vehicle_lresults/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
}

func getConfig() map[string]string {
	return tool.InitConfig("api_conf.txt")
}

func main() {
	//addFStrategy()
	//getFStrategys()
	//getFStrategy()
	//deleFStrategy()
	editFStrategy()

	//getStrategyVehicle()
	//getVehicleLearningResults()
	//getStrategyVehicleLearningResults()
}

func deleFStrategy() {
	configs := getConfig()
	dele_flow_vehicle_id := configs["dele_flow_strategy_id"]

	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

	reqUrl := fstrategyUrls["dele_fstrategy"] + dele_flow_vehicle_id

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func editFStrategy() {
	configs := getConfig()
	update_flow_vehicle_id := configs["update_flow_vehicle_id"]
	update_flow_strategy_id := configs["update_flow_strategy_id"]
	update_fips := configs["update_fips"]
	update_fports := configs["update_fports"]

	token := tool.GetVehicleToken()
	urlReq := fstrategyUrls["edit_fstrategy"] + update_flow_strategy_id

	diports := creatFastrategyIpPortData(update_fips, update_fports)
	queryParams := map[string]interface{}{
		"vehicle_id": update_flow_vehicle_id,
		"dip_ports":  diports,
	}

	resp, _ := tool.PutForm(urlReq, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func addFStrategy() {
	configs := getConfig()
	flow_vehicle_id := configs["flow_vehicle_id"]
	fips := configs["fips"]
	fports := configs["fports"]

	token := tool.GetVehicleToken()
	reqUrl := fstrategyUrls["post_fstrategy"]

	diports := creatFastrategyIpPortData(fips, fports)

	queryParams := map[string]interface{}{
		"vehicle_id": flow_vehicle_id,
		"dip_ports":  diports,
	}
	resp, _ := tool.PostForm(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func creatFastrategyIpPortData(fips string, fports string) string {

	ipList := strings.Split(fips, ",")
	portList := strings.Split(fports, ",")

	data := map[string][]uint32{}

	for i := 0; i < len(ipList); i++ {
		ip := ipList[i]

		var ipPort []uint32
		for j := 0; j < len(portList); j++ {
			port := portList[j]
			rt, _ := strconv.Atoi(port)
			ipPort = append(ipPort, uint32(rt))
		}
		data[ip] = ipPort
	}

	ret, _ := json.Marshal(data)
	return string(ret)
}

/**
获取一条会话策略信息
*/
func getFStrategy() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{}
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
