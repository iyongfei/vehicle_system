package main

import (
	"fmt"
	"time"
	"vehicle_system/src/vehicle/conf"
)

func main() {
	fmt.Println(time.Now().Unix())

fmt.Println(time.Now().Add(time.Hour * time.Duration(1)).Unix())

	r:=time.Now().Add(conf.Expires * time.Hour).Unix()
	fmt.Println(r,conf.Expires)
}
