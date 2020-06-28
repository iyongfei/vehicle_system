package main

import (
	"fmt"
	"strings"
)

type AuthVehicle struct {
	Start     uint32
	End       uint32
	VehicleId string
}

type AuthVehicleList struct {
	Server       string
	AuthVehicles []AuthVehicle
}

var aggg []*AuthVehicleList

func main() {
	a := "wjek"
	b := "wjekeee"

	c := strings.Index(a, b)
	fmt.Println(c)
}
