package main

import (
	"encoding/json"
)

func main() {

	//rate := fmt.Sprintf("%.3f", 1.0/1.0)
	//frate, _ := strconv.ParseFloat(rate, 64)
	//
	//fmt.Println(frate, frate)

	//fcollect := float64(float64(200)/float64(300)) * 0.2
	//fmt.Println(fcollect)

	//fcollect = float64(float64(distanceTime) / float64(ctime)) * MAX_COLLECT_RATE
	//-1592473012
	//fmt.Println(uint64(1 - 2))

	protoByteMap := map[string]interface{}{}

	_ = json.Unmarshal([]byte(""), &protoByteMap)

}
