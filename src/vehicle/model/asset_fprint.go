package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type AssetFprint struct {
	gorm.Model

	AssetId string

	AssetFprintId string
	CateId        string
}

//序列化为数字类型
func (tmp *AssetFprint) MarshalJSON() ([]byte, error) {
	type Type AssetFprint
	return json.Marshal(&struct {
		CreatedAt int64
		*Type
	}{
		CreatedAt: tmp.CreatedAt.Unix(),
		Type:      (*Type)(tmp),
	})
}

func (f *AssetFprint) InsertModel() error {
	return mysql.CreateModel(f)
}

func (f *AssetFprint) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(f, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}

func (f *AssetFprint) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(f, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (f *AssetFprint) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(f, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (f *AssetFprint) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (flow *AssetFprint) CreateModel(flowParam ...interface{}) interface{} {
	return flow
}

func (flow *AssetFprint) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(flow, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

type AssetFprintCateJoin struct {
	gorm.Model
	AssetId       string
	AssetFprintId string
	CateId        string

	CateName string
}

func GetAssetFprintCateJoin(query string, args ...interface{}) (*AssetFprintCateJoin, error) {
	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
	if err != nil {
		return nil, fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
	}
	assetFprintCateJoin := &AssetFprintCateJoin{}
	err = vgorm.Debug().
		Table("asset_fprints").
		Select("asset_fprints.*,categories.name as cate_name").
		Where(query, args...).
		Joins("inner JOIN categories ON categories.cate_id = asset_fprints.cate_id").
		Scan(&assetFprintCateJoin).
		Error
	return assetFprintCateJoin, err
}

func GetAssetFprintCateListJoin(query string, args ...interface{}) ([]*AssetFprintCateJoin, error) {
	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
	if err != nil {
		return nil, fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
	}
	assetFprintCateListJoin := []*AssetFprintCateJoin{}
	err = vgorm.Debug().
		Table("asset_fprints").
		Select("asset_fprints.*,categories.name as cate_name").
		Where(query, args...).
		Joins("inner JOIN categories ON categories.cate_id = asset_fprints.cate_id").
		Scan(&assetFprintCateListJoin).
		Error
	return assetFprintCateListJoin, err
}
