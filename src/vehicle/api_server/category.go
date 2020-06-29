package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

/**
添加指纹库
*/

func AddCategory(c *gin.Context) {
	cateName := c.PostForm("cate_name")

	argsTrimsEmpty := util.RrgsTrimsEmpty(cateName)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s category_name:%s", util.RunFuncName(), cateName)
		logger.Logger.Error("%s category_name:%s", util.RunFuncName(), cateName)
		return
	}

	logger.Logger.Print("%s category_name:%s", util.RunFuncName(), cateName)
	logger.Logger.Info("%s category_name:%s", util.RunFuncName(), cateName)

	cate := &model.Category{
		CateId: util.RandomString(32),
		Name:   cateName,
	}
	cateModelBase := model_base.ModelBaseImpl(cate)

	err, cateRecordNotFound := cateModelBase.GetModelByCondition("name = ?", []interface{}{cate.Name}...)
	if !cateRecordNotFound {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if err = cateModelBase.InsertModel(); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddCategoryFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseContent := map[string]interface{}{}
	responseContent["category"] = cate

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddCategorySuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}

/**
查询所有指纹库
*/

func GetCategorys(c *gin.Context) {

	cateModelBase := model_base.ModelBaseImpl(&model.Category{})
	cates := []*model.Category{}
	err := cateModelBase.GetModelListByCondition(&cates, "", []interface{}{}...)
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryListFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseContent := map[string]interface{}{}
	responseContent["categorys"] = cates

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqCategoryListSuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}

/**
删除指纹库
*/

func DeleCategory(c *gin.Context) {
	cateId := c.Param("cate_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(cateId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s cateId:%s", util.RunFuncName(), cateId)
		logger.Logger.Error("%s cateId:%s", util.RunFuncName(), cateId)
		return
	}

	logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId)
	logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId)

	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleCategoryFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	tx := vgorm.Begin()

	//dele 删除FstrategyVehicleItem表
	category := &model.Category{}

	if err := tx.Debug().Unscoped().Where("cate_id = ?", cateId).Delete(category).Error; err != nil {
		tx.Rollback()
		logger.Logger.Error("%s dele category id:%s, err:%s", util.RunFuncName(), cateId, err)
		logger.Logger.Print("%s dele category id:%s, err:%s", util.RunFuncName(), cateId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleCategoryFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	tx.Commit()

	logger.Logger.Print("%s cateId:%s", util.RunFuncName(), cateId)
	logger.Logger.Info("%s cateId:%s", util.RunFuncName(), cateId)
	ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqDeleCategorySuccessMsg, "")
	c.JSON(http.StatusOK, ret)
}

/**
编辑指纹库
*/

func EditCategory(c *gin.Context) {
	cateId := c.Param("cate_id")
	cateName := c.PostForm("cate_name")

	argsTrimsEmpty := util.RrgsTrimsEmpty(cateId, cateName)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)

		logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
		logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
		return
	}

	logger.Logger.Print("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)
	logger.Logger.Error("%s cateId:%s,cateName%s", util.RunFuncName(), cateId, cateName)

	cate := &model.Category{
		CateId: cateId,
	}
	cateModelBase := model_base.ModelBaseImpl(cate)

	err, cateRecordNotFound := cateModelBase.GetModelByCondition("cate_id = ?", []interface{}{cate.CateId}...)
	if cateRecordNotFound {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryNotExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqCategoryFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//编辑
	attrs := map[string]interface{}{
		"name": cateName,
	}
	if err := cateModelBase.UpdateModelsByCondition(attrs, "cate_id = ?",
		[]interface{}{cate.CateId}...); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateCategoryFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//ReqUpdateCategorySuccessMsg
	cateUpdated := &model.Category{
		CateId: cateId,
	}
	cateUpdatedModelBase := model_base.ModelBaseImpl(cateUpdated)
	_, _ = cateUpdatedModelBase.GetModelByCondition("cate_id = ?", []interface{}{cateUpdated.CateId}...)

	responseContent := map[string]interface{}{}
	responseContent["category"] = cateUpdated

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateCategorySuccessMsg, responseContent)
	c.JSON(http.StatusOK, retObj)
}
