package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	offLine string


}


func main() {

	jsonRet := `{"offline-qaxnet-lan-168": [{"addr": "1.249.171.83", "version": 4, "OS-EXT-IPS:type": "fixed", "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:8e:f6:3e"}]}`

	tempMap := map[string]interface{}{}
	_ = json.Unmarshal([]byte(jsonRet), &tempMap)


	for k,_:=range tempMap{
		fmt.Println(k)
	}
}