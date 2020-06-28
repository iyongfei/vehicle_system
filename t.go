package main

import (
	"fmt"
<<<<<<< HEAD
	"github.com/jinzhu/gorm"
	"strconv"
=======
>>>>>>> cq_1.0
	"time"
	"vehicle_system/src/vehicle/util"
)

<<<<<<< HEAD
func BytesToInt32(buf []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(buf))
}

//整形转换成字节
func IntToBytes(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

func UintToBytes(n uint32) []byte {
	var testBytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, n)
	return testBytes
}

func bytesToLittleEndian(testBytes []byte) (ret uint32) {
	ret = binary.LittleEndian.Uint32(testBytes)
	return
}

func Str2Stamp(formatTimeStr string) int64 {
	timeStruct := util.Str2Time(formatTimeStr)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

type CreatedAt time.Time

func (ut *CreatedAt) MarshalJSON() (data []byte, err error) {
	t := strconv.FormatInt(time.Time(*ut).Unix(), 10)
	data = []byte(t)
	return
}

type a struct {
	gorm.Model
	CreatedAt CreatedAt
	Name      string
}

type F func(int)

var aaa F

func main() {

	return

	fmt.Println(time.Now().Unix())
	defaultStartTime := util.GetFewDayAgo(5)
	fmt.Println(defaultStartTime.Unix())
	return
	vehicleFStrategyItemsMap := map[string][]string{}
	vehicleFStrategyItemsMap["a"] = []string{"s", "d"}

	//vehicleFStrategyItemsMap["a"]["b"] = append(vehicleFStrategyItemsMap["a"]["b"], []string{"c", "d"}...)
	//fmt.Println(vehicleFStrategyItemsMap["a"]["b"])

	if v, ok := vehicleFStrategyItemsMap["a"]; ok {
		if !util.IsExistInSlice("ddd", v) {
			vehicleFStrategyItemsMap["a"] = append(vehicleFStrategyItemsMap["a"], "ddd")
		}

	} else {
		fmt.Println("Key Not Found", v, ok)
	}

	fmt.Println(vehicleFStrategyItemsMap)
	return

	//dip := "3232235898"
	//strDIp, _ := strconv.Atoi(dip)
=======
func main() {
	time.Now()
>>>>>>> cq_1.0

	//ip->int
	s := util.StringIpToInt("192.168.1.122")
	fmt.Println("ip->int", s)
	//int->ip
	r := util.IpIntToString(s)
	fmt.Println("int->ip", r)

	//小->大
	sipBigEndian := util.BytesToBigEndian(util.LittleToBytes(uint32(s)))
	//转换////////////////
	dipf := int(sipBigEndian)

	fmt.Println(dipf)

	//大->小
	dipLittleEndian := util.BytesToLittleEndian(util.BigToBytes(uint32(dipf)))
	fss := util.IpIntToString(int(dipLittleEndian))
	fmt.Println(fss)

	//大端
	rr := util.IpIntToString(2046929088)
	fmt.Println("int-->ip", rr)
}
