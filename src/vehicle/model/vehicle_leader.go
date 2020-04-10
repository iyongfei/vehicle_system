package model

import "github.com/jinzhu/gorm"

type VehicleLeader struct {
	gorm.Model//创建时间在
	LeaderId  string
	Name  string

	Phone string//手机
}


