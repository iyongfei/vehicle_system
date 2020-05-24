package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type FingerPrint struct {
	gorm.Model
	FprintId    string
	CateId      string
	VehicleId   string
	DeviceMac   string
	FlowIds     string
	ProtoRate   string
	CollectType uint8
}

func (fingerPrint *FingerPrint) InsertModel() error {
	return mysql.CreateModel(fingerPrint)
}
func (fingerPrint *FingerPrint) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(fingerPrint, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (fingerPrint *FingerPrint) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(fingerPrint, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (fingerPrint *FingerPrint) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(fingerPrint, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fingerPrint *FingerPrint) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fingerPrint *FingerPrint) CreateModel(assetParams ...interface{}) interface{} {
	return fingerPrint
}

////////////////////////////////探测类型///////////////////////////
type FprintDetectInfo struct {
	gorm.Model
	DetectInfoId string
	DeviceMac    string
	TradeMark    string
	VehicleId    string
	Os           string
}

func (fprintDetectInfo *FprintDetectInfo) InsertModel() error {
	return mysql.CreateModel(fprintDetectInfo)
}
func (fprintDetectInfo *FprintDetectInfo) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(fprintDetectInfo, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (fprintDetectInfo *FprintDetectInfo) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(fprintDetectInfo, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (fprintDetectInfo *FprintDetectInfo) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(fprintDetectInfo, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintDetectInfo *FprintDetectInfo) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintDetectInfo *FprintDetectInfo) CreateModel(assetParams ...interface{}) interface{} {
	return fprintDetectInfo
}

///////////////////////////////被动探测/////////////////////////////////

type FprintPassiveInfo struct {
	gorm.Model
	PassiveInfoId string
	DeviceMac     string
	TradeMark     string
	VehicleId     string
	DstPort       uint32
}

func (fprintPassiveInfo *FprintPassiveInfo) InsertModel() error {
	return mysql.CreateModel(fprintPassiveInfo)
}
func (fprintPassiveInfo *FprintPassiveInfo) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(fprintPassiveInfo, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (fprintPassiveInfo *FprintPassiveInfo) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(fprintPassiveInfo, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (fprintPassiveInfo *FprintPassiveInfo) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(fprintPassiveInfo, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintPassiveInfo *FprintPassiveInfo) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintPassiveInfo *FprintPassiveInfo) CreateModel(assetParams ...interface{}) interface{} {
	return fprintPassiveInfo
}

/////////////////////////////////////////////////

type FprintDetectPassiveInfo struct {
	gorm.Model
	DeviceMac string
	TradeMark string
	VehicleId string
	Os        string
	DstPort   uint32
}

func GetPaginAssetFprints(pageIndex int, pageSize int, totalCount *int, query interface{}, args ...interface{}) ([]*FprintDetectPassiveInfo, error) {
	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
	if err != nil {
		return nil, fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
	}
	assetFprints := []*FprintDetectPassiveInfo{}

	err = vgorm.Debug().
		Table("fprint_detect_infos").
		Select("fprint_detect_infos.*,fprint_passive_infos.dst_port").
		Where(query, args...).
		Order("fprint_detect_infos.created_at desc").
		Joins("inner join fprint_passive_infos ON fprint_passive_infos.device_mac = fprint_detect_infos.device_mac").
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Scan(&assetFprints).
		Limit(-1).
		Count(totalCount).
		Error

	return assetFprints, err
}
