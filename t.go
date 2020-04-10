package main

import (
	"fmt"
	"reflect"
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
	LOGDIR      = "vlog"
)
const (
	TodayFormat     = "2006-01-02"
	TodayTimeFormat = "2006-01-02 15:04:05"
)

func RrgsTrim(args ...string) bool {
	var flag = false
	for _, arg := range args {
		if strings.Trim(arg, " ") == "" {
			flag = true
			break
		}
	}

	return flag
}

type H map[string]interface{}

func T(args ...string) {
	//for k,v:=range args{
	//
	//}
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

type Name struct {
	Sex    string
	Id     int
	School string
}

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}
func Stamp2Str(stamp int64) string{
	timeLayout := "2006-01-02 15:04:05"
	str:=time.Unix(stamp,0).Format(timeLayout)
	return str
}

type A struct {
	TT time.Time
}

func main() {
	a:=A{}
	fmt.Println(a.TT)
	fmt.Println(Stamp2Str(10000000000))

}

func AA(data interface{}) map[interface{}]interface{} {
	dataValue := reflect.ValueOf(data)
	typeF := dataValue.Type()
	//queryMap:=map[string]interface{}{}
	queryMap := map[interface{}]interface{}{}
	//nFiled:=dataValue.Type().NumField()
	for i := 0; i < dataValue.NumField(); i++ {
		field := dataValue.Field(i)
		//fmt.Println(field)
		fmt.Println(typeF.Field(i).Name, field.Interface())
		queryMap[field.Interface()] = typeF.Field(i).Name
	}

	return queryMap
}


func TimeFormat(format string) string {
	  today := time.Now().Format(format)
	return today
}
