package model_base

type ModelBaseImpl interface {
	//插入
	InsertModel(interface{}) error
	//查询
	GetModelByCondition(model interface{},query interface{}, args ...interface{})(error,bool)
	//查询
	GetModelListByCondition(model interface{},query interface{}, args ...interface{})(error)
	//修改
	UpdateModelsByCondition(query interface{}, args ...interface{})(error)
	//删除
	DeleModelsByCondition(query interface{}, args ...interface{})(error)
}
