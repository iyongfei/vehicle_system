package mysql

import (
	"fmt"
	"vehicle_system/src/vehicle/util"
)

func CreateModel(model interface{}) error  {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	if err = vgorm.Create(model).Error;err!= nil{
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}

	return nil
}


func CreatTable(model interface{})error{
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}

	tExist := vgorm.HasTable(model)

	if !tExist{
		if err := vgorm.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(model).Error;err!= nil{
			return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
		}
	}
	return nil
}

