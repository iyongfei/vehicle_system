package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type Fprint struct {
	gorm.Model

	FprintId  string
	VehicleId string
	AssetId   string

	//Categorys string //流量占比
	//CategorysRate float64

	CollectProtoRate  float64 //流量占比
	CollectProtoFlows string  //流量占比

	CollectBytesRate float64 //采集流量
	CollectBytes     uint64  //采集流量

	CollectHost     string  //hostname
	CollectHostRate float64 //hostname

	CollectTlsRate float64 //tls
	CollectTls     string  //tls

	CollectTimeRate float64 //采集时间
	CollectTime     uint32  //采集时间

	CollectTotalRate float64 //采集时间

	CollectStart  uint64
	CollectFinish bool
	AutoCateId    string
	AutoCateRate  float64 //采集时间
}

//序列化为数字类型
func (tmp *Fprint) MarshalJSON() ([]byte, error) {
	type Type Fprint
	return json.Marshal(&struct {
		CreatedAt int64
		*Type
	}{
		CreatedAt: tmp.CreatedAt.Unix(),
		Type:      (*Type)(tmp),
	})
}

func (f *Fprint) InsertModel() error {
	return mysql.CreateModel(f)
}

func (f *Fprint) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(f, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}

func (f *Fprint) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(f, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (f *Fprint) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(f, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (f *Fprint) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (flow *Fprint) CreateModel(flowParam ...interface{}) interface{} {
	return flow
}

func (flow *Fprint) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(flow, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

type FprintJoinAsset struct {
	gorm.Model
	VehicleId  string
	AssetId    string
	AutoCateId string

	//join category
	AutoCateName string
}

func GetFprintJoinAsset(query string, args ...interface{}) (*FprintJoinAsset, error) {
	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
	if err != nil {
		return nil, fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
	}
	fprintJoinAsset := &FprintJoinAsset{}
	err = vgorm.Debug().
		Table("fprints").
		Select("fprints.*,categories.name as auto_cate_name").
		Where(query, args...).
		Joins("inner JOIN assets ON fprints.asset_id = assets.asset_id").
		Joins("inner JOIN categories ON categories.cate_id = fprints.auto_cate_id").
		Scan(&fprintJoinAsset).
		Error
	return fprintJoinAsset, err
}

func GetFprintJoinAssetList(query string, args ...interface{}) ([]*FprintJoinAsset, error) {
	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
	if err != nil {
		return nil, fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
	}
	fprintJoinAsset := []*FprintJoinAsset{}
	err = vgorm.Debug().
		Table("fprints").
		Select("fprints.*,categories.name as auto_cate_name").
		Where(query, args...).
		Joins("inner JOIN assets ON fprints.asset_id = assets.asset_id").
		Joins("inner JOIN categories ON categories.cate_id = fprints.auto_cate_id").
		Scan(&fprintJoinAsset).
		Error
	return fprintJoinAsset, err
}
