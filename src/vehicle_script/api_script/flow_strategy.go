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
	"dele_fstrategy": "http://localhost:7001/api/v1/fstrategys/Ne114FumZ61ju946sBT3Mdr65PjZYv95",

	"edit_fstrategy": "http://localhost:7001/api/v1/fstrategys/xyhMowqwDQUCBtofp25Z2hP4CdDvDwk3",

	"get_strategy_vehicles":        "http://localhost:7001/api/v1/strategy_vehicles/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"get_vehicle_results":          "http://localhost:7001/api/v1/vehicle_lresults/cuMwUiDA2V8NLNWGznfVI2hP5Zi3PhMJ",
	"get_strategy_vehicle_results": "http://localhost:7001/api/v1/strategy_vehicle_lresults/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
}

func getConfig() map[string]string {
	return tool.InitConfig("api_conf.txt")
}

func main() {
	addFStrategy()
	//getFStrategys()
	//getFStrategy()

	//deleFStrategy()
	//editFStrategy() //

	//getStrategyVehicle()
	//getVehicleLearningResults()
	//getStrategyVehicleLearningResults()
}

func editFStrategy() {
	token := tool.GetVehicleToken()
	urlReq, _ := fstrategyUrls["edit_fstrategy"]

	queryParams := map[string]interface{}{
		"vehicle_id":  "gBDvsD4qnAdR9K4cTIZoq7Zr05Ox1udr",
		"type":        "1",
		"handle_mode": "2",
		"dips":        "192.190.1.1,192.108.1.3",
		"dst_ports":   "34,35",
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

func deleFStrategy() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

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
