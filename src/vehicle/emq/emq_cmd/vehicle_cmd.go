package emq_cmd

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"strconv"
	"strings"
)

type VehicleSetCmd struct {
	VehicleId   string
	CmdId int
	TaskType int

	Type   int
	Switch bool
}

func (setCmd *VehicleSetCmd) CreateProtoTopicMsg() interface{}{
	publishItem := &protobuf.Command{}

	//ItemType
	publishItem.ItemType = protobuf.Command_TaskType(setCmd.TaskType)
	//param
	gwSetParams := &protobuf.GwSetParam{}
	gwSetParams.Type = protobuf.GwSetParam_Type(setCmd.Type)
	gwSetParams.Switch = setCmd.Switch
	publishItem.Param, _ = proto.Marshal(gwSetParams)
	//CmdID
	resultcmdItemKey := strings.Join([]string{strconv.Itoa(setCmd.CmdId),
		strconv.Itoa(setCmd.Type)}, "_")
	publishItem.CmdID = resultcmdItemKey

	resultcmdItemsBys, _ := proto.Marshal(publishItem)
	return resultcmdItemsBys
}

