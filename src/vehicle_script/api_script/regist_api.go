package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

func main() {

	regist()

}

func regist() {
	apiConfigMap := tool.InitConfig("api_conf.txt")

	user_name := apiConfigMap["user_name"]
	ip := apiConfigMap["server_ip"]
	password := apiConfigMap["password"]

	req_url := fmt.Sprintf("http://%s:7001/regist", ip)

	bodyParams := map[string]interface{}{
		"user_name": user_name,
		"password":  password,
	}
	resp, _ := tool.PostForm(req_url, bodyParams, "")
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

//curl http://192.168.40.234:8320/regist -X POST -d "user_name=aa&password=aa"

//curl http://192.168.40.234:8320/api/v1/flow_statistics?vehicle_id=e4aa43208d213dc1a4372185a7774fcc -X GET -H "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNnc2dGRNd0pHMjZHZElQSEZIczhsWG1QckhIek5pS1ciLCJ1c2VyX25hbWUiOiJhYSIsInBhc3Nfd29yZCI6ImVhOWU1ZmE2ZGVjNzY1ZWEzNDM3MzlkZWE2YTA4NThmIiwiZXhwIjoxNjY1NDMxNzQyLCJpc3MiOiJ2ZWhpY2xlIn0.8_nC9gIloOKKXfzu1Mljx-u0330thTTd77sfo3sq1BI"
