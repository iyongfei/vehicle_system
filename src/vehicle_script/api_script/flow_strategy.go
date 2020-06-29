package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"vehicle_system/src/vehicle_script/tool"
)

var fstrategyUrls = map[string]string{
	"post_fstrategy":                   "http://%s:7001/api/v1/fstrategys",
	"dele_fstrategy":                   "http://%s:7001/api/v1/fstrategys/",
	"edit_fstrategy":                   "http://%s:7001/api/v1/fstrategys/",
	"get_fstrategy":                    "http://%s:7001/api/v1/fstrategys/",
	"get_active_fstrategs":             "http://%s:7001/api/v1/active/fstrategys",
	"get_all_fstrategs":                "http://%s:7001/api/v1/ids/fstrategys",
	"get_partial_fstrategs":            "http://%s:7001/api/v1/partial/fstrategys",
	"get_pagination_fstrategs":         "http://%s:7001/api/v1/pagination/fstrategys",
	"get_pagination_vehicle_fstrategs": "http://%s:7001/api/v1/pagination/vehicle/fstrategys",

	//
	"get_fstrategys":               "http://localhost:7001/api/v1/fstrategys",
	"get_strategy_vehicles":        "http://localhost:7001/api/v1/strategy_vehicles/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"get_vehicle_results":          "http://localhost:7001/api/v1/vehicle_lresults/cuMwUiDA2V8NLNWGznfVI2hP5Zi3PhMJ",
	"get_strategy_vehicle_results": "http://localhost:7001/api/v1/strategy_vehicle_lresults/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
}

func getConfig() map[string]string {
	return tool.InitConfig("api_conf.txt")
}

func main() {
	//addFStrategy()
<<<<<<< HEAD
	//getFStrategy()
	deleFStrategy()
	//editFStrategy()
=======
	editFStrategy()
	//deleFStrategy()
	//getFStrategy()
>>>>>>> test
	//getRecentFStrategy()
	//getAllFStrategys()
	//getPartialFStrategys()
	//getPaginationVehicleFstrategys()
	//getPaginationFstrategys()
	//unused
	//getFStrategys()
	//getStrategyVehicle()
	//getVehicleLearningResults()
	//getStrategyVehicleLearningResults()
}

<<<<<<< HEAD
func deleFStrategy() {
	configs := getConfig()
	dele_flow_vehicle_id := configs["get_flow_fstrategy_id"]
	fip := configs["server_ip"]
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

	reqUrl := fmt.Sprintf(fstrategyUrls["dele_fstrategy"], fip) + dele_flow_vehicle_id

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getPaginationFstrategys() {
	token := tool.GetVehicleToken()
	configs := getConfig()
	fip := configs["server_ip"]
=======
func getPaginationVehicleFstrategys() {
	token := tool.GetVehicleToken()
	configs := getConfig()
	fip := configs["server_ip"]
	vehicle_id := configs["vehicle_id"]
>>>>>>> test
	page_index := configs["page_index"]
	page_size := configs["page_size"]
	//start_time := configs["start_time"]
	//end_time := configs["end_time"]

	queryParams := map[string]interface{}{
<<<<<<< HEAD
=======
		"vehicle_id": vehicle_id,
>>>>>>> test
		"page_index": page_index,
		"page_size":  page_size,
		//"start_time": start_time,
		//"end_time":   end_time,
	}
<<<<<<< HEAD
	reqUrl := fmt.Sprintf(fstrategyUrls["get_pagination_fstrategs"], fip)
=======
	reqUrl := fmt.Sprintf(fstrategyUrls["get_pagination_vehicle_fstrategs"], fip)
>>>>>>> test
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func addFStrategy() {
	configs := getConfig()
	flow_vehicle_id := configs["vehicle_id"]
	fips := configs["fips"]
	fports := configs["fports"]
	fip := configs["server_ip"]

	token := tool.GetVehicleToken()
	reqUrl := fstrategyUrls["post_fstrategy"]
	reqUrl = fmt.Sprintf(reqUrl, fip)
	diports := creatFastrategyIpPortData(fips, fports)

	queryParams := map[string]interface{}{
		"vehicle_id": flow_vehicle_id,
		"dip_ports":  diports,
		"name":       "策略名称1",
	}
	fmt.Println("req::::::", reqUrl, diports)

	resp, _ := tool.PostForm(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func editFStrategy() {
	configs := getConfig()
	update_flow_vehicle_id := configs["vehicle_id"]
	update_flow_strategy_id := configs["get_flow_fstrategy_id"]
	fip := configs["server_ip"]
	update_fips := configs["update_fips"]
	update_fports := configs["update_fports"]

	token := tool.GetVehicleToken()
	fmt.Println("token::::::", token)
	urlReq := fmt.Sprintf(fstrategyUrls["edit_fstrategy"], fip) + update_flow_strategy_id

	diports := creatFastrategyIpPortData(update_fips, update_fports)
	queryParams := map[string]interface{}{
		"vehicle_id": update_flow_vehicle_id,
		"dip_ports":  diports,
		"name":       "new新名werwerwwerwe字",
	}

	resp, _ := tool.PutForm(urlReq, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
<<<<<<< HEAD
func getPaginationVehicleFstrategys() {
	token := tool.GetVehicleToken()
	configs := getConfig()
	fip := configs["server_ip"]
	vehicle_id := configs["vehicle_id"]
=======
func getPaginationFstrategys() {
	token := tool.GetVehicleToken()
	configs := getConfig()
	fip := configs["server_ip"]
>>>>>>> test
	page_index := configs["page_index"]
	page_size := configs["page_size"]
	//start_time := configs["start_time"]
	//end_time := configs["end_time"]

	queryParams := map[string]interface{}{
<<<<<<< HEAD
		"vehicle_id": vehicle_id,
=======
>>>>>>> test
		"page_index": page_index,
		"page_size":  page_size,
		//"start_time": start_time,
		//"end_time":   end_time,
	}
	reqUrl := fmt.Sprintf(fstrategyUrls["get_pagination_fstrategs"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getPartialFStrategys() {
	configs := getConfig()
	get_flow_vehicle_id := configs["vehicle_id"]
	partial_fstrategs := configs["partial_fstrategs"]
	fip := configs["server_ip"]
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id":    get_flow_vehicle_id,
		"fstrategy_ids": partial_fstrategs,
	}
	reqUrl := fmt.Sprintf(fstrategyUrls["get_partial_fstrategs"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getAllFStrategys() {
	configs := getConfig()
	get_flow_vehicle_id := configs["vehicle_id"]
	fip := configs["server_ip"]
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": get_flow_vehicle_id,
	}
	reqUrl := fmt.Sprintf(fstrategyUrls["get_all_fstrategs"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

/**
获取一条会话策略信息
*/
func getRecentFStrategy() {
	configs := getConfig()
	get_flow_vehicle_id := configs["vehicle_id"]
	fip := configs["server_ip"]
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": get_flow_vehicle_id,
	}
	reqUrl := fmt.Sprintf(fstrategyUrls["get_active_fstrategs"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

/**
获取一条会话策略信息
*/
func getFStrategy() {
	configs := getConfig()
	get_flow_fstrategy_id := configs["get_flow_fstrategy_id"]
	get_flow_vehicle_id := configs["vehicle_id"]
	fip := configs["server_ip"]
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": get_flow_vehicle_id,
	}
	reqUrl := fmt.Sprintf(fstrategyUrls["get_fstrategy"], fip) + get_flow_fstrategy_id
	resp, _ := tool.Get(reqUrl, queryParams, token)
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
获取所有的会话策略
*/
//func getFStrategys() {
//	token := tool.GetVehicleToken()
//	queryParams := map[string]interface{}{
//		"page_size":  "3",
//		"page_index": "1",
//	}
//	reqUrl := fstrategyUrls["get_fstrategys"]
//	resp, _ := tool.Get(reqUrl, queryParams, token)
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}

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
