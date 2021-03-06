package push

import (
	"fmt"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/emq/protobuf"
)

const ONLINE_STATUS = "ONLINE_STATUS"

var GWResult_ActionType_name = map[string]string{
	protobuf.GWResult_ActionType_name[0]: "defaults",
	protobuf.GWResult_ActionType_name[1]: "devices",
	protobuf.GWResult_ActionType_name[2]: "threats",
	//protobuf.GWResult_ActionType_name[3]:  "gw_infos",
	protobuf.GWResult_ActionType_name[3]:  "vehicle_infos",
	protobuf.GWResult_ActionType_name[4]:  "samples",
	protobuf.GWResult_ActionType_name[5]:  "protects",
	protobuf.GWResult_ActionType_name[6]:  "strategys",
	protobuf.GWResult_ActionType_name[7]:  "port_redirects",
	protobuf.GWResult_ActionType_name[8]:  "deployers",
	protobuf.GWResult_ActionType_name[9]:  "firmwares",
	protobuf.GWResult_ActionType_name[10]: "flowstats",
	protobuf.GWResult_ActionType_name[11]: "flow_strategystats",
	protobuf.GWResult_ActionType_name[12]: "monitor_infos",
	protobuf.GWResult_ActionType_name[13]: "flow_statistics",
	ONLINE_STATUS:                         "online_status",
}

func getPushReqUrl(interfaceName string) (url string) {

	urlName := GWResult_ActionType_name[interfaceName]

	url = fmt.Sprintf("http://%s:%d/%s", conf.PushHost, conf.PushPort, urlName)
	return
}
