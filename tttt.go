package main

import (
	"fmt"
)

func init() {

}

func main() {

	assetCateMarkMap := map[string]int{
		"a": 2,
		"b": 10,
		"c": 4,
		"d": 8,
	}

	var maxAssetIdKey string
	for assetId, value := range assetCateMarkMap {
		fmt.Println("aaa")
		maxAssetIdKey = assetId
		for tmpAssetId, tmpValue := range assetCateMarkMap {
			fmt.Println("bbb", tmpAssetId, tmpValue)

			if tmpValue > value {
				maxAssetIdKey = tmpAssetId
				value = tmpValue
			}
		}
		break
	}

	fmt.Println("sss", maxAssetIdKey)

}
