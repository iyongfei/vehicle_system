package main

import (
	"fmt"
	"strings"
	"time"
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
	msg:="$SYS/brokers/emqx@127.0.0.1/clients/tianqi-R201b-967E6D9A3001/disconnected"
	set:="$SYS/brokers/emqx@127.0.0.1/clients/"

	topIndex := strings.Index(msg,set)
	fmt.Println(topIndex)

	topicSlice:=strings.Split(msg,set)
	fmt.Println("topicSlice::",topicSlice)
	topicSlice_1 :=topicSlice[1]
	fmt.Println("topicSlice_1::",topicSlice_1)
	gwId := strings.Split(topicSlice_1,"/")[0]

	fmt.Println(gwId)

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
