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

func main() {
	//解密授权文件，放入内run
	//AESDecryptstr()

	//多个guid生成授权文件
	t := time.Now().Unix()
	fmt.Println(t)
	generateCertFile([]string{"12", "323"}, t, t+24*3600, "", "")

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
	fmt.Println(devs)
	bs, _ := json.Marshal(devs)

	encrypted, err := tool.AESEncrypt(bs, BroadcastKey, BroadcastIV)
	if err != nil {
		fmt.Println("AES encrypt error:%v", err, encrypted)
		return
	}

	tool.WriteFile(ReadPath, encrypted)
}
