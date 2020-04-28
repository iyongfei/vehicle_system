package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonRet := `{"offline-qaxnet-lan-168": [{"addr": "1.249.171.83", "version": 4, "OS-EXT-IPS:type": "fixed", "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:8e:f6:3e"},
{"addr": "1.249.171.99", "version": 4, "OS-EXT-IPS:type": "fixed", "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:8e:f6:3e"}]}`

	tempMap := map[string][]map[string]interface{}{}
	_ = json.Unmarshal([]byte(jsonRet), &tempMap)

	var addrList []string

	for _, ListMapV := range tempMap {
		for _, obj := range ListMapV {
			addrList = append(addrList, obj["addr"].(string))
		}
	}

	fmt.Printf("addr ret:%+v", addrList)
}
