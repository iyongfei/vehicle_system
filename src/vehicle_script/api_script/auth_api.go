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

	req_url := fmt.Sprintf("http://%s:7001/auth", ip)
	bodyParams := map[string]interface{}{
		"user_name": user_name,
		"password":  password,
	}
	resp, _ := tool.PostForm(req_url, bodyParams, "")
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZG53MTdBd0pNTWJpYzE2NnhHb3daUGZ1cWNiY2xqSkIiLCJ1c2VyX25hbWUiOiJhIiwicGFzc193b3JkIjoiYjYyZTM4OTc5MzE3NzhlNjk2OGFhZjBlZWVlZjMyZDYiLCJleHAiOjE2NjM2OTE3NjgsImlzcyI6InZlaGljbGUifQ.wVR9n2FzCSNDKgvCmwwWzwNaWwjTql9ehnc1NqYklPE

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZG53MTdBd0pNTWJpYzE2NnhHb3daUGZ1cWNiY2xqSkIiLCJ1c2VyX25hbWUiOiJhIiwicGFzc193b3JkIjoiYjYyZTM4OTc5MzE3NzhlNjk2OGFhZjBlZWVlZjMyZDYiLCJleHAiOjE2NjM2OTE4NjksImlzcyI6InZlaGljbGUifQ.kTEJAmic14s0305EGKhwXbuezrxqJUIwEHyARjYdVzo
