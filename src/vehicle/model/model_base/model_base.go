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

type ModelBaseImplPagination interface {
	//查询
	GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
		paginModel interface{}, query interface{}, args ...interface{})(error)
}
//
//
//
//func PaginationT(whichPageNum int, perPageNum int, count *int, totalModel interface{},
//	paginModel interface{}, query interface{}, args interface{}) error {
//	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
//	if err != nil {
//		return errors.Wrapf(err, "Pagination open gorm error %v", err)
//	}
//
//	//这里判断的不严谨
//	if whichPageNum < 1 && perPageNum < 1 {
//		return err
//	}
//
//	err = vgorm.Offset((whichPageNum-1)*perPageNum).Limit(perPageNum).Where(query, args).Find(paginModel).
//		Offset(-1).Find(totalModel).Count(count).Error
//
//	return err
//}

//db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)















