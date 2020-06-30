package main

import (
	"fmt"
	"vehicle_system/src/vehicle/util"
)

type AuthVehicle struct {
	Start     uint32
	End       uint32
	VehicleId string
}

type AuthVehicleList struct {
	Server       string
	AuthVehicles []AuthVehicle
}

var aggg []*AuthVehicleList

func main() {
	return

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

//r := bytesToLittleEndian(UintToBytes(690530496))
//fmt.Println(util.IpIntToString(int(r)))
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
