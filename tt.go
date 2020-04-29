package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "tianqi-R201b-967E6D9A3001/disconnected"
	//patter := ".*/disconnecte"
	//r, err := regexp.MatchString(patter, s)
	//fmt.Println(r, err)

	r := strings.HasSuffix("ssswe", "we")
	fmt.Println(r)

}
