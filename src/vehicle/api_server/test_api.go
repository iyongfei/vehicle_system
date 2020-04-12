package api_server

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
)

func TFlow(c *gin.Context)  {




	body,_:= ioutil.ReadAll(c.Request.Body)
	strBody:= string(body)
	fmt.Println("TFlow",strBody)


}
