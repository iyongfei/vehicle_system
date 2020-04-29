package main

import (
	"fmt"
	"time"
)

func main() {

	s := 0.000000000000000001

	fmt.Println(s)
}

func SubTime(t time.Time) float64 {
	now := time.Now()

	subM := now.Sub(t)

	return subM.Seconds()
}
