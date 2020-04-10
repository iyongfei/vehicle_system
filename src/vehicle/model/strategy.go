package model

import "github.com/jinzhu/gorm"

type Strategy struct {
	gorm.Model
	StrategyId string

	Type uint8 //策略模式
	HandleMode uint8 //处理方式
	Enable  bool//策略启用状态

	Name string//策略名称
	Introduce string//策略说明
}


type StrategyGroup struct {
	gorm.Model
	StrategyId string
	GroupId string //终端分组
}



type GroupLearningResult struct {
	gorm.Model
	GroupId string
	LearningResultId string//id
}
