package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
)

/**
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build

*/
const VPasswordSecret = "vgw-1214-pwd-key-vgw-1214-pwd-key"

func main() {
	var extraVehicleAuth string
	flag.StringVar(&extraVehicleAuth, "v", "", "vid")
	flag.Parse()

	if extraVehicleAuth != "" {
		vehicleIdMd5 := VMd5(extraVehicleAuth + VPasswordSecret)

		fmt.Println("vmd5->", vehicleIdMd5)
	}

}

//生成32位md5字串
func VMd5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
