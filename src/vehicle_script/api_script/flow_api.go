package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"vehicle_system/src/vehicle_script/tool"
)

var urls = map[string]string{
	"get_flow":   "http://%s:7001/api/v1/flows/%s",
	"get_flows":  "http://%s:7001/api/v1/flows",
	"flow_dps":   "http://%s:7001/api/v1/flow_dps",
	"pagination": "http://%s:7001/api/v1/pagination/flows",
	"post_flows": "http://localhost:7001/api/v1/flows",
	"edit_flows": "http://localhost:7001/api/v1/flows/1111",
	"dele_flows": "http://localhost:7001/api/v1/flows/3648327872",
}

func main() {
	//getFlow()
	getPaginationFlows()
	getflowsDps()

	//unused
	//getFlows()
	//addFlows()

	//editFlows()
	//deleFlows()

	//pushFlow()
}

var flowip string
var flowId string

var flowvehicleId string

func init() {
	apiConfigMap := tool.InitConfig("api_conf.txt")
	flowip = apiConfigMap["server_ip"]
	flowvehicleId = apiConfigMap["vehicle_id"]
	flowId = apiConfigMap["flowid"]
}

func getflowsDps() {
	defaultStartTime := GetFewDayAgo(10) //2
	now := strconv.Itoa(int(time.Now().Unix()))

	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": flowvehicleId,
		"start_time": defaultStartTime,
		"end_time":   now,
	}

	reqUrl := urls["flow_dps"]
	reqUrl = fmt.Sprintf(reqUrl, flowip)

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getFlows() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": flowvehicleId,
	}

	reqUrl := urls["get_flows"]
	reqUrl = fmt.Sprintf(reqUrl, flowip)

	fmt.Println(reqUrl, token)

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func GetFewDayAgo(days int) string {
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -days).Unix()
	r := strconv.Itoa(int(oldTime))
	return r
}

func getPaginationFlows() {
	defaultStartTime := GetFewDayAgo(10) //2
	now := strconv.Itoa(int(time.Now().Unix()))

	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": flowvehicleId,
		"start_time": defaultStartTime,
		"end_time":   now,
	}

	reqUrl := urls["pagination"]
	reqUrl = fmt.Sprintf(reqUrl, flowip)

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
func getFlow() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"vehicle_id": flowvehicleId,
	}

	reqUrl := urls["get_flow"]
	reqUrl = fmt.Sprintf(reqUrl, flowip, flowId)

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
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
