package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var categoryUrls = map[string]string{
	"post_cate": "http://%s:7001/api/v1/categorys",
	"get_cates": "http://%s:7001/api/v1/all/categorys",
	"edit_cate": "http://%s:7001/api/v1/categorys/%s",
	///
	"get_assets": "http://localhost:7001/api/v1/assets",

	"dele_assets": "http://localhost:7001/api/v1/assets/ypBH0VIQ",
}

func getCateConfig() map[string]string {
	return tool.InitConfig("api_conf.txt")
}

func main() {
	//addCategory()
	getCategorys()
	//editCategory()
	//getAssets()
	//getCategory()

	//deleAsset()

}

/**
//获取所有的车载信息
//*/
func getCategorys() {
	configs := getCateConfig()
	fip := configs["server_ip"]
	token := tool.GetVehicleToken()
	queryParams := map[string]interface{}{}

	urlReq := fmt.Sprintf(categoryUrls["get_cates"], fip)

	resp, _ := tool.Get(urlReq, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
func editCategory() {
	configs := getCateConfig()
	fip := configs["server_ip"]
	cate_id := configs["cate_id"]
	cate_name := configs["cate_name"]
	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(categoryUrls["edit_cate"], fip, cate_id)

	bodyParams := map[string]interface{}{
		"cate_name": cate_name,
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

//
//

func getCategory() {
	configs := getCateConfig()
	fip := configs["server_ip"]
	vehicle_id := configs["vehicle_id"]
	cate_id := configs["cate_id"]

	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(categoryUrls["get_cate"], fip) + cate_id

	queryParams := map[string]interface{}{
		"vehicle_id": vehicle_id,
		"cate_id":    cate_id,
	}

	resp, _ := tool.PostForm(urlReq, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func addCategory() {
	configs := getCateConfig()
	fip := configs["server_ip"]
	//vehicle_id := configs["vehicle_id"]
	cateName := configs["cate_name"]

	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(categoryUrls["post_cate"], fip)

	queryParams := map[string]interface{}{
		"cate_name": cateName,
	}

	resp, _ := tool.PostForm(urlReq, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

//func deleAsset() {
//	token := tool.GetVehicleToken()
//
//	queryParams := map[string]interface{}{}
//
//	reqUrl := assetUrls["dele_assets"]
//
//	resp, _ := tool.Delete(reqUrl, queryParams, token)
//
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}
//
