package main

import (
	"fmt"
	"strings"
)

func main() {

	rt:=RrgsTrimsAllEmpty(" "," d ","")
	fmt.Println(rt)
}


func RrgsTrimsAllEmpty(args... string) bool {
	nullCount := 0
	var flag = false
	for _,arg:=range args{
		if strings.Trim(arg, " ") == ""{
			nullCount ++
		}
	}
	if nullCount == len(args){
		flag = true
	}

	return flag
}
