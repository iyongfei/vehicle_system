package emq_cmd

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

type AssetSetCmd struct {
	VehicleId string
	CmdId     int
	TaskType  int

	Mac    string
	Type   int
	Switch bool
}

func (setCmd *AssetSetCmd) CreateAssetTopicMsg() interface{} {
	publishItem := &protobuf.Command{}

	//ItemType
	publishItem.ItemType = protobuf.Command_TaskType(setCmd.TaskType)
	//param
	assetSetParams := &protobuf.DeviceSetParam{}
	assetSetParams.DeviceMac = setCmd.Mac
	assetSetParams.Type = protobuf.DeviceSetParam_Type(setCmd.Type)
	assetSetParams.Switch = setCmd.Switch
	publishItem.Param, _ = proto.Marshal(assetSetParams)
	//CmdID
	var resultcmdItemKey string

	switch setCmd.Type {
	case int(protobuf.DeviceSetParam_DEFAULT):

	case int(protobuf.DeviceSetParam_PROTECT):
		taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction := protobuf.GwSetParam_Type_name[int32(protobuf.DeviceSetParam_PROTECT)]
		resultcmdItemKey = createCmdId(taskTypeName, taskTypeAction, util.RandomString(16))

	case int(protobuf.DeviceSetParam_INTERNET):
		taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction := protobuf.GwSetParam_Type_name[int32(protobuf.DeviceSetParam_INTERNET)]
		resultcmdItemKey = createCmdId(taskTypeName, taskTypeAction, util.RandomString(16))

	case int(protobuf.DeviceSetParam_GUEST_ACCESS_DEVICE):
		taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction := protobuf.GwSetParam_Type_name[int32(protobuf.DeviceSetParam_GUEST_ACCESS_DEVICE)]
		resultcmdItemKey = createCmdId(taskTypeName, taskTypeAction, util.RandomString(16))

	case int(protobuf.DeviceSetParam_LANVISIT):
		taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
		taskTypeAction := protobuf.GwSetParam_Type_name[int32(protobuf.DeviceSetParam_LANVISIT)]
		resultcmdItemKey = createCmdId(taskTypeName, taskTypeAction, util.RandomString(16))
	default:
		//类型不对
	}
	publishItem.CmdID = resultcmdItemKey
	resultcmdItemsBys, _ := proto.Marshal(publishItem)

	logger.Logger.Info("%s createAssetTopicMsg publishItem:%+v", util.RunFuncName(), publishItem)
	logger.Logger.Print("%s createAssetTopicMsg publishItem:%+v", util.RunFuncName(), publishItem)

	return resultcmdItemsBys
}
