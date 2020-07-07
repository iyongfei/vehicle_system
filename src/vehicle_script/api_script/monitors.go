package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

func main() {
	monitors()
}

var monitorip string
var monitorvehicleId string

func init() {
	apiConfigMap := tool.InitConfig("api_conf.txt")
	monitorip = apiConfigMap["server_ip"]
	monitorvehicleId = apiConfigMap["vehicle_id"]
}

func monitors() {
	token := tool.GetVehicleToken()

	req_url := fmt.Sprintf("http://%s:7001/api/v1/monitors", monitorip)
	bodyParams := map[string]interface{}{
		"vehicle_id": monitorvehicleId,
	}
	resp, _ := tool.Get(req_url, bodyParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
