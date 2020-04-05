package tool

import (
	"fmt"
	"math/rand"
	"time"
)
func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

func TimeNowToUnix()uint32{
	t:=time.Now()
	return  uint32(t.Unix())
}

func GenIpAddr() string {
	rand.Seed(time.Now().UnixNano())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}
func GenVersion() string {
	rand.Seed(time.Now().UnixNano())
	version := fmt.Sprintf("%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10))
	return version
}


func GenBrandType(ln int) uint8 {
	letters := []rune("12")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	var ret uint8
	switch string(b) {
	case "1":
		ret =  0
	case "2":
		ret = 1
	}
	return ret
}


func RandomAlternativeBool(ln int) bool {
	letters := []rune("12")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	var ret bool
	switch string(b) {
	case "1":
		ret =  true
	case "2":
		ret = false
	}
	return ret
}


func GenBrand(ln int) string {
	letters := []rune("1234")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	var ret string
	switch string(b) {
	case "1":
		ret =  "apple"
	case "2":
		ret = "huawei"
	case "3":
		ret = "小米"
	case "4":
		ret = "格力"
	}
	return ret
}

