package main

import (
	"io"
	"os"
	"vehicle_system/src/vehicle_script/tool"
)

var fStrategyCsvUrls = map[string]string{
	"csv_url": "http://localhost:7001/fstrategy_csv/N5gqNSN0lpV30gKJOfBkYvGudNUfj1V5.csv",

	"get_fstrategy_csv": "http://localhost:7001/api/v1/fstrategy_csvs/N5gqNSN0lpV30gKJOfBkYvGudNUfj1V5",
	//"post_strategy": "http://localhost:7001/api/v1/strategys",
	"get_strategys": "http://localhost:7001/api/v1/strategys",

	"dele_strategy": "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"edit_strategy": "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",

	"get_strategy_vehicles":        "http://localhost:7001/api/v1/strategy_vehicles/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"get_vehicle_results":          "http://localhost:7001/api/v1/vehicle_lresults/cuMwUiDA2V8NLNWGznfVI2hP5Zi3PhMJ",
	"get_strategy_vehicle_results": "http://localhost:7001/api/v1/strategy_vehicle_lresults/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
}

func main() {
	//getFstrategyCsv()
	getFstrategyCsvTemp()
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
