package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"vehicle_system/src/vehicle_script/tool"
)

//https://www.cnblogs.com/zhaof/p/11346412.html
func main()  {
	//getFlow()
	getFlows()
}


func getFlows()  {
	token := tool.GetVehicleToken()
	//fmt.Println(token)
	vehicleId:= "ff"

	urlReq, _ := url.Parse("http://localhost:7001/api/v1/flows")
	params := url.Values{}
	params.Set("vehicle_id",vehicleId)
	urlReq.RawQuery = params.Encode()

	resp,_:=tool.Get(urlReq.String(),token)
	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}

func getFlow()  {
	token := tool.GetVehicleToken()
	//fmt.Println(token)
	vehicleId:= "ff"

	urlReq, _ := url.Parse("http://localhost:7001/api/v1/flows/1")
	params := url.Values{}
	params.Set("vehicle_id",vehicleId)
	urlReq.RawQuery = params.Encode()

	resp,_:=tool.Get(urlReq.String(),token)
	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}