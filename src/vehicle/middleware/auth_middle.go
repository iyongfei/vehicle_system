package middleware

import "github.com/gin-gonic/gin"

func AuthMiddle() gin.HandlerFunc {

	return authMiddleHandlerFunc
}

func authMiddleHandlerFunc(c *gin.Context)  {


}