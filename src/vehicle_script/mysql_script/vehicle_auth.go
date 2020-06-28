package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
	"vehicle_system/src/vehicle_script/mysql_script/tools"
)

/**
授权脚本
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build
*/

func init() {

	tools.GetMysqlInstance()
}

const PasswordSecret = "vgw-1214-pwd-key-vgw-1214-pwd-key"
const AuthDays = 30

func main() {

	var extraVehicleAuth string
	flag.StringVar(&extraVehicleAuth, "v", "", "vid")
	flag.Parse()

	extraVehicleAuth = strings.Trim(extraVehicleAuth, " ")

	TdataVehicleAuthCheck(extraVehicleAuth)
}

/**
初始化授权列表
*/
func TdataVehicleAuthCheck(vid string) error {
	vehicleAuths := GetVehicleAuths(vid)

	//授权天数，起始时间

	authTime := uint32(time.Now().Unix())

	authDaysSeconds := uint32(AuthDays * 24 * 3600)
	authExpire := authTime + authDaysSeconds

	for _, vehicleId := range vehicleAuths {

		vehicleAuth := &VehicleAuth{
			VehicleId:  vehicleId,
			AuthDays:   AuthDays,
			AuthTime:   authTime,
			AuthExpire: authExpire,
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

func GetVehicleAuths(vid string) []string {
	//添加需要授权的唯一标识码
	vehicleAuths := []string{
		//"754d2728b4e549c5a16c0180fcacb800",
		"derXH5DghbCV3UVHFQaCNbmHitQHcTfj",
		"754d2728b4e549c5a16c0180fcacb80",
	}

	fvehicleAuths := []string{}
	for _, vi := range vehicleAuths {
		vehicleIdMd5 := Md5(vi + PasswordSecret)
		fvehicleAuths = append(fvehicleAuths, vehicleIdMd5)
	}

	if vid != "" {
		fvehicleAuths = append(fvehicleAuths, vid)
	}
	return fvehicleAuths
}

type VehicleAuth struct {
	gorm.Model
	VehicleId  string //关联的小v ID
	AuthTime   uint32
	AuthDays   uint32
	AuthExpire uint32
}

//生成32位md5字串
func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
