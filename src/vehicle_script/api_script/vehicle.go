package main

import (
	"encoding/json"
	"fmt"
	"time"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle_script/tool"
)

var vehicleUrls = map[string]string{
	"get_vehicle":  "http://%s:7001/api/v1/vehicles/%s",
	"get_vehicles": "http://%s:7001/api/v1/pagination/vehicles",

	"post_vehicles": "http://localhost:7001/api/v1/vehicles",

	"edit_vehicles":     "http://%s:7001/api/v1/vehicles/%s",
	"edit_vehicle_name": "http://%s:7001/api/v1/vehicles/%s/vehicle_info",
	"dele_vehicles":     "http://localhost:7001/api/v1/vehicles/WDHIAeGImCklIqrzQ2fBfojPL0kg4D7d",
}

var ip string
var vehicleId string
var vehicleName string

func init() {
	apiConfigMap := tool.InitConfig("api_conf.txt")
	ip = apiConfigMap["server_ip"]
	vehicleId = apiConfigMap["vehicle_id"]
	vehicleName = apiConfigMap["vehicle_new_name"]
}

func main() {
	//editVehicleInfo()
	//getVehicle()
	//editVehicles()
	getVehicles()
	//unused
	//addVehicle()
	//deleVehicles()
}

/**
获取所有的车载信息
*/
func getVehicles() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{
		"page_size":  "5",
		"page_index": "1",
	}
	reqUrl := vehicleUrls["get_vehicles"]
	reqUrl = fmt.Sprintf(reqUrl, ip)

	fmt.Println(reqUrl, "tok,,,,,,,,,,,")

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func editVehicleInfo() {
	token := tool.GetVehicleToken()
	urlReq, _ := vehicleUrls["edit_vehicle_name"]

	urlReq = fmt.Sprintf(urlReq, ip, vehicleId)

	bodyParams := map[string]interface{}{
		"name": vehicleName,
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func editVehicles() {
	token := tool.GetVehicleToken()
	urlReq, _ := vehicleUrls["edit_vehicles"]

	urlReq = fmt.Sprintf(urlReq, ip, vehicleId)

	bodyParams := map[string]interface{}{
		"type":   "1",
		"switch": "false",
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

/**
获取一条车载信息

*/
func getVehicle() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

	reqUrl := vehicleUrls["get_vehicle"]
	reqUrl = fmt.Sprintf(reqUrl, ip, vehicleId)

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func deleVehicles() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

	reqUrl := vehicleUrls["dele_vehicles"]

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func addVehicle() {
	token := tool.GetVehicleToken()

	reqUrl := vehicleUrls["post_vehicles"]

	queryParams := &model.VehicleInfo{
		VehicleId: tool.RandomString(32),
		Name:      tool.RandomString(8),
		Version:   tool.GenVersion(),
		//StartTime:model_base.UnixTime(time.Now()),
		StartTime:        time.Now(),
		FirmwareVersion:  tool.RandomString(8),
		HardwareModel:    tool.RandomString(8),
		Module:           tool.RandomString(8),
		SupplyId:         tool.RandomString(8),
		UpRouterIp:       tool.GenIpAddr(),
		Type:             1,
		Mac:              tool.RandomString(8),
		TimeStamp:        tool.TimeNowToUnix(),
		HbTimeout:        88,
		DeployMode:       1,
		FlowIdleTimeSlot: 23,
		OnlineStatus:     true,
		ProtectStatus:    1,
		LeaderId:         tool.RandomString(8),
		GroupId:          tool.RandomString(8),
	}

	resp, _ := tool.PostJson(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
