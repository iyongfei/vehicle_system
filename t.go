package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/util"
)

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

func Str2Stamp(formatTimeStr string) int64 {
	timeStruct := util.Str2Time(formatTimeStr)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

type CreatedAt time.Time

func (ut *CreatedAt) MarshalJSON() (data []byte, err error) {
	t := strconv.FormatInt(time.Time(*ut).Unix(), 10)
	data = []byte(t)
	return
}

type a struct {
	gorm.Model
	CreatedAt CreatedAt
	Name      string
}

type F func(int)

var aaa F

func main() {

	a := []string{"a", "b", "c", "d", "e"}

	fmt.Println(a[0:])
	return
	claims1 := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Second).Unix(), // 过期时间，必须设置
		Issuer:    "wang",                                  // 可不必设置，也可以填充用户名，
	}
	//expired := time.Now().Add(148 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims1) //生成token
	accessToken, _ := token.SignedString([]byte("vector.sign"))
	fmt.Println(accessToken)
	////////////////////////////

	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims2 := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 10,
			Issuer:    "test",
		},
	}

	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	ss, err := token2.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return

	//dip := "3232235898"
	//strDIp, _ := strconv.Atoi(dip)

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

type FlowProtos int32
type FlowSafetype int32
type FlowStat int32
type FlowParam_FItem struct {
	Hash                 uint32       `protobuf:"varint,1,opt,name=hash,proto3" json:"hash,omitempty"`
	SrcIp                uint32       `protobuf:"varint,2,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	SrcPort              uint32       `protobuf:"varint,3,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstIp                uint32       `protobuf:"varint,4,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	DstPort              uint32       `protobuf:"varint,5,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	Protocol             FlowProtos   `protobuf:"varint,6,opt,name=protocol,proto3,enum=protobuf.FlowProtos" json:"protocol,omitempty"`
	FlowInfo             string       `protobuf:"bytes,7,opt,name=flow_info,json=flowInfo,proto3" json:"flow_info,omitempty"`
	SafeType             FlowSafetype `protobuf:"varint,8,opt,name=safe_type,json=safeType,proto3,enum=protobuf.FlowSafetype" json:"safe_type,omitempty"`
	SafeInfo             string       `protobuf:"bytes,9,opt,name=safe_info,json=safeInfo,proto3" json:"safe_info,omitempty"`
	StartTime            uint32       `protobuf:"varint,10,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	LastSeenTime         uint32       `protobuf:"varint,11,opt,name=last_seen_time,json=lastSeenTime,proto3" json:"last_seen_time,omitempty"`
	Src2DstBytes         uint64       `protobuf:"varint,12,opt,name=src2dst_bytes,json=src2dstBytes,proto3" json:"src2dst_bytes,omitempty"`
	Dst2SrcBytes         uint64       `protobuf:"varint,13,opt,name=dst2src_bytes,json=dst2srcBytes,proto3" json:"dst2src_bytes,omitempty"`
	FlowStat             FlowStat     `protobuf:"varint,14,opt,name=flow_stat,json=flowStat,proto3,enum=protobuf.FlowStat" json:"flow_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
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

func InitDataBase() *gorm.DB {
	GormDb, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/vehicle?charset=utf8&parseTime=True&loc=Local")
	//GormDb.LogMode(true)

	GormDb.DB().SetMaxIdleConns(conf.MaxIdleConns)
	GormDb.DB().SetMaxOpenConns(conf.MaxOpenConns)
	return GormDb
}
