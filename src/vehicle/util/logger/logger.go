package logger

import (
	"errors"
	"fmt"
)

var Logger *VLogger
var loggerErrInfo = errors.New("vlogger init error")

func init()  {
	var loggerErr error
	Logger,loggerErr = NewRealStLogger()
	if loggerErr!=nil || Logger == nil{
		fmt.Println(loggerErrInfo)
	}
}
