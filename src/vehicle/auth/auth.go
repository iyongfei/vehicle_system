package auth

import (
	"encoding/json"
	"fmt"
	"time"
	"vehicle_system/src/vehicle/util"
)

var (
	EmpowerDevices []*EmpowerDevice
)

var (
	AuthFile     = "vehicle_auth.ret" //授权文件存放目录
	BroadcastKey = []byte("v.secure-gAteway-for-Venushalo60")
	BroadcastIV  = []byte("v.Venus-b33csy5F")
)

//授权对象
type EmpowerDevice struct {
	Guid        string //guid
	StartTime   int64  //start stamp
	EndTime     int64  //end stamp
	CompanyName string //公司名称
	EmpowerName string //授权人
}

func Setup() {
	AESDecryptstr()
	//r := AuthVehicleAllExpire()
	//fmt.Println("rrr", r)

	//for k, v := range EmpowerDevices {
	//	fmt.Println("sjdklf", k, v)
	//}
}

func AESDecryptstr() {
	encrypted, err := util.ReadFileByte(AuthFile)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	decrypted, err := util.AESDecrypt(encrypted, BroadcastKey, BroadcastIV)
	devs := []*EmpowerDevice{}
	err = json.Unmarshal(decrypted, &devs)
	EmpowerDevices = devs
}

//1先判断是否授权，在判断是否过期
func AuthVehicleAllExpire() bool {
	var allExpire bool = true
	timeNow := time.Now().Unix()

	for _, empower := range EmpowerDevices {
		endtime := empower.EndTime

		if timeNow < endtime {
			allExpire = false
		}
	}
	return allExpire
}

//2某终端是否授权
func VehicleAuth(vehicleId string) bool {
	aVehicleIdList := AuthVehicleIdList()
	if util.IsExistInSlice(vehicleId, aVehicleIdList) {
		return true
	}
	return false
}

//获取授权的列表
func AuthVehicleIdList() []string {
	authVehicleIdList := []string{}
	for _, empower := range EmpowerDevices {
		vehicleId := empower.Guid

		authVehicleIdList = append(authVehicleIdList, vehicleId)
	}
	return authVehicleIdList
}

//3授权文件是否存在
func VehicleAllUnAuth(vehicleIds []string) bool {

	//授权文件
	exist, _ := util.IsExist(AuthFile)

	//列表为空
	EmpowerNull := EmpowerDevices == nil

	//vehicleId都不在列表中
	aVehicleIdList := AuthVehicleIdList()

	var result []interface{}
	for _, leftE := range vehicleIds { //db
		for _, rightE := range aVehicleIdList { //auth
			if leftE == rightE {
				result = append(result, leftE)
			}
		}
	}

	if /*len(result) == 0 ||*/ !exist || EmpowerNull {
		return true
	}

	return false
}
