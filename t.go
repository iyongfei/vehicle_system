package main

import (
	"fmt"
	"time"
	"vehicle_system/src/vehicle/util"
)

func main() {
	time.Now()

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
