package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var deployerUrls = map[string]string{
	//"post_strategy": "http://localhost:7001/api/v1/strategys",
	//"get_strategys": "http://localhost:7001/api/v1/strategys",
	//"get_strategy":  "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	//
	//"dele_strategy": "http://localhost:7001/api/v1/strategys/9xR5vYZweMb3aRoGGEQYaIw6xhRetYV8",
	"edit_deployer": "http://localhost:7001/api/v1/deployers/9PaWCc7wbEhg2UNQuACGPSyV5BNPRFll",
}

func main()  {
	editDeployer()
}

func editDeployer() {
	token := tool.GetVehicleToken()
	urlReq, _ := deployerUrls["edit_deployer"]

	bodyParams := map[string]interface{}{
		"vehicle_id":   "kQ8XKqP57cNwt0CwXLWoXWF0UfwaYFX8",
		"dev_name": "devname...",
		"name": "name...",
		"phone": "13581922339",
	}
	resp, _ := tool.PutForm(urlReq, bodyParams, token)

	respMarshal, _ := json.Marshal(resp)
	fmt.Printf("resp %+v", string(respMarshal))
}

