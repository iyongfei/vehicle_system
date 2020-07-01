package main

import (
	"encoding/json"
	"fmt"
	"vehicle_system/src/vehicle_script/tool"
)

var (
	EmpowerDevices []*EmpowerDevice
)

// BroadcastKey(32bit) and BroadcastIV(16bit) are two salts for AES encryption.
// Don't rewrite them.
var (
	BroadcastKey = []byte("v.secure-gAteway-for-Venushalo60")
	BroadcastIV  = []byte("v.Venus-b33csy5F")

	ReadPath = "vehicle_auth.ret" //授权文件存放目录
)

//授权对象
type EmpowerDevice struct {
	Guid        string //guid
	StartTime   int64  //start stamp
	EndTime     int64  //end stamp
	CompanyName string //公司名称
	EmpowerName string //授权人
}

const (
	AddTime = 90 * 24 * 3600
	//AddTime = 24 * 3600
)

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
var authVehicleList = []string{
	"1597d612ffc99f61f39a7d491e2d48ec",
	"a42af4258b12288ea4239c3ef6e03888",
	"2833db1cff4e66b116ff7789c1c64cd1",
	"2086b93ded62ba449fd69106f23c0311",
	"fbee8075428a7b86dc8b4dde2b56e951",
	"8ca7f2b8a2d2b31a2c6c2233cab7d9c1",
}

func main() {
	startTimeStamp := tool.Str2Time("2020-07-01 10:00:00")
	//fmt.Println("startTimeStamp:", startTimeStamp)
	//解密授权文件，放入内run

	//多个guid生成授权文件
	t := startTimeStamp

	generateCertFile(authVehicleList, t, t+AddTime, "bohui", "test")
}

//获取guid后用来生成授权证书,存放再ReadPath目录
func generateCertFile(guids []string, start int64, end int64, company string, contacter string) {
	devs := []*EmpowerDevice{}
	for _, v := range guids {
		emp := EmpowerDevice{}
		emp.Guid = v
		emp.StartTime = start
		emp.EndTime = end
		emp.CompanyName = company
		emp.EmpowerName = contacter
		devs = append(devs, &emp)
	}
	bs, _ := json.Marshal(devs)

	encrypted, err := tool.AESEncrypt(bs, BroadcastKey, BroadcastIV)
	if err != nil {
		fmt.Println("AES encrypt error:%v", err, encrypted)
		return
	}

	err = tool.WriteFile(ReadPath, encrypted)
	fmt.Printf("err:%+v", err)
}
