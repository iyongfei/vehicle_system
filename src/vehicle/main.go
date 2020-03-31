package main

import (
	"vehicle_system/src/vehicle/api_server"
	_ "vehicle_system/src/vehicle/conf"
	_ "vehicle_system/src/vehicle/logger"
	_ "vehicle_system/src/vehicle/db"
)

func  main()  {
	api_server.ApiServerStart()
}
