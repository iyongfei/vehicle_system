package model

import "github.com/jinzhu/gorm"

type AssetLeader struct {
	gorm.Model//创建时间在
	LeaderId  string
	Name  string
	Phone string//手机
}


