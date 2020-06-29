package main

import (
	"encoding/json"
	"fmt"
	"time"
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
)

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
var authVehicleList = []string{
	"a42af4258b12288ea4239c3ef6e03888",
	"8ca7f2b8a2d2b31a2c6c2233cab7d9c1",
}

func main() {
	//str2Time := tool.Str2Time("2006-01-02 15:04:05")
	//fmt.Println("str2Time:", str2Time)
	//解密授权文件，放入内run
	//AESDecryptstr()

	//多个guid生成授权文件
	t := time.Now().Unix()

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
