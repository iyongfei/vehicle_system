package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle_script/tool"
)

var assetUrls = map[string]string{
	"get_assets_simple": "http://%s:7001/api/v1/all/assets",
	"get_assets":        "http://%s:7001/api/v1/pagination/assets",
	"get_asset":         "http://%s:7001/api/v1/assets/%s",
	"post_assets":       "http://localhost:7001/api/v1/assets",

	"edit_assets":             "http://%s:7001/api/v1/assets/%s",
	"dele_assets":             "http://localhost:7001/api/v1/assets/ypBH0VIQ",
	"edit_vehicle_asset_name": "http://%s:7001/api/v1/assets/%s/asset_info",
}

var config map[string]string

func getAssetConfig() {
	config = tool.InitConfig("api_conf.txt")
}
func init() {
	getAssetConfig()
}

func main() {
	//getAssetsimple()
	//getAssets()
	//getAsset()
	//addAsset()
	//deleAsset()
	editAsset()

	//editVehicleAssetInfo()
}

func editAsset() {
	fip := config["server_ip"]
	assetId := config["asset_id"]
	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(assetUrls["edit_assets"], fip, assetId)
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
func getAsset() {
	fip := config["server_ip"]
	assetId := config["asset_id"]

	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(assetUrls["get_asset"], fip, assetId)

	queryParams := map[string]interface{}{}
	resp, _ := tool.Get(urlReq, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
func getAssetsimple() {
	fip := config["server_ip"]

	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{}
	urlReq := fmt.Sprintf(assetUrls["get_assets_simple"], fip)
	resp, _ := tool.Get(urlReq, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))

}

/**
获取所有的车载信息
*/
func getAssets() {
	fip := config["server_ip"]
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{
		"page_size":  "3",
		"page_index": "1",
		"vehicle_id": "1",
	}
	reqUrl := fmt.Sprintf(assetUrls["get_assets"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func editVehicleAssetInfo() {
	fip := config["server_ip"]
	assetId := config["asset_id"]
	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(assetUrls["edit_vehicle_asset_name"], fip, assetId)

	vehicle_id := config["vehicle_id"]
	asset_new_name := config["asset_new_name"]

	bodyParams := map[string]interface{}{
		"vehicle_id": vehicle_id,
		"name":       asset_new_name,
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

//
//func getAssets() {
//	fip := config["server_ip"]
//
//	token := tool.GetVehicleToken()
//	queryParams := map[string]interface{}{}
//	urlReq := fmt.Sprintf(assetUrls["get_assets_simple"], fip)
//	resp, _ := tool.Get(urlReq, queryParams, token)
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//
//}

func deleAsset() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

	reqUrl := assetUrls["dele_assets"]

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func addAsset() {
	token := tool.GetVehicleToken()
	reqUrl := assetUrls["post_assets"]
	queryParams := &model.Asset{
		VehicleId:      "cQ9U6wV1Zpj7dAH9aa9rmOzLD6JAEKCE",
		AssetId:        tool.RandomString(8),
		IP:             tool.RandomString(8),
		Mac:            tool.RandomString(8),
		Name:           tool.RandomString(8),
		TradeMark:      tool.RandomString(8),
		OnlineStatus:   true,
		LastOnline:     tool.TimeNowToUnix(),
		InternetSwitch: true,
		ProtectStatus:  true,
		LanVisitSwitch: true,
	}

	resp, _ := tool.PostJson(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
