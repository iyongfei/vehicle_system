package tool

//map[string]interface{}

func GetVehicleToken() string {
	reqUrl:="http://localhost:7001/auth"

	bodyParams := map[string]interface{}{
		"user_name":"safly",
		"password":"838facfab6e49cd2d54d064864ceeb00",
	}
	resp,_:=PostForm(reqUrl,bodyParams)

	var token string
	switch resp["data"].(type) {
	case map[string]interface{}:
		ret:=resp["data"].(map[string]interface{})
		token = ret["token"].(string)
		return token
	}
	return token
}