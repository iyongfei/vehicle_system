package api_server

import (
	"github.com/gin-gonic/gin"
	"vehicle_system/src/vehicle/middleware/cors"
	"vehicle_system/src/vehicle/util/logger"
)

func ApiServerStart()  {
	logger.Logger.INFO("%s","ApiServerStart")

	router := gin.Default()
	router.Use(cors.Default())

	router.Run()

}

/**
type HandlerFunc func(*Context)
 */