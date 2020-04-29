package main

import (
	"fmt"
	"time"
)

func main() {

	maper := map[string]string{
		"a": "b",
	}
	for k, _ := range maper {
		delete(maper, k)
	}

	fmt.Println(maper)
}

func SubTime(t time.Time) float64 {
	now := time.Now()

	subM := now.Sub(t)

	return subM.Seconds()
}
