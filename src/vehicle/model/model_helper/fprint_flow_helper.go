package model_helper

import (
	"fmt"
	"vehicle_system/src/vehicle/conf"
)

/**
collect_time=300
proto_count=5
collect_total=1048576
是否采集到了hostname
是否采集到了tls_client_info
*/
/**
判断某个设备采集时长是否达标
*/
func JudgeAssetCollectTime(assetId string) {

	ctime := conf.CollectTime
	fmt.Println(ctime, "ctime")

}
