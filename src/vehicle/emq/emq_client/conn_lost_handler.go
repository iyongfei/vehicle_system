package emq_client

import (
	"github.com/eclipse/paho.mqtt.golang"
	"time"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

/**
失去连接
 */
var connLostFlag = false
var ConconnlostReconnectTime time.Duration = 30

func Conconnlost(client mqtt.Client, err error)  {
	if connLostFlag == true{
		return
	}

	connLostFlag = true
	timer :=time.NewTimer(ConconnlostReconnectTime * time.Second)

	defer timer.Stop()
	select {
	case <- timer.C:
		logger.Logger.Print("%s",util.RunFuncName())
		logger.Logger.Info("%s",util.RunFuncName())

		connLostFlag = false
		GetEmqInstance().InitEmqClient()
		return
	}
}

