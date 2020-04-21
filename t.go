package main

import (
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/conf"
)

func main() {
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
