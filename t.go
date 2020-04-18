package main

import "fmt"

func main() {

	r:=IsStringExistInSlice(true,[]interface{}{1,"1",true})
	fmt.Println(r)
}


func IsStringExistInSlice(valueParam interface{}, array []interface{}) bool {
	for _, value := range array {
		if value == valueParam {
			return true
		}
	}
	return false
}