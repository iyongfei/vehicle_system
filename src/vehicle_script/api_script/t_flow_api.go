package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var tFlow_urls = map[string]string{
	"get_flows": "http://%s:7001/api/v1/tflows",
	"tflow_dps": "http://%s:7001/api/v1/flow_dps",
}
var tflowip string
var tflowvehicleId string

func init() {
	apiConfigMap := tool.InitConfig("api_conf.txt")
	tflowip = apiConfigMap["server_ip"]
	tflowvehicleId = apiConfigMap["vehicle_id"]
}

func main() {
	//getTflows()
	getTflowsDps()
}

func getTflowsDps() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": tflowvehicleId,
	}

	reqUrl := tFlow_urls["tflow_dps"]
	reqUrl = fmt.Sprintf(reqUrl, tflowip)

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getTflows() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": tflowvehicleId,
	}

	reqUrl := tFlow_urls["get_flows"]
	reqUrl = fmt.Sprintf(reqUrl, tflowip)

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
