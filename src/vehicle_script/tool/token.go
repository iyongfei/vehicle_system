package tool

import "fmt"

//map[string]interface{}

func GetVehicleToken() string {

	apiConfigMap := InitConfig("api_conf.txt")
	ip := apiConfigMap["server_ip"]

	user_name := apiConfigMap["user_name"]
	password := apiConfigMap["password"]

	reqUrl := fmt.Sprintf("http://%s:7001/auth", ip)

	bodyParams := map[string]interface{}{
		"user_name": user_name,
		"password":  password,
	}
	resp, _ := PostForm(reqUrl, bodyParams, "")

	var token string
	switch resp["data"].(type) {
	case map[string]interface{}:
		ret := resp["data"].(map[string]interface{})
		token = ret["token"].(string)
		return token
	}

	return token
}
