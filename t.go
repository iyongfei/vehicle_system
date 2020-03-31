package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
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


type Name struct{
	Sex interface{}
}

func runFuncName()string{
	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func AAA()  {
	fmt.Println(runFuncName())

}

func main()  {
	N:=Name{nil}
	fmt.Println(N.Sex==nil)
	S:=Name{""}
	fmt.Println(S.Sex == "")
	var a = func() time.Time {
		return time.Unix(0, 0)
	}


	fmt.Println(a)
	return
	r:=runFuncName()

	AAA()
	fmt.Println(r)

	pwd, err := os.Getwd()

	fmt.Println(pwd,err)

	ArgsInvaild      := errors.New("args can be vaild")
	fmt.Println(ArgsInvaild)


	//game:= Name{"sjlkfs"}



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
