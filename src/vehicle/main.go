package main

import (
	"fmt"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/cron"
	"vehicle_system/src/vehicle/db"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/mac"
	"vehicle_system/src/vehicle/router"
	"vehicle_system/src/vehicle/service/push"
	"vehicle_system/src/vehicle/timing"
	"vehicle_system/src/vehicle/vgo"
)

func init() {
	logger.Setup()
	conf.Setup()
	mac.Setup()
	timing.Setup()
	db.Setup()
	emq.Setup()
	cron.Setup()
	push.Setup()
	vgo.Setup()

	fprintsMacs := []string{}
	//查找指纹库所有的mac
	_ = mysql.QueryRawsqlScanVariable("finger_prints",
		"device_mac", &fprintsMacs, "", []interface{}{}...)

	fmt.Println(fprintsMacs, "---------------->")
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
func main() {
	router.RouterHandler()

}
