package model

import (
	"github.com/jinzhu/gorm"
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
	LearningResultId string
	SampleId            string

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
