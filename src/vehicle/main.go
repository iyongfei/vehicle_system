package main

import (
	"vehicle_system/src/vehicle/api_server"
	_ "vehicle_system/src/vehicle/conf"
	_ "vehicle_system/src/vehicle/util/logger"
)

func  main()  {
	api_server.ApiServerStart()
}
