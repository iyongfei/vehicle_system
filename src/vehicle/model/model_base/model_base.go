package model_base

import (
	"fmt"
	"time"
)

type ModelBaseImpl interface {
	//插入
	InsertModel() error
	//查询
	GetModelByCondition(query interface{}, args ...interface{})(error,bool)
	//查询
	GetModelListByCondition(model interface{},query interface{}, args ...interface{})(error)
	//修改
	UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{})(error)
	//删除
	DeleModelsByCondition(query interface{}, args ...interface{})(error)

	CreateModel(args ...interface{}) interface{}
}

type ModelBaseImplPagination interface {
	//查询
	GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
		paginModel interface{}, query interface{}, args ...interface{})(error)
}


//插入
type InsertModelImpl interface {
	InsertModel() error
}
//查询
type GetModelImpl interface {
	GetModelByCondition(query interface{}, args ...interface{})(error,bool)
}
//查询
type GetModelListImpl interface {
	GetModelListByCondition(model interface{},query interface{}, args ...interface{})(error)
}

//修改
type UpdateModelImpl interface {
	UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{})(error)
}

//删除
type DeleModelImpl interface {
	DeleModelsByCondition(query interface{}, args ...interface{})(error)
}
//插入
type CreateModelImpl interface {
	CreateModel(args ...interface{}) interface{}
}



type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}




