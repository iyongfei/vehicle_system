package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle_script/mysql_script/tools"
)

/**
授权脚本
*/

func init() {

	tools.GetMysqlInstance()
}

const PasswordSecret = "vgw-1214-pwd-key-vgw-1214-pwd-key"

func main() {
	TdataVehicleAuthCheck()
}

/**
初始化授权列表
*/
func TdataVehicleAuthCheck() error {
	vehicleAuths := GetVehicleAuths()
	for _, vehicleId := range vehicleAuths {
		vehicleIdMd5 := Md5(vehicleId + PasswordSecret)

		vehicleAuth := &VehicleAuth{
			VehicleId: vehicleIdMd5,
		}

		recordNotFound := tools.QueryFirstModelRecord(vehicleAuth, "vehicle_id = ?", []interface{}{vehicleAuth.VehicleId}...)

		//不存在插入
		if recordNotFound {
			err := tools.CreateModel(vehicleAuth)
			if err != nil {
				continue
			}
		}
	}

	return nil
}

func GetVehicleAuths() []string {
	vehicleAuths := []string{
		"ad29258414587f8aab117932a470df45",
	}
	return vehicleAuths
}

type VehicleAuth struct {
	gorm.Model
	VehicleId string //关联的小v ID
}

//生成32位md5字串
func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
