package api_server

import (
	"github.com/gin-gonic/gin"
	"vehicle_system/src/vehicle/middleware/cors"
	"vehicle_system/src/vehicle/util/logger"
)

type T struct {
	Name string
	Id int
	Sex string
}

func ApiServerStart()  {
	t:=T{
		"name",1,"sjdlkf",
	}
	logger.Logger.Info("%+v,%+v,%s",t,t,"sjdflk")
	logger.Logger.Print("%+v,%+v,%s",t,t,"sjdfgggglk")

	router := gin.Default()
	router.Use(cors.Default())

	router.Run()

}

/**
type HandlerFunc func(*Context)
 */