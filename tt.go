package main

import (
	"fmt"
	"time"
)

func main() {

	maper := map[string]string{
		"a": "b",
	}
	delete(maper, "b")
	fmt.Println(maper)
}

func SubTime(t time.Time) float64 {
	now := time.Now()

	subM := now.Sub(t)

	return subM.Seconds()
}
