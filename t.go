package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/util"
)

/**
flowItemId 654576970
flowItemIdBefore &{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>} 654576970 754d2728b4e549c5a16c0180fcacb800
 0  0  0 0  0  0 0 0 0 0}
vehicle_system/src/vehicle/service/flow.(*FlowService).SendFlow send flow info map[754d2728b4e549c5a16c0180fcacb800
:[0xc00054eee0]]
TFlow {"754d2728b4e549c5a16c0180fcacb800\n":[{"ID":1070,"CreatedAt":"2020-04-24T12:46:15+08:00","UpdatedAt":"2020-04-24T12:46:47.540655+08:00","DeletedAt":null,"FlowId":1376020295,"VehicleId":"754d2728b4e549c5a16c0180fcacb800\n","Hash":1376020295,"SrcIp":"192.168.1.235","SrcPort":49240,"DstIp":"192.168.1.103","DstPort":445,"Protocol":41,"FlowInfo":"","SafeType":0,"SafeInfo":"","StartTime":1587703561,"LastSeenTime":1587703561,"SrcDstBytes":246,"DstSrcBytes":120,"Stat":2}]}
[GIN] 2020/04/24 - 12:46:47 | 200 |       9.644µs |             ::1 | POST     "/t_flow"
vehicle_system/src/vehicle/service/flow.(*FlowService).SendFlow sendFlow resp json:null, type:map[string]interface {}

(/Users/mac/go/vehicle_system/src/vehicle/db/mysql/query.go:111)
[2020-04-24 12:46:47]  [0.58ms]  SELECT * FROM `flows`  WHERE `flows`.`deleted_at` IS NULL AND ((flow_id = 654576970 and vehicle_id = '754d2728b4e549c5a16c0180fcacb800
')) ORDER BY `flows`.`id` ASC LIMIT 1
[0 rows affected or returned ]
flowItemIdBeforfe &{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>} 654576970 754d2728b4e549c5a16c0180fcacb800
 654576970 192.168.1.103 41520 192.168.1.192 8843 0  0  1587697020 1587703600 247861 83292 1}
recordNotFound true &{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>} 654576970 754d2728b4e549c5a16c0180fcacb800
 654576970 192.168.1.103 41520 192.168.1.192 8843 0  0  1587697020 1587703600 247861 83292 1}

(/Users/mac/go/vehicle_system/src/vehicle/db/mysql/add.go:13)
[2020-04-24 12:46:47]  Error 1062: Duplicate entry '654576970' for key 'flow_id'

(/Users/mac/go/vehicle_system/src/vehicle/db/mysql/add.go:13)
[2020-04-24 12:46:47]  [0.54ms]  INSERT INTO `flows` (`created_at`,`updated_at`,`deleted_at`,`flow_id`,`vehicle_id`,`hash`,`src_ip`,`src_port`,`dst_ip`,`dst_port`,`protocol`,`flow_info`,`safe_type`,`safe_info`,`start_time`,`last_seen_time`,`src_dst_bytes`,`dst_src_bytes`,`stat`) VALUES ('2020-04-24 12:46:47','2020-04-24 12:46:47',NULL,654576970,'754d2728b4e549c5a16c0180fcacb800
',654576970,'192.168.1.103',41520,'192.168.1.192',8843,0,'',0,'',1587697020,1587703600,247861,83292,1)
[0 rows affected or returned ]
vehicle_system/src/vehicle/emq/topic_subscribe_handler.HandleVehicleFlow insert flowParam err:vehicle_system/src/vehicle/db/mysql.CreateModel err Error 1062: Duplicate entry '654576970' for key 'flow_id'
*/

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

func main() {
	trimStr := "sjkdlf" +
		"" +
		""

	strings.Trim(trimStr, "\n")
	fmt.Println(trimStr)

	r := bytesToLittleEndian(UintToBytes(690530496))
	fmt.Println(util.IpIntToString(int(r)))
	//
	//rrr := util.StringIpToInt("255.255.255.255")
	//fmt.Println(rrr)
	//
	//var testInt int32 = 690530496
	//fmt.Printf("%d use big endian: \n", testInt)
	//
	//rr := util.IpIntToString(int(testInt))
	//fmt.Println(rr)
	//
	//var testBytes []byte = make([]byte, 4)
	//binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	//fmt.Println("int32 to bytes:", testBytes)
	//
	//convInt := binary.LittleEndian.Uint32(testBytes)
	//fmt.Printf("bytes to int32: %d\n\n", convInt)
	//
	//r := util.IpIntToString(int(convInt))
	//fmt.Println(r)

	//uint32(binary.BigEndian.Uint32(buf))
	//diportsMap := map[string][]uint32{
	//	"192.167.1.1": []uint32{198, 2, 3},
	//	"b":           []uint32{18989898, 2, 3},
	//}
	//
	//for _, ports := range diportsMap {
	//	for k, dport := range ports {
	//		dporStr := strconv.Itoa(int(dport))
	//		destPortValid := util.VerifyIpPort(dporStr)
	//		if !destPortValid {
	//			ports = append(ports[:k], ports[k+1:]...)
	//		}
	//	}
	//}
	//
	//fmt.Println(diportsMap)
	//
	//a := []string{"a", "b"}
	//fmt.Println(a[:0])

	//vgorm := InitDataBase()

	//strategyVehicleLearningResultJoins := []*model.StrategyVehicleLearningResultJoin{}
	//_ = vgorm.Debug().
	//	Table("strategies").
	//	Select("strategies.*,strategy_vehicles.vehicle_id").
	//	Where("strategy_vehicles.vehicle_id = ?","TDavCZX6IyE3NDa2OVRaMlt92pMOG3Hw").
	//	Joins("inner join strategy_vehicles ON strategies.strategy_id = strategy_vehicles.strategy_id").
	//	Order("strategies.created_at desc").
	//	Scan(&strategyVehicleLearningResultJoins).
	//	Error
	//for k,v:=range strategyVehicleLearningResultJoins{
	//	fmt.Println(k,v.StrategyId)
	//}

	//fdportSlice := []uint32{}
	//ori:="1,3,4,5,5"
	//
	//dstPortSlice := strings.Split(ori, ",")
	//for _, dport := range dstPortSlice {
	//	destPortValid := util.VerifyIpPort(dport)
	//	dpInt,_:=strconv.Atoi(dport)
	//	if destPortValid && !util.IsExistInSlice(dpInt,fdportSlice)  {
	//		fmt.Println(dpInt)
	//		fdportSlice = append(fdportSlice, uint32(dpInt))
	//	}
	//}

	//fstrategyVehicleItems := []*model.FstrategyItem{}
	//mapper := map[string][]uint32{}
	//_ = model_base.ModelBaseImpl(&model.AutomatedLearningResult{}).
	//	GetModelListByCondition(&fstrategyVehicleItems,
	//		"fstrategy_item_id in (?)",
	//		[]interface{}{"Anvl7c2xEdm85wVwstHfNDj6TJeruWpZ","vmtsrkxsI87EoCLOtxag5Dh9V4CkW9GN"}...)

}

func InitDataBase() *gorm.DB {
	GormDb, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/vehicle?charset=utf8&parseTime=True&loc=Local")
	//GormDb.LogMode(true)

	GormDb.DB().SetMaxIdleConns(conf.MaxIdleConns)
	GormDb.DB().SetMaxOpenConns(conf.MaxOpenConns)
	return GormDb
}
