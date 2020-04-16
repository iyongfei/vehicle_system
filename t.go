package main

import (
	"bytes"
	"fmt"
	"strings"
)
var TaskTypeFlowStrategySet = "flowstrategyset"


func main() {

	fmt.Println(createCmdId("sdf","wej"))

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