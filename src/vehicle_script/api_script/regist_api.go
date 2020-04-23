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
