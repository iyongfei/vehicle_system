package main

import (
	"encoding/json"
	"fmt"
	"time"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle_script/tool"
)

var vehicleUrls = map[string]string{
	"get_vehicle":  "http://localhost:7001/api/v1/vehicles/b020eccdf33d48b4aa246a89a6f04609",
	"get_vehicles": "http://localhost:7001/api/v1/vehicles",

	"post_vehicles": "http://localhost:7001/api/v1/vehicles",

	"edit_vehicles": "http://localhost:7001/api/v1/vehicles/derXH5DghbCV3UVHFQaCNbmHitQHcTfj",
	"dele_vehicles": "http://localhost:7001/api/v1/vehicles/WDHIAeGImCklIqrzQ2fBfojPL0kg4D7d",
}

func main() {
	//getVehicles()
	//getVehicle()
	//addVehicle()
	//deleVehicles()
	editVehicles()
}

func editVehicles() {
	token := tool.GetVehicleToken()
	urlReq, _ := vehicleUrls["edit_vehicles"]

	bodyParams := map[string]interface{}{
		"type":   "1",
		"switch": "true",
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

	resp, _ := tool.Get(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
