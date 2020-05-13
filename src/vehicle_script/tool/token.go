package tool

//map[string]interface{}

func GetVehicleToken() string {
	reqUrl := "http://192.168.1.103:7001/auth"

	bodyParams := map[string]interface{}{
		"user_name": "saflyer",
		"password":  "123456",
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
