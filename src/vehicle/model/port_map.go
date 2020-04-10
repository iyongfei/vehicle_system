package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

type PortMap struct {
	gorm.Model
	PortMapId             string
	VehicleId  			  string //关联某设备
	SrcPort               string //外部端口小v
	DstPort               string//内部端口
	DstIp 				  string//内部ip

	Switch                bool//是否开启端口映射
	ProtocolType          uint8//网络协议类型

}


func (p *PortMap) InsertModel() error {
	return mysql.CreateModel(p)
}
func (p *PortMap) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(p, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (p *PortMap) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(p, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (p *PortMap) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (p *PortMap) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	return nil
}

func (p *PortMap) CreateModel(param ...interface{}) interface{} {
	portMap := param[0].(*protobuf.PortRedirectParam_Item)
	p.DstPort = portMap.GetDestPort()
	p.DstIp = portMap.GetDestIp()
	p.Switch = portMap.GetSwitch()
	p.ProtocolType = uint8(portMap.GetProto())
	return p
}
