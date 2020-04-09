package main

import (
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/cron"
	"vehicle_system/src/vehicle/emq"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/db"
	"vehicle_system/src/vehicle/timing"

	"vehicle_system/src/vehicle/router"
)

func init() {
	logger.Setup()
	conf.Setup()
	timing.Setup()
	db.Setup()
	emq.Setup()
	cron.Setup()
}


// @title vehicle API
// @version 1.0
// @description This is a sample server vehicle server.
// @termsOfService http://swagger.io/terms/

// @contact.name vehicle API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7001
// @BasePath /
func  main()  {
	router.RouterHandler()
}
