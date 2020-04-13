package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"vehicle_system/src/vehicle_script/tool"
)

//https://www.cnblogs.com/zhaof/p/11346412.html

var urls = map[string]string{
	"get_flow":"http://localhost:7001/api/v1/flows/1",
	"get_flows":"http://localhost:7001/api/v1/flows",
	"pagination":"http://localhost:7001/api/v1/pagination/flows",
	"post_flows":"http://localhost:7001/api/v1/flows",
}

var (
	page_size = "2"
	page_index = "2"
	vehicleId = "ff"
)
func main()  {
	//getFlow()
	//getFlows()
	//getPaginationFlows()
	addFlows()


}

type Flow struct {
	FlowId          uint32
	VehicleId       string
	Hash            uint32
	SrcIp           uint32
	SrcPort         uint32
	DstIp           uint32
	DstPort         uint32
	Protocol        uint8
	FlowInfo        string
	SafeType        uint8
	SafeInfo        string
	StartTime       uint32
	LastSeenTime    uint32
	SrcDstBytes     uint64
	DstSrcBytes     uint64
	Stat            uint8
}

func addFlows()  {
	token := tool.GetVehicleToken()
	urlReq, _ := url.Parse(urls["post_flows"])

	body:= map[string]interface{}{
		//"VehicleId":
	}

	resp,_:=tool.Post(urlReq.String(),body,token)

	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}
/**
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
 */

func getPaginationFlows()  {
	token := tool.GetVehicleToken()

	urlReq, _ := url.Parse(urls["pagination"])
	params := url.Values{}
	params.Set("vehicle_id",vehicleId)
	params.Set("page_size",page_size)
	params.Set("page_index",page_index)
	urlReq.RawQuery = params.Encode()

	resp,_:=tool.Get(urlReq.String(),token)
	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}


func getFlows()  {
	token := tool.GetVehicleToken()
	urlReq, _ := url.Parse(urls["get_flows"])
	params := url.Values{}
	params.Set("vehicle_id",vehicleId)
	urlReq.RawQuery = params.Encode()

	resp,_:=tool.Get(urlReq.String(),token)
	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}

func getFlow()  {
	token := tool.GetVehicleToken()
	urlReq, _ := url.Parse(urls["get_flow"])
	params := url.Values{}
	params.Set("vehicle_id",vehicleId)
	urlReq.RawQuery = params.Encode()

	resp,_:=tool.Get(urlReq.String(),token)
	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}