package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type AutomatedLearning  struct {
	gorm.Model
	LearningId string
	SampleId string

	FileName string
	Description string
}


/**
学习结果
SampleId,GwId任务采集学习
RuleId自定义规则
ThreatWhiteId事件加白
*/
type AutomatedLearningResult struct {
	gorm.Model
	LearningResultId    string
	OriginId            string
	OriginType          uint8

	//CollectStatus       uint8
	//Name  string
	//Introduce           string
	//GwId       			string
	//
	//
	//RuleId string
	//
	//ThreatWhiteId string//
	//
	//StudyGroupId        string
	//ResultFetchStyle  uint8
	//BeginStudyTime  time.Time
	//StudyRemainTime uint32
}


func (automatedLearningResult *AutomatedLearningResult) InsertModel() error {
	return mysql.CreateModel(automatedLearningResult)
}
func (automatedLearningResult *AutomatedLearningResult) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(automatedLearningResult, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (automatedLearningResult *AutomatedLearningResult) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(automatedLearningResult, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (automatedLearningResult *AutomatedLearningResult) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (automatedLearningResult *AutomatedLearningResult) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (automatedLearningResult *AutomatedLearningResult) CreateModel(strategyParams ...interface{}) interface{} {
	return automatedLearningResult
}
