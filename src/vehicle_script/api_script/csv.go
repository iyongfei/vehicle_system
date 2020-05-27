package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"vehicle_system/src/vehicle_script/tool"
)

var fStrategyCsvUrls = map[string]string{
	"csv_url":                 "http://localhost:7001/fstrategy_csv/3832kYyxG3uD9DhF9VDvV5HwLLyhrAkG.csv",
	"get_fstrategy_csv":       "http://localhost:7001/api/v1/fstrategy_csvs/HFiYobVy2dqYiVcGpcsrk6GVRxUdqpuy",
	"post_strategy_csv":       "http://localhost:7001/api/v1/fstrategy_csvs",
	"edit_strategy_csv":       "http://localhost:7001/api/v1/fstrategy_csvs/RaZ8yLTOjDqybBrsQ7Tf5i3ZOJQhKfK9",
	"post_asset_fprints_csvs": "http://%s:7001/api/v1/asset_fprints_csvs",
	//////////////////////////////////////////////
	"get_strategys": "http://localhost:7001/api/v1/strategys",

	"dele_strategy": "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",

	"get_strategy_vehicles":        "http://localhost:7001/api/v1/strategy_vehicles/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"get_vehicle_results":          "http://localhost:7001/api/v1/vehicle_lresults/cuMwUiDA2V8NLNWGznfVI2hP5Zi3PhMJ",
	"get_strategy_vehicle_results": "http://localhost:7001/api/v1/strategy_vehicle_lresults/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
}

func getConfiger() map[string]string {
	return tool.InitConfig("api_conf.txt")
}
func main() {
	//getFstrategyCsv()
	//getFstrategyCsvTemp()
	//uploadFstrategyCsv()
	//editFstrategyCsv()

	//上传资产白名单
	uploadAssetPrintCsv()
}

func uploadAssetPrintCsv() {
	configer := getConfiger()
	ip := configer["server_id"]
	token := tool.GetVehicleToken()
	reqUrl := fmt.Sprintf(fStrategyCsvUrls["post_asset_fprints_csvs"], ip)
	mapArgs := map[string]string{}

	nameField := "upload_csv"
	fileName := "upload_csver"
	file, _ := os.Open("/Users/mac/go/vehicle_system/whte_asset_print.csv")

	resp, _ := tool.UploadFile(reqUrl, token, mapArgs, nameField, fileName, file)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func editFstrategyCsv() {
	url := fStrategyCsvUrls["edit_strategy_csv"]

	nameField := "upload_csv"
	fileName := "upload_csver"

	mapArgs := map[string]string{
		"vehicle_id": "754d2728b4e549c5a16c0180fcacb800",
	}

	file, _ := os.Open("/Users/mac/go/vehicle_system/safly.csv")

	resp, _ := tool.UploadEditFile(url, mapArgs, nameField, fileName, file)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

func getFstrategyCsvTemp() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

	reqUrl := fStrategyCsvUrls["get_fstrategy_csv"]
	res, err := tool.GetFile(reqUrl, queryParams, token)
	if err != nil {
	}
	defer res.Body.Close()

	f, _ := os.Create("safly.csv")

	io.Copy(f, res.Body)

}

func getFstrategyCsv() {
	token := tool.GetVehicleToken()

	queryParams := map[string]interface{}{}

	reqUrl := fStrategyCsvUrls["csv_url"]

	resp, _ := tool.GetFile(reqUrl, queryParams, token)
	defer resp.Body.Close()
	//写入文件

	f, _ := os.Create("safly.csv")

	io.Copy(f, resp.Body)
}

func uploadFstrategyCsv() {
	token := tool.GetVehicleToken()
	url := fStrategyCsvUrls["post_strategy_csv"]
	mapArgs := map[string]string{
		"vehicle_id": "754d2728b4e549c5a16c0180fcacb800",
	}

	nameField := "upload_csv"
	fileName := "upload_csver"
	file, _ := os.Open("/Users/mac/go/vehicle_system/safly.csv")

	resp, _ := tool.UploadFile(url, token, mapArgs, nameField, fileName, file)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}
