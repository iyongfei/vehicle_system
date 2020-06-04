package api_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

/*
获取一条消息流
*/
func GetFlow(c *gin.Context) {
	flowId := c.Param("flow_id")
	vehicleId := c.Query("vehicle_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, flowId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty", util.RunFuncName(), flowId, vehicleId)
		logger.Logger.Print("%s argsTrimsEmpty flowId:%s,vehicleId:%s argsTrimsEmpty", util.RunFuncName(), flowId, vehicleId)
		return
	}
	flowObj := &model.Flow{}

	modelBase := model_base.ModelBaseImpl(flowObj)

	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ? and flow_id = ?", []interface{}{vehicleId, flowId}...)

	if err != nil {
		logger.Logger.Error("%s flowId:%s,err:%s", util.RunFuncName(), flowId, err)
		logger.Logger.Print("%s flowId:%s,err:%s", util.RunFuncName(), flowId, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if recordNotFound {
		logger.Logger.Error("%s flowId:%s,recordNotFound", util.RunFuncName(), flowId)
		logger.Logger.Print("%s flowId:%s,recordNotFound", util.RunFuncName(), flowId)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFlowUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFlowSuccessMsg, flowObj)
	c.JSON(http.StatusOK, retObj)
}

func GetFlowsDps(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	startTimeP := c.Query("start_time")
	endTimeP := c.Query("end_time")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s vehicleId:%s argsTrimsEmpty", util.RunFuncName(), vehicleId)
		logger.Logger.Print("%s vehicleId:%s argsTrimsEmpty", util.RunFuncName(), vehicleId)
		return
	}
	logger.Logger.Info("%s vehicleId:%s,startTimeP:%s,endTimeP:%s", util.RunFuncName(), vehicleId, startTimeP, endTimeP)
	logger.Logger.Print("%s vehicleId:%s", util.RunFuncName(), vehicleId, startTimeP, endTimeP)

	var fStartTime time.Time
	var fEndTime time.Time

	startTime, _ := strconv.Atoi(startTimeP)
	endTime, _ := strconv.Atoi(endTimeP)

	if startTime == 0 {
		fStartTime = util.GetFewDayAgo(2)
	} else {
		fStartTime = util.StampUnix2Time(int64(startTime))
	}

	if endTime == 0 {
		fEndTime = time.Now()
	} else {
		fEndTime = util.StampUnix2Time(int64(endTime))
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

	flows := []*model.Flow{}

	flowModelBase := model_base.ModelBaseImpl(&model.Flow{})

	err = flowModelBase.GetModelListByCondition(&flows, "vehicle_id = ? and flows.created_at BETWEEN ? AND ?",
		[]interface{}{vehicleId, fStartTime, fEndTime}...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	dps := map[string][]uint32{}
	for _, tflow := range flows {
		dip := tflow.DstIp
		dport := tflow.DstPort
		if util.RrgsTrimEmpty(dip) || dport <= 0 {
			continue
		}
		if keyValue, ok := dps[dip]; ok {
			exist := util.IsExistInSlice(dport, keyValue)
			if !exist {
				dps[dip] = append(dps[dip], dport)
			}
		} else {
			dps[dip] = []uint32{dport}
		}
	}

	responseData := &model.TempFlowDp{
		VehicleId: vehicleId,
		Dps:       dps,
	}
	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFlowSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)

}

func GetFlowTypeCounts(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")

	logger.Logger.Info("%s request params vehicle_id:%s", util.RunFuncName(), vehicleId)
	logger.Logger.Print("%s request params vehicle_id:%s", util.RunFuncName(), vehicleId)

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
		return
	}

	type FlowSafeType struct {
		SafeType    int
		Count       int
		SafeTypeStr string
	}

	flowSafeTypeModelList := []*FlowSafeType{}

	execSql := "SELECT COUNT(safe_type) as count,safe_type FROM flows WHERE vehicle_id = ? GROUP BY safe_type;"
	err := mysql.QueryRawsqlScanStruct(execSql, vehicleId, &flowSafeTypeModelList)
	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	var FlowSafetype_name = map[int]string{
		0: "FS_UNSET",
		1: "FS_SAFE",
		2: "FS_VERY_DANGER",
		3: "FS_EMAL_LEAKAGE",
		4: "FS_LEAKAGE",
		5: "FS_WARNING",
	}

	totalCount := 0
	flowMap := map[string]int{}
	for _, flowSafeType := range flowSafeTypeModelList {

		ftype := flowSafeType.SafeType
		SafeTypeStr := FlowSafetype_name[ftype]
		fcount := flowSafeType.Count
		flowMap[SafeTypeStr] = fcount
		totalCount += fcount
	}
	responseData := map[string]interface{}{
		"flow_type_counts": flowMap,
		"total_count":      totalCount,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFlowSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

/*
获取所有消息会话
GetModelPaginationByCondition
*/
func GetPaginationFlows(c *gin.Context) {
	vehicleId := c.Query("vehicle_id")
	pageSizeP := c.Query("page_size")
	pageIndexP := c.Query("page_index")
	startTimeP := c.Query("start_time")
	endTimeP := c.Query("end_time")

	logger.Logger.Info("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)
	logger.Logger.Print("%s request params vehicle_id:%s,page_size:%s,page_index:%s,start_time%s,endtime%s",
		util.RunFuncName(), vehicleId, pageSizeP, pageIndexP, startTimeP, endTimeP)

	vehicleId = util.RrgsTrim(vehicleId)

	//argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId)
	//if argsTrimsEmpty {
	//	ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
	//	c.JSON(http.StatusOK, ret)
	//	logger.Logger.Error("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
	//	logger.Logger.Print("%s argsTrimsEmpty threatId:%s", util.RunFuncName(), argsTrimsEmpty)
	//	return
	//}

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
	//defaultStartTime := util.GetFewDayAgo(2) //2
	if startTime == 0 {
		fStartTime = util.StampUnix2Time(int64(0))
		//fStartTime = defaultStartTime
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

	logger.Logger.Info("%s frequest params vehicle_id:%s,fpageSize:%s,fpageIndex:%s,fStartTime%s,fEndTime%s",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)
	logger.Logger.Print("%s frequest params vehicle_id:%s,fpageSize:%s,fpageIndex:%s,fStartTime%s,fEndTime%s",
		util.RunFuncName(), vehicleId, fpageSize, fpageIndex, fStartTime, fEndTime)

	flows := []*model.Flow{}
	var total int

	modelBase := model_base.ModelBaseImplPagination(&model.Flow{})

	var sqlQuery string
	var sqlArgs []interface{}
	if vehicleId == "" {
		sqlQuery = "flows.created_at BETWEEN ? AND ?"
		sqlArgs = append(sqlArgs, fStartTime, fEndTime)
	} else {
		sqlQuery = "vehicle_id = ? and flows.created_at BETWEEN ? AND ?"
		sqlArgs = append(sqlArgs, vehicleId, fStartTime, fEndTime)
	}

	err := modelBase.GetModelPaginationByCondition(fpageIndex, fpageSize,
		&total, &flows, "flows.created_at desc", sqlQuery,
		sqlArgs...)

	if err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	responseData := map[string]interface{}{
		"flows":       flows,
		"total_count": total,
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetFlowSuccessMsg, responseData)
	c.JSON(http.StatusOK, retObj)
}

func AddFlow(c *gin.Context) {
	vehicleId := c.PostForm("vehicle_id")
	hashP := c.PostForm("hash")
	srcIpArp := c.PostForm("src_ip")
	dstIpArg := c.PostForm("dst_ip")

	logger.Logger.Print("%s vehicleId:%s,hash:%s,srcIpArp:%s,dstIpArg:%s", util.RunFuncName(), vehicleId, hashP, srcIpArp, dstIpArg)

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, hashP, srcIpArp, dstIpArg)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty", util.RunFuncName(), vehicleId, hashP)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty", util.RunFuncName(), vehicleId, hashP)
	}

	hash, _ := strconv.Atoi(hashP)

	flowObj := &model.Flow{
		VehicleId: vehicleId,
		Hash:      uint32(hash),
		SrcIp:     srcIpArp,
		DstIp:     dstIpArg,
		FlowId:    uint32(hash),
	}
	modelBase := model_base.ModelBaseImpl(flowObj)
	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ? and hash = ?", []interface{}{vehicleId, hash}...)

	if err != nil {
		logger.Logger.Error("%s vehicleId:%s,hash:%d,get flow info err:%s", util.RunFuncName(), vehicleId, hash, err)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,get flow info err:%s", util.RunFuncName(), vehicleId, hash, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if !recordNotFound {
		logger.Logger.Error("%s vehicleId:%s,hash:%d,record exist", util.RunFuncName(), vehicleId, hash)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,record exist", util.RunFuncName(), vehicleId, hash)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFlowExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	if err := modelBase.InsertModel(); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqAddFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	retObj := response.StructResponseObj(response.VStatusOK, response.ReqAddFlowSuccessMsg, flowObj)
	c.JSON(http.StatusOK, retObj)
}

func EditFlow(c *gin.Context) {
	hashP := c.Param("flow_id")
	vehicleId := c.PostForm("vehicle_id")
	srcIpArp := c.PostForm("src_ip")
	dstIpArg := c.PostForm("dst_ip")

	logger.Logger.Print("%s vehicleId:%s,hash:%s,srcIpArp:%s,dstIpArg:%s", util.RunFuncName(), vehicleId, hashP, srcIpArp, dstIpArg)

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, hashP, srcIpArp, dstIpArg)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty", util.RunFuncName(), vehicleId, hashP)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty", util.RunFuncName(), vehicleId, hashP)
	}

	hash, _ := strconv.Atoi(hashP)

	flowObj := &model.Flow{
		VehicleId: vehicleId,
		Hash:      uint32(hash),
	}

	modelBase := model_base.ModelBaseImpl(flowObj)
	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ? and hash = ?", []interface{}{vehicleId, hash}...)

	if err != nil {
		logger.Logger.Error("%s vehicleId:%s,hash:%d,get flow info err:%s", util.RunFuncName(), vehicleId, hash, err)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,get flow info err:%s", util.RunFuncName(), vehicleId, hash, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqEditFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if recordNotFound {
		logger.Logger.Error("%s vehicleId:%s,hash:%d,record not exist", util.RunFuncName(), vehicleId, hash)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,record not exist", util.RunFuncName(), vehicleId, hash)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFlowUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	//赋值
	flowObj.SrcIp = srcIpArp
	flowObj.DstIp = dstIpArg

	attrs := map[string]interface{}{
		"src_ip": flowObj.SrcIp,
		"dst_ip": flowObj.DstIp,
	}
	if err := modelBase.UpdateModelsByCondition(attrs, "vehicle_id = ? and hash = ?",
		[]interface{}{flowObj.VehicleId, flowObj.Hash}...); err != nil {
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqEditFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqEditFlowSuccessMsg, flowObj)
	c.JSON(http.StatusOK, retObj)
}

func DeleFlow(c *gin.Context) {
	hashP := c.Param("flow_id")
	vehicleId := c.Query("vehicle_id")

	argsTrimsEmpty := util.RrgsTrimsEmpty(vehicleId, hashP)
	if argsTrimsEmpty {
		ret := response.StructResponseObj(response.VStatusBadRequest, response.ReqArgsIllegalMsg, "")
		c.JSON(http.StatusOK, ret)
		logger.Logger.Error("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty", util.RunFuncName(), vehicleId, hashP)
		logger.Logger.Print("%s argsTrimsEmpty vehicleId:%s,hash:%s argsTrimsEmpty", util.RunFuncName(), vehicleId, hashP)
		return
	}

	hash, _ := strconv.Atoi(hashP)

	flowObj := &model.Flow{
		VehicleId: vehicleId,
		Hash:      uint32(hash),
	}

	modelBase := model_base.ModelBaseImpl(flowObj)
	err, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ? and hash = ?", []interface{}{vehicleId, hash}...)

	if err != nil {
		logger.Logger.Error("%s vehicleId:%s,hash:%d,get flow info err:%s", util.RunFuncName(), vehicleId, hash, err)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,get flow info err:%s", util.RunFuncName(), vehicleId, hash, err)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqDeleFlowFailMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if recordNotFound {
		logger.Logger.Error("%s vehicleId:%s,hash:%d,record not exist", util.RunFuncName(), vehicleId, hash)
		logger.Logger.Print("%s vehicleId:%s,hash:%d,record not exist", util.RunFuncName(), vehicleId, hash)
		ret := response.StructResponseObj(response.VStatusServerError, response.ReqGetFlowUnExistMsg, "")
		c.JSON(http.StatusOK, ret)
		return
	}
	if err := modelBase.DeleModelsByCondition("vehicle_id = ? and hash = ?",
		[]interface{}{flowObj.VehicleId, flowObj.Hash}...); err != nil {
	}

	retObj := response.StructResponseObj(response.VStatusOK, response.ReqGetWhiteListSuccessMsg, flowObj)
	c.JSON(http.StatusOK, retObj)
}
