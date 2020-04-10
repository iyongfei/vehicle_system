package model

import "github.com/jinzhu/gorm"

/**管理来源*/
type StudyOrigin struct {
	gorm.Model
	StudyOriginId  string
	Name  string
}