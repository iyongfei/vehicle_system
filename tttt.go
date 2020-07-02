package main

import (
	"fmt"
	"reflect"
)

func init() {

}
func main() {

	var a = 0.1
	r := reflect.TypeOf(a)

	fmt.Println(r)
	fmt.Printf("%s", r)

}
