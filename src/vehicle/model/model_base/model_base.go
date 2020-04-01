package model_base

type ModelBaseImpl interface {
	//插入
	InsertModel(interface{}) error
	//查询
	GetModelsByCondition(model interface{},query interface{}, args ...interface{})(error,bool)
	//修改
	UpdateModelsByCondition(query interface{}, args ...interface{})(error)
	//删除
	DeleModelsByCondition(query interface{}, args ...interface{})(error)
}
