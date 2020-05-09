package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"vehicle_system/src/vehicle_script/tool"
)

func main() {

	ip, port := initConfIni()
	router := gin.Default()

	router.POST("/flowstats", TFlowStats)
	router.POST("/flow_statistics", TFlowStatistics)
	router.POST("/monitor_infos", TMonitorInfos)
	router.POST("/vehicle_infos", TVehicleInfos)
	router.POST("/protects", TProtects)
	router.POST("/online_status", TOnlineStatus)

	router.Run(fmt.Sprintf("%s:%s", ip, port))
}

func initConfIni() (string, string) {
	apiConfigMap := tool.InitConfig("conf.ini")
	push_host := apiConfigMap["push_host"]
	push_port := apiConfigMap["push_port"]

	fmt.Println("ini_pushhost-->", push_host, ",ini_pushport-->", push_port)

	return push_host, push_port
}

//./push_script -ip 192.168.1.1 -port 7001
//func initConfIni() (string, string) {
//	var ip string
//	var port string
//
//	flag.StringVar(&ip, "h", "192.1", "")
//	flag.StringVar(&port, "p", "", "")
//
//	flag.Parse()
//
//	fmt.Println("ip==>", ip, ",port==>", port)
//	return ip, port
//}

func TFlowStats(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TFlowStats", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)

	//log
	fileName := time.Now().Format("20180102") + ".txt"
	outfile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开

	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("TFlowStats:%s \n", bodyJsonStr) //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

	defer outfile.Close()

	c.JSON(http.StatusOK, tempMap)
}

func TFlowStatistics(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TFlowStatistics", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)

	//log
	fileName := time.Now().Format("20180102") + ".txt"
	outfile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开

	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("TFlowStatistics:%s \n", bodyJsonStr) //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

	defer outfile.Close()

	c.JSON(http.StatusOK, tempMap)
}

func TMonitorInfos(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TMonitorInfos", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)

	//log
	fileName := time.Now().Format("20180102") + ".txt"
	outfile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开

	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("TMonitorInfos:%s \n", bodyJsonStr) //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

	defer outfile.Close()

	c.JSON(http.StatusOK, tempMap)
}

func TVehicleInfos(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TVehicleInfos", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)
	//log
	fileName := time.Now().Format("20180102") + ".txt"
	outfile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开

	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("TVehicleInfos:%s \n", bodyJsonStr) //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

	defer outfile.Close()

	c.JSON(http.StatusOK, tempMap)
}

func TProtects(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TProtects", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)
	//log
	fileName := time.Now().Format("20180102") + ".txt"
	outfile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开

	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("TProtects:%s \n", bodyJsonStr) //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

	defer outfile.Close()

	c.JSON(http.StatusOK, tempMap)
}

func TOnlineStatus(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	bodyJsonStr := string(body)
	fmt.Println("TOnlineStatus", bodyJsonStr)

	var tempMap map[string]interface{}

	_ = json.Unmarshal([]byte(bodyJsonStr), &tempMap)
	//log
	fileName := time.Now().Format("20180102") + ".txt"
	outfile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开

	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("TOnlineStatus:%s \n", bodyJsonStr) //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

	defer outfile.Close()

	c.JSON(http.StatusOK, tempMap)
}
