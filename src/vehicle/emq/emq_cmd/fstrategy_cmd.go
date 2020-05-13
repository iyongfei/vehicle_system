package emq_cmd

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

type FStrategySetCmd struct {
	VehicleId string
	CmdId     int
	TaskType  int

	FstrategyId string
	Type        int
	HandleMode  int
	Enable      bool
	GroupId     string
}

func (setCmd *FStrategySetCmd) CreateFStrategyTopicMsg() interface{} {
	publishItem := &protobuf.Command{}

	//ItemType
	publishItem.ItemType = protobuf.Command_TaskType(setCmd.TaskType)
	//param
	fstrategySetParams := &protobuf.FlowStrategyAddParam{}
	fstrategySetParams.FlowStrategyId = setCmd.FstrategyId
	fstrategySetParams.HandleMode = protobuf.FlowStrategyAddParam_HandleMode(setCmd.HandleMode)
	fstrategySetParams.DefenseType = protobuf.FlowStrategyAddParam_Type(setCmd.Type)
	fstrategySetParams.Enable = setCmd.Enable

	dipPortList, dipPortMap := FetchDipPortList(setCmd)
	fstrategySetParams.DIPList = dipPortList

	publishItem.Param, _ = proto.Marshal(fstrategySetParams)
	//CmdID
	var resultcmdItemKey string
	taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
	resultcmdItemKey = createCmdId(taskTypeName, util.RandomString(16))

	publishItem.CmdID = resultcmdItemKey
	resultcmdItemsBys, _ := proto.Marshal(publishItem)

	dipPortMapMarshal, _ := json.Marshal(dipPortMap)

	logger.Logger.Info("%s createFStrategyTopicMsg taskType:%s,cmdId:%s,"+
		"fstrategy_id:%s,type_name:%s,handle_mode:%s,enable:%v,dipPortList:%s",
		util.RunFuncName(),
		protobuf.Command_TaskType_name[int32(publishItem.ItemType)],
		publishItem.CmdID,
		fstrategySetParams.FlowStrategyId,
		protobuf.FlowStrategyAddParam_Type_name[int32(fstrategySetParams.DefenseType)],
		protobuf.FlowStrategyAddParam_HandleMode_name[int32(fstrategySetParams.HandleMode)],
		fstrategySetParams.Enable, dipPortMapMarshal)

	logger.Logger.Print("%s createFStrategyTopicMsg taskType:%s,cmdId:%s,"+
		"fstrategy_id:%s,type_name:%s,handle_mode:%s,enable:%v,dipPortList:%s",
		util.RunFuncName(),
		protobuf.Command_TaskType_name[int32(publishItem.ItemType)],
		publishItem.CmdID,
		fstrategySetParams.FlowStrategyId,
		protobuf.FlowStrategyAddParam_Type_name[int32(fstrategySetParams.DefenseType)],
		protobuf.FlowStrategyAddParam_HandleMode_name[int32(fstrategySetParams.HandleMode)],
		fstrategySetParams.Enable, dipPortMapMarshal)

	return resultcmdItemsBys
}

func FetchDipPortList(setCmd *FStrategySetCmd) ([]*protobuf.FlowStrategyAddParam_FlowStrategyItem, map[string][]uint32) {
	//获取learning_result_ids
	strategyVehicleLearningResultJoin, _ := model.GetFlowStrategyVehicleItems(
		"fstrategy_vehicles.vehicle_id = ?", []interface{}{setCmd.VehicleId}...)

	fstrategyVehicleItemIds := []string{}
	for _, v := range strategyVehicleLearningResultJoin {
		fstrategyVehicleItemIds = append(fstrategyVehicleItemIds, v.FstrategyItemId)
	}

	//去重learning_result_ids
	fFstrategyVehicleItems := util.RemoveRepeatedForSlice(fstrategyVehicleItemIds)

	fstrategyVehicleItems := []*model.FstrategyItem{}

	_ = model_base.ModelBaseImpl(&model.AutomatedLearningResult{}).
		GetModelListByCondition(&fstrategyVehicleItems,
			"fstrategy_item_id in (?)", []interface{}{fFstrategyVehicleItems}...)

	fProtobufStrategyVehicleItems := []*protobuf.FlowStrategyAddParam_FlowStrategyItem{}
	mapper := map[string][]uint32{}
	for _, fItem := range fstrategyVehicleItems {

		//去重
		dip := fItem.DstIp //string
		//strDip, _ := strconv.Atoi(dstIp)

		//sipBigEndian := util.BytesToBigEndian(util.LittleToBytes(uint32(strDip)))
		//转换
		//
		//dip := strconv.Itoa(int(sipBigEndian))

		dport := fItem.DstPort

		if len(mapper[dip]) == 0 {

			mapper[dip] = []uint32{dport}
		} else {
			if !util.IsExistInSlice(dport, mapper[dip]) {
				mapper[dip] = append(mapper[dip], dport)
			}
		}
	}

	for dip, ports := range mapper {
		for _, port := range ports {
			fProtobufStrategyVehicleItem := &protobuf.FlowStrategyAddParam_FlowStrategyItem{}
			ipI := uint32(util.InetAton(dip))

			sipBigEndian := util.BytesToBigEndian(util.LittleToBytes(uint32(ipI)))
			//转换

			fProtobufStrategyVehicleItem.DstIp = sipBigEndian
			fProtobufStrategyVehicleItem.DstPort = port

			fProtobufStrategyVehicleItems = append(fProtobufStrategyVehicleItems, fProtobufStrategyVehicleItem)
		}
	}

	return fProtobufStrategyVehicleItems, mapper
}
