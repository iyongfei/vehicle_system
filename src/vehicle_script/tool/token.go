package tool

//map[string]interface{}

func GetVehicleToken() string {
	reqUrl := "http://localhost:7001/auth"

	bodyParams := map[string]interface{}{
		"user_name": "safly",
		"password":  "4c35c166cc5d28cb96ad5c606cd2f263",
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
