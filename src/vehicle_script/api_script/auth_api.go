package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

func main() {
	auth()
}

func auth() {
	apiConfigMap := tool.InitConfig("api_conf.txt")
	user_name := apiConfigMap["user_name"]
	password := apiConfigMap["password"]
	ip := apiConfigMap["server_ip"]

	fmt.Println(user_name, password, ip)

	req_url := fmt.Sprintf("http://%s:7001/auth", ip)
	bodyParams := map[string]interface{}{
		"user_name": user_name,
		"password":  password,
	}
	resp, _ := tool.PostForm(req_url, bodyParams, "")
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
