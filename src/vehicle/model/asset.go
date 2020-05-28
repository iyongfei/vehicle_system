package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/mac"
	"vehicle_system/src/vehicle/util"
)

type Asset struct {
	gorm.Model
	VehicleId string //关联的小v ID
	AssetId   string `gorm:"unique"` //资产id

	IP        string //小v资产 IP
	Mac       string //资产Mac地址
	Name      string //小v资产名称
	TradeMark string //资产品牌

	OnlineStatus bool   //在线状态
	LastOnline   uint32 //最近活跃时间

	InternetSwitch bool //是否允许联网
	ProtectStatus  bool //是否受小V保护
	LanVisitSwitch bool //是否可以访问内网

	AssetGroup  string
	AssetLeader string
}

//序列化为数字类型
func (asset *Asset) MarshalJSON() ([]byte, error) {
	type AssetType Asset
	return json.Marshal(&struct {
		CreatedAt int64
		*AssetType
	}{
		CreatedAt: asset.CreatedAt.Unix(),
		AssetType: (*AssetType)(asset),
	})
}

func (asset *Asset) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(asset, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (asset *Asset) InsertModel() error {
	return mysql.CreateModel(asset)
}
func (asset *Asset) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(asset, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (asset *Asset) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(asset, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (asset *Asset) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(asset, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (asset *Asset) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (asset *Asset) CreateModel(assetParams ...interface{}) interface{} {
	assetParam := assetParams[0].(*protobuf.DeviceParam_Item)
	asset.IP = assetParam.GetIp()
	asset.Mac = assetParam.GetMac()
	asset.Name = assetParam.GetName()
	asset.TradeMark = assetParam.GetTrademark()
	asset.TradeMark = mac.GetOrgByMAC(assetParam.GetMac())
	asset.OnlineStatus = assetParam.GetIsOnline()
	asset.LastOnline = assetParam.GetLastOnline()
	asset.InternetSwitch = assetParam.GetInternetSwitch()
	asset.ProtectStatus = assetParam.GetProtectSwitch()
	asset.LanVisitSwitch = assetParam.GetLanVisitSwitch()
	return asset
}
