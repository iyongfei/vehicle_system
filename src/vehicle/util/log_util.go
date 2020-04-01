package util

import (
	"fmt"
	"log"
)

var (
	//								 0x1B  [    9  色值 ;   4 背景色值 m
	greenBgYellow      	= string([]byte{27, 91, 57, 51, 59, 52, 50, 109})
	greenBgRed      	= string([]byte{27, 91, 57, 49, 59, 52, 50, 109})
	greenBg      		= string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	grayBg      		= string([]byte{27, 91, 57, 55, 59, 52, 55, 109})
	yellowBg     		= string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	redBg        		= string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg       		= string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg    		= string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg       		= string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green        		= string([]byte{27, 91, 51, 50, 109})
	gray        		= string([]byte{27, 91, 51, 55, 109})
	yellow       		= string([]byte{27, 91, 51, 51, 109})
	red          		= string([]byte{27, 91, 51, 49, 109})
	blue         		= string([]byte{27, 91, 51, 52, 109})
	magenta      		= string([]byte{27, 91, 51, 53, 109})
	cyan         		= string([]byte{27, 91, 51, 54, 109})
	reset        		= string([]byte{27, 91, 48, 109})
	disableColor 		= false
)

func Vprintf(tag string, format string, v ...interface{}) {
	format = "(%s)" + format
	v = append([]interface{}{tag}, v...)
	log.Printf(format, v...)
}

func Vfmtf(tag string, format string, v ...interface{}){
	format = "(%s)" + format
	v = append([]interface{}{tag}, v...)
	fmt.Printf(format,v...)
}

func VfmtfSimple(format string, v ...interface{})  {
	v = append([]interface{}{}, v...)
	fmt.Printf(format,v...)
}

func Vfmt(a ...interface{}){
	fmt.Println(a...)
}

func Vspritf(format string, v ...interface{})  string{
	return  fmt.Sprintf(format,v...)
}


func RedbackgroundPrintf(a ...interface{})  {
	fmt.Println(redBg,a)
}

