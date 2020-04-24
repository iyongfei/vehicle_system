package mysql

import (
	"fmt"
	"vehicle_system/src/vehicle/util"
)

func CreateModel(model interface{}) error {
	vgorm, err := GetMysqlInstance().GetMysqlDB()
	if err != nil {
		return fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
	}
	if err = vgorm.Debug().Create(model).Error; err != nil {
		return fmt.Errorf("%s err %v", util.RunFuncName(), err.Error())
	}

	return nil
}

func CreatTable(model interface{}) (error, int64) {
	vgorm, err := GetMysqlInstance().GetMysqlDB()
	if err != nil {
		return fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error()), 0
	}

	tExist := vgorm.HasTable(model)

	if !tExist {
		db := vgorm.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(model)

		rowsAffected := db.RowsAffected
		err := db.Error

		if err != nil {
			return fmt.Errorf("%s err %v", util.RunFuncName(), err.Error()), 0
		}
		return nil, rowsAffected
	}

	return nil, 0
}
