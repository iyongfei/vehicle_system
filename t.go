package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)
var TaskTypeFlowStrategySet = "flowstrategyset"


func main() {

	a:=[]interface{}{"s","sd"}
	r,_:=json.Marshal(a)
	fmt.Println(string(r))
}

//会话策略设置
func createCmdId(args ...string) string {
	var buffer bytes.Buffer

	for _,arg:=range args{
		if strings.Trim(arg, " ") != ""{
			buffer.WriteString(arg)
		}
	}
	return buffer.String()
}