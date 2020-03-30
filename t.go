package main

import (
	"errors"
	"fmt"
	"os"
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

func (e Name) Errorr() string {
	return e.Sex
}

type Name struct{
	Sex string
}


func main()  {

	pwd, err := os.Getwd()

	fmt.Println(pwd,err)

	ArgsInvaild      := errors.New("args can be vaild")
	fmt.Println(ArgsInvaild)


	game:= Name{"sjlkfs"}




	fmt.Println(game)

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
