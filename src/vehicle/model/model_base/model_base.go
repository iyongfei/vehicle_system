package model_base

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
