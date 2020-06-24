package main

import "fmt"

func main() {

	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 1,
		"d": 3,
	}

	maxKey := ""
	for k, v := range a {

		maxKey = k

		for k1, v1 := range a {

			if v1 > v {
				maxKey = k1
			}

		}
	}

	fmt.Println(maxKey)
}
