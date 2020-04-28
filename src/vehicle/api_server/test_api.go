package api_server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func TFlowStats(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TFlowStats", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)

	c.JSON(http.StatusOK, tempMap)
}

func TFlowStatistics(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TFlowStatistics", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)

	c.JSON(http.StatusOK, tempMap)
}

func TMonitorInfos(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TMonitorInfos", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)

	c.JSON(http.StatusOK, tempMap)
}

func TVehicleInfos(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TVehicleInfos", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)

	c.JSON(http.StatusOK, tempMap)
}
