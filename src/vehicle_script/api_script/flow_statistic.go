package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

func main() {
	flowStatistic()
}

func flowStatistic() {
	apiConfigMap := tool.InitConfig("api_conf.txt")
	vehicleId := apiConfigMap["flowstatistic_vehicle_id"]
	ip := apiConfigMap["server_ip"]

	req_url := fmt.Sprintf("http://%s:7001/api/v1/flow_statistics", ip)
	bodyParams := map[string]interface{}{
		"vehicle_id": vehicleId,
	}
	resp, _ := tool.PostForm(req_url, bodyParams, "")
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
