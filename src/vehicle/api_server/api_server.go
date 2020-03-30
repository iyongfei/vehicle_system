package api_server

import (
	"github.com/gin-gonic/gin"
	"vehicle_system/src/vehicle/middleware/cors"
)

func ApiServerStart()  {

	router := gin.Default()
	router.Use(cors.Default())

	router.Run()

}

/**
type HandlerFunc func(*Context)
 */