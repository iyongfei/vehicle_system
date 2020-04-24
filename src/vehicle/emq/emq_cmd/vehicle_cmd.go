package emq_cmd

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

type VehicleSetCmd struct {
	VehicleId string
	CmdId     int
	TaskType  int

	Type   int
	Switch bool
}

func (setCmd *VehicleSetCmd) CreateVehicleTopicMsg() interface{} {
	publishItem := &protobuf.Command{}

	//ItemType
	publishItem.ItemType = protobuf.Command_TaskType(setCmd.TaskType)
	//param
	gwSetParams := &protobuf.GwSetParam{}
	gwSetParams.Type = protobuf.GwSetParam_Type(setCmd.Type)
	gwSetParams.Switch = setCmd.Switch
	publishItem.Param, _ = proto.Marshal(gwSetParams)
	//CmdID
	var resultcmdItemKey string

	switch setCmd.Type {
	case int(protobuf.GwSetParam_PROTECT):
		taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction := protobuf.GwSetParam_Type_name[int32(protobuf.GwSetParam_PROTECT)]
		resultcmdItemKey = createCmdId(taskTypeName, taskTypeAction, util.RandomString(16))

	case int(protobuf.GwSetParam_RESTART):
		taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction := protobuf.GwSetParam_Type_name[int32(protobuf.GwSetParam_RESTART)]
		resultcmdItemKey = createCmdId(taskTypeName, taskTypeAction, util.RandomString(16))

	case int(protobuf.GwSetParam_DEFAULT):
		taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction := protobuf.GwSetParam_Type_name[int32(protobuf.GwSetParam_DEFAULT)]
		resultcmdItemKey = createCmdId(taskTypeName, taskTypeAction, util.RandomString(16))
	default:
		//类型不对
	}
	publishItem.CmdID = resultcmdItemKey
	resultcmdItemsBys, _ := proto.Marshal(publishItem)

	logger.Logger.Info("%s createVehicleTopicMsg taskType:%s,cmdId:%s,"+
		"type:%s,switch:%+v",
		util.RunFuncName(),
		protobuf.Command_TaskType_name[int32(publishItem.ItemType)],
		publishItem.CmdID,
		protobuf.GwSetParam_Type_name[int32(gwSetParams.Type)],
		gwSetParams.Switch)

	logger.Logger.Print("%s createVehicleTopicMsg taskType:%s,cmdId:%s,"+
		"type:%s,switch:%+v",
		util.RunFuncName(),
		protobuf.Command_TaskType_name[int32(publishItem.ItemType)],
		publishItem.CmdID,
		protobuf.GwSetParam_Type_name[int32(gwSetParams.Type)],
		gwSetParams.Switch)
	return resultcmdItemsBys
}
