package tool

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
生成特定位数随机数
*/
func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

/**
生成几位随机数字
*/
func RandomNumber(ln int) int {
	letters := []rune("1234567890")
	lettersPre := []rune("123456789")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		if i == 0 {
			b[i] = letters[r.Intn(len(lettersPre))]
		} else {
			b[i] = letters[r.Intn(len(letters))]
		}
	}
	ret, _ := strconv.Atoi(string(b))
	return ret
}

/**
生成多少以内的数字
*/
func RandToMaxNumber(ln int) int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(pow(10, ln)))
	return r
}

/**
次方
*/
func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

//传10，生成[1,10]的随机数
func RandOneToMaxNumber(max int) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(max))
	r++
	return r
}

//[0,10)
func RandOneToMaxNumberT(max int) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(max))
	return r
}

func GenVersion() string {
	rand.Seed(time.Now().UnixNano())
	version := fmt.Sprintf("%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10))
	return version
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
		ret = true
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
		ret = "apple"
	case "2":
		ret = "huawei"
	case "3":
		ret = "小米"
	case "4":
		ret = "格力"
	}
	return ret
}
