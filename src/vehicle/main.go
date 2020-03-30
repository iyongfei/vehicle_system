package main

import (
	"vehicle_system/src/vehicle/api_server"
	"vehicle_system/src/vehicle/util/logger"
)

func  main()  {
	Logger:=logger.NewRealStLogger(0)
	Logger.DEBUG("sdjfl%s","sdfs")
	api_server.ApiServerStart()
}
