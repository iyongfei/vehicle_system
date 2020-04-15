package model_base

import (
	"strconv"
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

func (ut *UnixTime) MarshalJSON() (data []byte, err error) {
	t := strconv.FormatInt(time.Time(*ut).Unix(), 10)
	data = []byte(t)
	return
}

func (ut *UnixTime) UnmarshalJSON(data []byte) (err error) {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return
	}
	t := time.Unix(i, 0)
	*ut = UnixTime(t)
	return
}


//或者如下的方法
//func (u *MyUser) MarshalJSON() ([]byte, error) {
//	type Alias MyUser
//	return json.Marshal(&struct {
//		LastSeen int64 `json:"lastSeen"`
//		*Alias
//	}{
//		LastSeen: u.LastSeen.Unix(),
//		Alias:    (*Alias)(u),
//	})
//}
//
//func (u *MyUser) UnmarshalJSON(data []byte) error {
//	type Alias MyUser
//	aux := &struct {
//		LastSeen int64 `json:"lastSeen"`
//		*Alias
//	}{
//		Alias: (*Alias)(u),
//	}
//	if err := json.Unmarshal(data, &aux); err != nil {
//		return err
//	}
//	u.LastSeen = time.Unix(aux.LastSeen, 0)
//	return nil
//}



