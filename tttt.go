package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var weightRate float64
	hostnameWeight := 0.2
	fpHost := `["tffes","gggg"]`
	assetCateStdHost := `["tffes","ggg","s"]`
	var fpHostslice []string
	_ = json.Unmarshal([]byte(fpHost), &fpHostslice)

	var assetCateStdHostslice []string
	_ = json.Unmarshal([]byte(assetCateStdHost), &assetCateStdHostslice)

	hostCommonMap := []string{}
	for _, stdhost := range assetCateStdHostslice {
		for _, host := range fpHostslice {
			if host == stdhost {
				hostCommonMap = append(hostCommonMap, host)
			}
		}
	}
	fmt.Println(hostCommonMap)
	weightRate += float64(len(hostCommonMap)) / float64(len(assetCateStdHostslice)) * hostnameWeight

	fmt.Println(weightRate)
}
