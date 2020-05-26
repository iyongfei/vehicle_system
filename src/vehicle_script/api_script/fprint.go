package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var fprintUrls = map[string]string{
	"get_pagination_fprints": "http://%s:7001/api/v1/pagination/asset_fprints",
	"get_asset_fprints":      "http://%s:7001/api/v1/asset_fprints",
	"post_finger_prints":     "http://%s:7001/api/v1/finger_prints",
	"get_all_fprints":        "http://%s:7001/api/v1/pagination/finger_prints",
	"examine_fprints":        "http://%s:7001/api/v1/pagination/examine/asset_fprints",
	"examine_fprint":         "http://%s:7001/api/v1/examine/asset_fprints/%s",
	//"get_cates": "http://%s:7001/api/v1/all/categorys",
	//"edit_cate": "http://%s:7001/api/v1/categorys/%s",
	//"get_assets": "http://localhost:7001/api/v1/assets",
	//
	"dele_fprints": "http://%s:7001/api/v1/finger_prints/%s",
}

func getFprintConfig() map[string]string {
	return tool.InitConfig("api_conf.txt")
}

func main() {
	//getAssetPaginationFprint() //或者资产指纹信息列表
	//asset_fprints()
	//addFingerPrints()
	getAllFprints()
	//deleFprint()

	//getExamineNetAssetPaginationFprint()
	//examine_fprint()
	//
	//getCategorys()
	//editCategory()
	//getAssets()
	//getCategory()

}

func getAllFprints() {
	token := tool.GetVehicleToken()
	configs := getFprintConfig()
	fip := configs["server_ip"]
	vehicle_id := configs["vehicle_id"]
	//start_time := configs["start_time"]
	//end_time := configs["end_time"]

	queryParams := map[string]interface{}{
		"vehicle_id": vehicle_id,
		//"page_index": page_index,
		//"page_size":  page_size,
		//"start_time": start_time,
		//"end_time":   end_time,
	}
	reqUrl := fmt.Sprintf(fprintUrls["get_all_fprints"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
func examine_fprint() {
	configs := getFprintConfig()
	fip := configs["server_ip"]
	//vehicle_id := configs["vehicle_id"]
	//cateId := configs["cate_id"]
	//assetIds := configs["asset_ids"]

	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(fprintUrls["examine_fprint"], fip, "Ad29vaXCd52A1KxREzBFlLyY6vFOlguG")

	queryParams := map[string]interface{}{
		//"cate_id":   cateId,
		//"asset_ids": assetIds,
	}

	resp, _ := tool.PostForm(urlReq, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))

}

func asset_fprints() {
	token := tool.GetVehicleToken()
	configs := getFprintConfig()
	fip := configs["server_ip"]
	vehicle_id := configs["vehicle_id"]
	//start_time := configs["start_time"]
	//end_time := configs["end_time"]

	queryParams := map[string]interface{}{
		"vehicle_id": vehicle_id,
		//"start_time": start_time,
		//"end_time":   end_time,
	}
	reqUrl := fmt.Sprintf(fprintUrls["get_asset_fprints"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))

}

//获取需要入网审批的资产指纹列表
func getExamineNetAssetPaginationFprint() {
	token := tool.GetVehicleToken()
	configs := getFprintConfig()
	fip := configs["server_ip"]
	vehicle_id := configs["vehicle_id"]
	page_index := configs["page_index"]
	page_size := configs["page_size"]
	//start_time := configs["start_time"]
	//end_time := configs["end_time"]

	queryParams := map[string]interface{}{
		"vehicle_id": vehicle_id,
		"page_index": page_index,
		"page_size":  page_size,
		//"start_time": start_time,
		//"end_time":   end_time,
	}
	reqUrl := fmt.Sprintf(fprintUrls["examine_fprints"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
func deleFprint() {
	token := tool.GetVehicleToken()
	configs := getFprintConfig()
	fip := configs["server_ip"]
	fprint_id := configs["fprint_id"]

	queryParams := map[string]interface{}{}

	reqUrl := fmt.Sprintf(fprintUrls["dele_fprints"], fip, fprint_id)

	resp, _ := tool.Delete(reqUrl, queryParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))

}

func addFingerPrints() {
	configs := getFprintConfig()
	fip := configs["server_ip"]
	//vehicle_id := configs["vehicle_id"]
	cateId := configs["cate_id"]
	assetIds := configs["asset_ids"]

	token := tool.GetVehicleToken()
	urlReq := fmt.Sprintf(fprintUrls["post_finger_prints"], fip)

	queryParams := map[string]interface{}{
		"cate_id":   cateId,
		"asset_ids": assetIds,
	}

	resp, _ := tool.PostForm(urlReq, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getAssetPaginationFprint() {
	token := tool.GetVehicleToken()
	configs := getFprintConfig()
	fip := configs["server_ip"]
	vehicle_id := configs["vehicle_id"]
	page_index := configs["page_index"]
	page_size := configs["page_size"]
	//start_time := configs["start_time"]
	//end_time := configs["end_time"]

	queryParams := map[string]interface{}{
		"vehicle_id": vehicle_id,
		"page_index": page_index,
		"page_size":  page_size,
		//"start_time": start_time,
		//"end_time":   end_time,
	}
	reqUrl := fmt.Sprintf(fprintUrls["get_pagination_fprints"], fip)
	resp, _ := tool.Get(reqUrl, queryParams, token)
	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

//
//func editCategory() {
//	configs := getCateConfig()
//	fip := configs["server_ip"]
//	cate_id := configs["cate_id"]
//	cate_name := configs["cate_name"]
//	token := tool.GetVehicleToken()
//	urlReq := fmt.Sprintf(categoryUrls["edit_cate"], fip, cate_id)
//
//	bodyParams := map[string]interface{}{
//		"cate_name": cate_name,
//	}
//	resp, _ := tool.PutForm(urlReq, bodyParams, token)
//
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}
//
////
////
/////**
////获取所有的车载信息
////*/
//func getCategorys() {
//	configs := getCateConfig()
//	fip := configs["server_ip"]
//	token := tool.GetVehicleToken()
//	queryParams := map[string]interface{}{}
//
//	urlReq := fmt.Sprintf(categoryUrls["get_cates"], fip)
//
//	resp, _ := tool.Get(urlReq, queryParams, token)
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}
//
//func getCategory() {
//	configs := getCateConfig()
//	fip := configs["server_ip"]
//	vehicle_id := configs["vehicle_id"]
//	cate_id := configs["cate_id"]
//
//	token := tool.GetVehicleToken()
//	urlReq := fmt.Sprintf(categoryUrls["get_cate"], fip) + cate_id
//
//	queryParams := map[string]interface{}{
//		"vehicle_id": vehicle_id,
//		"cate_id":    cate_id,
//	}
//
//	resp, _ := tool.PostForm(urlReq, queryParams, token)
//	respMarshal, _ := json.Marshal(resp)
//	fmt.Printf("resp %+v", string(respMarshal))
//}
//

//
////func deleAsset() {
////	token := tool.GetVehicleToken()
////
////	queryParams := map[string]interface{}{}
////
////	reqUrl := assetUrls["dele_assets"]
////
////	resp, _ := tool.Delete(reqUrl, queryParams, token)
////
////	respMarshal, _ := json.Marshal(resp)
////	fmt.Printf("resp %+v", string(respMarshal))
////}
////
