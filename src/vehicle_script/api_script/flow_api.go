package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var urls = map[string]string{
	"get_flow":   "http://localhost:7001/api/v1/flows/2768455442",
	"get_flows":  "http://192.168.40.14:7001/api/v1/flows",
	"pagination": "http://localhost:7001/api/v1/pagination/flows",
	"post_flows": "http://localhost:7001/api/v1/flows",
	"edit_flows": "http://localhost:7001/api/v1/flows/1111",
	"dele_flows": "http://localhost:7001/api/v1/flows/3648327872",
}

func main() {
	getFlow()
	//getFlows()
	//getPaginationFlows()
	//addFlows()
	//editFlows()
	//deleFlows()

	//pushFlow()
}

type Requester struct {
	Name string
}

func pushFlow() {
	//token := tool.GetVehicleToken()
	urlReq := "http://localhost:7001/t_flow"

	req := Requester{
		Name: "safly",
	}
	resp, _ := tool.PostJson(urlReq, req, "")

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
func getFlow() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": "754d2728b4e549c5a16c0180fcacb800",
	}

	reqUrl := urls["get_flow"]

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func deleFlows() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": "754d2728b4e549c5a16c0180fcacb800",
	}

	reqUrl := urls["dele_flows"]

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
func editFlows() {
	token := tool.GetVehicleToken()
	urlReq, _ := urls["edit_flows"]

	bodyParams := map[string]interface{}{
		"vehicle_id": "b020eccdf33d48b4aa246a89a6f04609",
		"src_ip":     "3",
		"dst_ip":     "42",
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getPaginationFlows() {

	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": "754d2728b4e549c5a16c0180fcacb800",
		"page_size":  "2",
	}

	reqUrl := urls["pagination"]

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getFlows() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": "b020eccdf33d48b4aa246a89a6f04609",
	}

	reqUrl := urls["get_flows"]

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func addFlows() {
	token := tool.GetVehicleToken()
	urlReq, _ := urls["post_flows"]

	bodyParams := map[string]interface{}{
		"vehicle_id": "1234567890123",
		"hash":       "1111",
		"src_ip":     "1111",
		"dst_ip":     "111",
	}
	resp, _ := tool.PostForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
