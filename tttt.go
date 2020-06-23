package main

import (
	"fmt"
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

	stdFpProtoFlowMap := map[string]int{
		"a": 1,
		"b": 3,
		"c": 1,
		"d": 14,
		"e": 5,
	}

	maxKey := ""
	for k, max := range stdFpProtoFlowMap {
		maxKey = k
		for k1, v1 := range stdFpProtoFlowMap {
			if v1 > max {
				maxKey = k1
				max = v1
			}
		}
		break
	}

	fmt.Println(maxKey)
}
