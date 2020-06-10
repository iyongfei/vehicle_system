package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

//////////////////////////////////探测类型///////////////////////////
type FprintInfo struct {
	gorm.Model
	FprintInfoId string
	VehicleId    string
	DeviceMac    string
	TradeMark    string
	Os           string
	DstPort      uint32
	ExamineNet   string
	AccessNet    bool
}

//序列化为数字类型
func (temp *FprintInfo) MarshalJSON() ([]byte, error) {
	type tempType FprintInfo
	return json.Marshal(&struct {
		CreatedAt int64
		*tempType
	}{
		CreatedAt: temp.CreatedAt.Unix(),
		tempType:  (*tempType)(temp),
	})
}

func (fprintDetectInfo *FprintInfo) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(fprintDetectInfo, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintDetectInfo *FprintInfo) InsertModel() error {
	return mysql.CreateModel(fprintDetectInfo)
}
func (fprintDetectInfo *FprintInfo) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(fprintDetectInfo, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (fprintDetectInfo *FprintInfo) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(fprintDetectInfo, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (fprintDetectInfo *FprintInfo) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(fprintDetectInfo, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintDetectInfo *FprintInfo) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintDetectInfo *FprintInfo) CreateModel(assetParams ...interface{}) interface{} {
	return fprintDetectInfo
}

/////////////////////////////////主动探测/////////////////////////////////
type FprintInfoActive struct {
	gorm.Model
	FprintInfoId string
	Os           string
}

func (fprintInfoActive *FprintInfoActive) InsertModel() error {
	return mysql.CreateModel(fprintInfoActive)
}
func (fprintInfoActive *FprintInfoActive) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(fprintInfoActive, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (fprintInfoActive *FprintInfoActive) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(fprintInfoActive, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (fprintInfoActive *FprintInfoActive) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(fprintInfoActive, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintInfoActive *FprintInfoActive) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintInfoActive *FprintInfoActive) CreateModel(params ...interface{}) interface{} {
	fprintActiveParams := params[0].(*protobuf.FingerprintParam_ActiveDetect)
	fprintInfoActive.Os = fprintActiveParams.GetOs()

	return fprintInfoActive
}

/////////////////////////////////被动探测/////////////////////////////////

type FprintInfoPassive struct {
	gorm.Model
	FprintInfoId string
	DstPort      uint32
}

func (fprintInfoPassive *FprintInfoPassive) InsertModel() error {
	return mysql.CreateModel(fprintInfoPassive)
}
func (fprintInfoPassive *FprintInfoPassive) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(fprintInfoPassive, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (fprintInfoPassive *FprintInfoPassive) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(fprintInfoPassive, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (fprintInfoPassive *FprintInfoPassive) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(fprintInfoPassive, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintInfoPassive *FprintInfoPassive) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (fprintInfoPassive *FprintInfoPassive) CreateModel(params ...interface{}) interface{} {
	fprintPassiveParams := params[0].(*protobuf.FingerprintParam_PassiveLearn)
	fprintInfoPassive.DstPort = fprintPassiveParams.GetDstPort()

	return fprintInfoPassive
}

//////////////////////////////////////////////////

type FprintPassiveInfo struct {
	gorm.Model

	FprintInfoId string

	Hash         uint32
	SrcIp        string
	SrcPort      uint32
	DstIp        string
	DstPort      uint32
	Protocol     uint8
	FlowInfo     string
	SafeType     uint8
	SafeInfo     string
	StartTime    uint32
	LastSeenTime uint32
	SrcDstBytes  uint64
	DstSrcBytes  uint64
	Stat         uint8
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

func (fprintPassiveInfo *FprintPassiveInfo) CreateModel(params ...interface{}) interface{} {
	fprintPassiveInfoParams := params[0].(*protobuf.PassiveInfoItem)
	fprintPassiveInfo.Hash = fprintPassiveInfoParams.GetHash()

	sipLittleEndian := util.BytesToLittleEndian(util.BigToBytes(fprintPassiveInfoParams.GetSrcIp()))
	fprintPassiveInfo.SrcIp = util.IpIntToString(int(sipLittleEndian))

	fprintPassiveInfo.SrcPort = fprintPassiveInfoParams.GetSrcPort()

	dipLittleEndian := util.BytesToLittleEndian(util.BigToBytes(fprintPassiveInfoParams.GetDstIp()))
	fprintPassiveInfo.DstIp = util.IpIntToString(int(dipLittleEndian))

	fprintPassiveInfo.DstPort = fprintPassiveInfoParams.GetDstPort()
	fprintPassiveInfo.Protocol = uint8(fprintPassiveInfoParams.GetProtocol())
	fprintPassiveInfo.FlowInfo = fprintPassiveInfoParams.GetFlowInfo()
	fprintPassiveInfo.SafeType = uint8(fprintPassiveInfoParams.GetSafeType())
	fprintPassiveInfo.SafeInfo = fprintPassiveInfoParams.GetSafeInfo()
	fprintPassiveInfo.StartTime = fprintPassiveInfoParams.GetStartTime()
	fprintPassiveInfo.LastSeenTime = fprintPassiveInfoParams.GetLastSeenTime()
	fprintPassiveInfo.SrcDstBytes = fprintPassiveInfoParams.GetSrc2DstBytes()
	fprintPassiveInfo.DstSrcBytes = fprintPassiveInfoParams.GetDst2SrcBytes()
	fprintPassiveInfo.Stat = uint8(fprintPassiveInfoParams.GetFlowStat())
	return fprintPassiveInfo
}

//
///////////////////////////////////////////////////
//
//type FprintDetectPassiveInfo struct {
//	gorm.Model
//	DeviceMac string
//	TradeMark string
//	VehicleId string
//	Os        string
//
//	ExamineNet string //入网审批
//	AccessNet  bool   //允许入网
//
//	DstPort uint32
//}
//
//func GetPaginAssetFprints(pageIndex int, pageSize int, totalCount *int, query interface{}, args ...interface{}) ([]*FprintDetectPassiveInfo, error) {
//	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
//	if err != nil {
//		return nil, fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
//	}
//	assetFprints := []*FprintDetectPassiveInfo{}
//
//	err = vgorm.Debug().
//		Table("fprint_detect_infos").
//		Select("fprint_detect_infos.*,fprint_passive_infos.dst_port").
//		Where(query, args...).
//		Order("fprint_detect_infos.created_at desc").
//		Joins("inner join fprint_passive_infos ON fprint_passive_infos.device_mac = fprint_detect_infos.device_mac").
//		Offset((pageIndex - 1) * pageSize).
//		Limit(pageSize).
//		Scan(&assetFprints).
//		Limit(-1).
//		Count(totalCount).
//		Error
//
//	return assetFprints, err
//}
