package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var vehicleUrls = map[string]string{
	"get_vehicle":"http://localhost:7001/api/v1/vehicles/DyG7IvSc4ds9vNkRuNCzmbwpJrG8MQeH",
	"get_vehicles":"http://localhost:7001/api/v1/vehicles",

	"post_vehicles":"http://localhost:7001/api/v1/vehicles",
	"edit_vehicles":"http://localhost:7001/api/v1/vehicles/1111",
	"dele_vehicles":"http://localhost:7001/api/v1/vehicles/113907034",
}


func main()  {
	//getVehicles()
	getVehicle()

	//getPaginationFlows()
	//addFlows()
	//editFlows()
	//deleFlows()
}



/**
获取一条车载信息
 */
func getVehicle()  {
	token := tool.GetVehicleToken()

	queryParams:=map[string]interface{}{
	}

	reqUrl:=vehicleUrls["get_vehicle"]

	resp,_:=tool.Get(reqUrl,queryParams,token)

	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}


/**
获取所有的车载信息
 */
func getVehicles()  {
	token := tool.GetVehicleToken()

	queryParams:=map[string]interface{}{
		"page_size":"5",
		"page_index":"1",
	}

	reqUrl:=vehicleUrls["get_vehicles"]

	resp,_:=tool.Get(reqUrl,queryParams,token)

	respMarshal ,_:= json.Marshal(resp)
	fmt.Printf("resp %+v",string(respMarshal))
}