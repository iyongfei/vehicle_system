package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

func main()  {
	//map[string]interface{}
	req_url:="http://localhost:7001/auth"

	bodyParams := map[string]interface{}{
		"user_name":"safly",
		"password":"838facfab6e49cd2d54d064864ceeb00",
	}
	resp,_:=tool.PostForm(req_url,bodyParams)
	fmt.Println(resp["data"])

	switch resp["data"].(type) {
	case map[string]interface{}:
		ret:=resp["data"].(map[string]interface{})
		token :=ret["token"]
		fmt.Println("token,,,,,,",token)
	}
	
	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}

