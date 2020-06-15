package main

import (
	"fmt"
	"strconv"
)

func main() {

	rate := fmt.Sprintf("%.3f", 1.0/1.0)
	frate, _ := strconv.ParseFloat(rate, 64)

	fmt.Println(frate, frate)
}
