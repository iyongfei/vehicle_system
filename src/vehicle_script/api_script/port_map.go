package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var portMapUrls = map[string]string{
	//"post_strategy": "http://localhost:7001/api/v1/strategys",
	//"get_strategys": "http://localhost:7001/api/v1/strategys",
	//"get_strategy":  "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	//"dele_strategy": "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"edit_port_map": "http://localhost:7001/api/v1/port_maps/W80uZD9MRbzk7Z5xUKQN0xIIetZk1Oqj",
}

func main()  {
	editPortMap()
}

func editPortMap() {
	token := tool.GetVehicleToken()
	urlReq, _ := portMapUrls["edit_port_map"]

	bodyParams := map[string]interface{}{
		"vehicle_id":"fds",
		"src_port":"77",
		"dest_port":"777",
		"dest_ip":"192.168.1.18",
		"switch":"true",
		"protocol":"3",
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

