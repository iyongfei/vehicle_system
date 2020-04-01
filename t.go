package main

import (
	"fmt"
	"strings"
	"time"
	"vehicle_system/src/vehicle/util"
)
/**
go run t.go -logdir "ss"
 */

const (
	LOG_GW_PULL = "pull"
	LOG_GW_PUSH = "push"
	LOG_WEB     = "web"
	LOGDIR = "vlog"
)
const (
	TodayFormat = "2006-01-02"
	TodayTimeFormat = "2006-01-02 15:04:05"
)


func RrgsTrim(args... string) bool {
	var flag = false
	for _,arg:=range args{
		if strings.Trim(arg, " ") == ""{
			 flag = true
			 break
		}
	}

	return flag
}
type H map[string]interface{}

func T(args... string)  {
	//for k,v:=range args{
	//
	//}
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func main()  {

	//user_id := "sef"
	user_name := "userNameg"
	//game:= Name{"sjlkfs"}
	r:=util.CamelCase(user_name)
	fmt.Println(r)




	re:=fmt.Errorf("%s,%d","sj为ldkf",23)
	re1:=fmt.Errorf("%s,%d","sj为ldkf",23)
	fmt.Println(re)
	fmt.Println(re1)

	//fmt.Println(game)

	//fileName := LOGDIR + "/" + "web" + "-" + TimeFormat(TodayFormat) + ".log"
	//
	//
	//_, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//if err!= nil{
	//	log.Fatalf("write2File open logFile err:%s",err)
	//}

}
func TimeFormat(format string) string {
	today := time.Now().Format(format)
	return today
}
