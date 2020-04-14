package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

//https://www.cnblogs.com/zhaof/p/11346412.html

var urls = map[string]string{
	"get_flow":"http://localhost:7001/api/v1/flows/6656653",
	"get_flows":"http://localhost:7001/api/v1/flows",
	"pagination":"http://localhost:7001/api/v1/pagination/flows",
	"post_flows":"http://localhost:7001/api/v1/flows",
	"edit_flows":"http://localhost:7001/api/v1/flows/1111",
	"dele_flows":"http://localhost:7001/api/v1/flows/113907034",
}


func main()  {
	//getFlow()
	//getFlows()
	//getPaginationFlows()
	//addFlows()
	editFlows()
	//deleFlows()
}


func deleFlows()  {
	token := tool.GetVehicleToken()

	queryParams:=map[string]interface{}{
		"vehicle_id":"1234567890123",
	}

	reqUrl:=urls["dele_flows"]

	resp,_:=tool.Delete(reqUrl,queryParams,token)

	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}


func editFlows()  {
	token := tool.GetVehicleToken()
	urlReq, _ := urls["edit_flows"]

	bodyParams := map[string]interface{}{
		"vehicle_id":"1234567890123",
		"src_ip":"3",
		"dst_ip":"42",
	}
	resp,_:=tool.PutForm(urlReq,bodyParams,token)

	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}


func addFlows()  {
	token := tool.GetVehicleToken()
	urlReq, _ := urls["post_flows"]

	bodyParams := map[string]interface{}{
		"vehicle_id":"1234567890123",
		"hash":"1111",
		"src_ip":"1111",
		"dst_ip":"111",
	}
	resp,_:=tool.PostForm(urlReq,bodyParams,token)

	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}


func getPaginationFlows()  {

	token := tool.GetVehicleToken()

	queryParams:=map[string]interface{}{
		"vehicle_id":"1234567890123",
		"page_size":"10",
		"page_index":"1",
	}

	reqUrl:=urls["pagination"]

	resp,_:=tool.Get(reqUrl,queryParams,token)


	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}


func getFlows()  {
	token := tool.GetVehicleToken()

	queryParams:=map[string]interface{}{
		"vehicle_id":"1234567890123",
	}

	reqUrl:=urls["get_flows"]

	resp,_:=tool.Get(reqUrl,queryParams,token)

	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}

/**
flowId:=c.Param("flow_id")
	vehicleId:=c.Query("vehicle_id")
 */
func getFlow()  {
	token := tool.GetVehicleToken()

	queryParams:=map[string]interface{}{
		"vehicle_id":"1234567890123",
	}

	reqUrl:=urls["get_flow"]


	resp,_:=tool.Get(reqUrl,queryParams,token)

	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}