package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var tFlow_urls = map[string]string{
	"get_flow":   "http://192.168.1.103:7001/api/v1/flows/2768455442",
	"get_flows":  "http://192.168.1.103:7001/api/v1/tflows",
	"tflow_dps":  "http://localhost:7001/api/v1/flow_dps",
	"pagination": "http://localhost:7001/api/v1/pagination/flows",
	"post_flows": "http://localhost:7001/api/v1/flows",
	"edit_flows": "http://localhost:7001/api/v1/flows/1111",
	"dele_flows": "http://localhost:7001/api/v1/flows/3648327872",
}

func main() {
	getTflows()
	//getTflowsDps()
}

func getTflowsDps() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": "ada81f6c788e40d4bbb9bfd2ee476a80",
	}

	reqUrl := tFlow_urls["tflow_dps"]

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getTflows() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": "ada81f6c788e40d4bbb9bfd2ee476a80",
	}

	reqUrl := tFlow_urls["get_flows"]

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
