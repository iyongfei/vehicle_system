package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle_script/tool"
)

var assetUrls = map[string]string{
	"get_assets":        "http://localhost:7001/api/v1/assets",
	"get_asset":         "http://localhost:7001/api/v1/assets/YP4wZffU",
	"post_assets":       "http://localhost:7001/api/v1/assets",
	"post_white_assets": "http://%s:7001/api/v1/white/assets",

	"edit_assets": "http://localhost:7001/api/v1/assets/XdUylhnx",
	"dele_assets": "http://localhost:7001/api/v1/assets/ypBH0VIQ",
}

func getAssetConfig() map[string]string {
	return tool.InitConfig("api_conf.txt")
}

func main() {
	addWhiteAsset()
	//getAssets()
	//getAsset()
	//addAsset()
	//deleAsset()
	//editAsset()
}

func addWhiteAsset() {
	configs := getAssetConfig()
	fip := configs["server_ip"]
	asset_ids := configs["white_asset_ids"]

	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(assetUrls["post_white_assets"], fip)

	queryParams := map[string]interface{}{
		"asset_ids": asset_ids,
	}

	resp, _ := tool.PostForm(urlReq, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func deleAsset() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

	reqUrl := assetUrls["dele_assets"]

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func editAsset() {
	token := tool.GetVehicleToken()
	urlReq, _ := assetUrls["edit_assets"]

	bodyParams := map[string]interface{}{
		"type":   "1",
		"switch": "true",
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

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

/**
获取一条车载信息
*/
func getAsset() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{}
	reqUrl := assetUrls["get_asset"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

/**
获取所有的车载信息
*/
func getAssets() {
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{
		"page_size":  "3",
		"page_index": "1",
	}
	reqUrl := assetUrls["get_assets"]
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
