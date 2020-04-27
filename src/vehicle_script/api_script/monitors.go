package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

func main() {
	monitors()
}

func monitors() {
	apiConfigMap := tool.InitConfig("api_conf.txt")
	vehicleId := apiConfigMap["monitors_vehicle_id"]

	req_url := "http://localhost:7001/api/v1/monitors"
	bodyParams := map[string]interface{}{
		"vehicle_id": vehicleId,
	}
	resp, _ := tool.Get(req_url, bodyParams, "")
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
