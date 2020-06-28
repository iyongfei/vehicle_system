package api_server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/emq_cmd"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/emq/topic_publish_handler"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func EditVehicleInfo(c *gin.Context) {

	vehicleId := c.Param("vehicle_id")
	name := c.PostForm("name")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, name)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
	}

	//查询是否存在
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//更新名字

	//编辑
	attrs := map[string]interface{}{
		"name": name,
	}
	if err := modelBase.UpdateModelsByCondition(attrs, "vehicle_id = ?",
		[]interface{}{vehicleInfo.VehicleId}...); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqUpdateVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	_, _ = modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	responseData := map[string]interface{}{
		"vehicle": vehicleInfo,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetVehicleSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}

func EditVehicle(c *gin.Context) {
	vehicleId := c.Param("vehicle_id")
	setTypeP := c.PostForm("type")
	setSwitchP := c.PostForm("switch")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, setTypeP, setSwitchP)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty", util.RunFuncName())
		logger.Logger.Print("%s argsTrimsEmpty", util.RunFuncName())
	}
	setType, _ := strconv.Atoi(setTypeP)
	setSwitch := true
	switch setSwitchP {
	case "true":
		setSwitch = true
	case "false":
		setSwitch = false
	}

	//查询是否存在
	vehicleInfo := &model.VehicleInfo{}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	//更新
	vehicleCmd := &emq_cmd.VehicleSetCmd{
		VehicleId: vehicleId,
		TaskType:  int(protobuf.Command_GW_SET),

		Type:   setType,
		Switch: setSwitch,
	}

	topic_publish_handler.GetPublishService().PutMsg2PublicChan(vehicleCmd)

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqUpdateVehicleSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}

func GetVehicles(c *gin.Context) {
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")
	startTimeP := c.Query("start_time")
	endTimeP := c.Query("end_time")

	logger.Logger.Info("%s request params page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), pageSizeP, pageIndexP, startTimeP, endTimeP)
	logger.Logger.Print("%s request params page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), pageSizeP, pageIndexP, startTimeP, endTimeP)

	fpageSize, _ := strconv.Atoi(pageSizeP)
	fpageIndex, _ := strconv.Atoi(pageIndexP)

	var fStartTime time.Time
	var fEndTime time.Time

	startTime, _ := strconv.Atoi(startTimeP)
	endTime, _ := strconv.Atoi(endTimeP)

	//默认20
	defaultPageSize := 20
	if fpageSize == 0 {
		fpageSize = defaultPageSize
	}
	//默认第一页
	defaultPageIndex := 1
	if fpageIndex == 0 {
		fpageIndex = defaultPageIndex
	}
	//默认2天前
	defaultStartTime := util.GetFewDayAgo(2) //2
	if startTime == 0 {
		fStartTime = defaultStartTime
	} else {
		fStartTime = util.StampUnix2Time(int64(startTime))
	}

	//默认当前时间
	defaultEndTime := time.Now()
	if endTime == 0 {
		fEndTime = defaultEndTime
	} else {
		fEndTime = util.StampUnix2Time(int64(endTime))
	}

	logger.Logger.Info("%s request params page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), pageSizeP, pageIndexP, startTimeP, endTimeP)
	logger.Logger.Print("%s request params page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), pageSizeP, pageIndexP, startTimeP, endTimeP)

	/////////////////////
	vehicleInfos := []*model.VehicleInfo{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.VehicleInfo{})

	var sqlQuery string
	var sqlArgs []interface{}

	sqlQuery = "vehicle_infos.created_at BETWEEN ? AND ?"
	sqlArgs = append(sqlArgs, fStartTime, fEndTime)

	err := modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
		&total, &vehicleInfos, "vehicle_infos.created_at desc", sqlQuery,
		sqlArgs...)

	////////////threat///////////////
	vehicleIdMaps := []string{}
	for _, vehicleInfo := range vehicleInfos {
		vehicleIdMaps = append(vehicleIdMaps, vehicleInfo.VehicleId)
	}
	type FlowCount struct {
		Count     int
		VehicleId string
	}
	flowCountSql := "SELECT * from (SELECT COUNT(*) as count,vehicle_id FROM flows WHERE vehicle_id IN (?) GROUP BY vehicle_id) as temp"
	FlowCountModelList := []*FlowCount{}

	_ = mysql.QueryRawsqlScanStruct(flowCountSql, vehicleIdMaps, &FlowCountModelList)

	//map
	mapThreatCountModel := map[string]int{}
	for _, threatCountModel := range FlowCountModelList {
		mapThreatCountModel[threatCountModel.VehicleId] = threatCountModel.Count
	}

	var VehicleInfoResponse []*model.VehicleInfoT
	for _, vehicle := range vehicleInfos {
		flowCount := mapThreatCountModel[vehicle.VehicleId]

		vehicleInfoTmp := model.CreateVehicleInfoT(vehicle, flowCount)

		VehicleInfoResponse = append(VehicleInfoResponse, vehicleInfoTmp)
	}

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehiclesFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"vehicles":   VehicleInfoResponse,
		"totalCount": total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetVehiclesSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func GetVehicle(c *gin.Context) {

	vehicleId := c.Param("vehicle_id")
	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicle_id:%s", util.RunFuncName(), vehicleId)
		return
	}
	logger.Logger.Info("%s vehicleId:%s", util.RunFuncName(), vehicleId)
	logger.Logger.Print("%s vehicleId:%s", util.RunFuncName(), vehicleId)

	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicle_id:%s,err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicle_id:%s,recordNotFound", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"vehicle": vehicleInfo,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetVehicleSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

/**
添加
*/

func AddVehicle(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	vehicleInfo := &model.VehicleInfo{}
	err := json.Unmarshal(body, vehicleInfo)

	if err != nil {
		logger.Logger.Error("%s unmarshal vehicle err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Print("%s unmarshal vehicle err:%s", util.RunFuncName(), err.Error())
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleInfo.VehicleId}...)

	if !recordNotFound {
		logger.Logger.Error("%s vehicleId:%s exist", util.RunFuncName(), vehicleInfo.VehicleId)
		logger.Logger.Print("%s vehicleId:%s exist", util.RunFuncName(), vehicleInfo.VehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if err := modelBase.InsertModel(); err != nil {
		logger.Logger.Error("%s add vehicleId:%s err:%s", util.RunFuncName(), vehicleInfo.VehicleId, err.Error())
		logger.Logger.Print("%s add vehicleId:%s err:%s", util.RunFuncName(), vehicleInfo.VehicleId, err.Error())
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"vehicle": vehicleInfo,
	}

	ret := response.StructResponseObj(response.VStatusOK, response.ReqAddVehicleSuccessMsg, responseData)
	c.JSON(http.StatusOK, ret)
}

func DeleVehicle(c *gin.Context) {
	vehicleId := c.Param("vehicle_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s argsTrimsEmpty", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s argsTrimsEmpty", util.RunFuncName(), vehicleId)
		return
	}

	fmt.Println(vehicleId, "jsldfjs")

	vehicleObj := &model.VehicleInfo{
		VehicleId: vehicleId,
	}

	modelBase := model_base.ModelBaseImpl(vehicleObj)
	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", []interface{}{vehicleObj.VehicleId}...)

	if err != nil {
		logger.Logger.Error("%s vehicleId:%s err:%s", util.RunFuncName(), vehicleId, err)
		logger.Logger.Print("%s vehicleId:%s err:%s", util.RunFuncName(), vehicleId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleVehicleFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if recordNotFound {
		logger.Logger.Error("%s vehicleId:%s,record not exist", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicleId:%s,record not exist", util.RunFuncName(), vehicleId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetVehicleUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err := modelBase.DeleModelsByCondition("vehicle_id = ?",
		[]interface{}{vehicleObj.VehicleId}...); err != nil {
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqDeleVehicleSuccessMsg, "")
	c.JSON(http.StatusOK, retObj)
}
