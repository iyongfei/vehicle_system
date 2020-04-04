package main

import (
	_ "vehicle_system/src/vehicle/conf"
	_ "vehicle_system/src/vehicle/db"
	_ "vehicle_system/src/vehicle/emq"
	_ "vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/router"
)

func  main()  {
	router.RouterHandler()
}
