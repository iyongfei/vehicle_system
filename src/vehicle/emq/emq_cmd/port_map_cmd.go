package emq_cmd

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

type PortMapSetCmd struct {
	VehicleId   string
	CmdId int
	TaskType int


	DestPort string
	SrcPort  string
	Switch   bool
	Protocol int
	DestIp   string
}

func (setCmd *PortMapSetCmd) CreatePortMapSetCmdTopicMsg() interface{}{
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
		taskTypeName:=protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction:=protobuf.GwSetParam_Type_name[int32(protobuf.GwSetParam_PROTECT)]
		resultcmdItemKey = createCmdId(taskTypeName,taskTypeAction)

	case int(protobuf.GwSetParam_RESTART):
		taskTypeName:=protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction:=protobuf.GwSetParam_Type_name[int32(protobuf.GwSetParam_RESTART)]
		resultcmdItemKey = createCmdId(taskTypeName,taskTypeAction)

	case int(protobuf.GwSetParam_DEFAULT):
		taskTypeName:=protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction:=protobuf.GwSetParam_Type_name[int32(protobuf.GwSetParam_DEFAULT)]
		resultcmdItemKey = createCmdId(taskTypeName,taskTypeAction)
	default:
		//类型不对
	}
	publishItem.CmdID = resultcmdItemKey
	resultcmdItemsBys, _ := proto.Marshal(publishItem)

	logger.Logger.Info("%s createVehicleTopicMsg publishItem:%+v",util.RunFuncName(),publishItem)
	logger.Logger.Print("%s createVehicleTopicMsg publishItem:%+v",util.RunFuncName(),publishItem)

	return resultcmdItemsBys
}

