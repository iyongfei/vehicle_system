package emq_service

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)

type Emqx struct {
	mqtt mqtt.Client
}

func NewEmqx() *Emqx {
	emqx:=&Emqx{
		mqtt:mqtt.NewClient(initClientOps()),
	}

	token:=emqx.mqtt.Connect()

	if token.Wait();token.Error() != nil{
		emqx.ConnectErr(token.Error())
		return  nil
	}
	return emqx
}

func  (emqx *Emqx) CreateTopic(vehicleId string) string {
	return fmt.Sprintf("%s/s/p",vehicleId)
}

func (emqx *Emqx) Publish(vehicleId string,payload interface{})  {
	token := emqx.mqtt.Publish(emqx.CreateTopic(vehicleId), 0, false, payload)
	if token.Wait() && token.Error() != nil {
		emqx.ConnectErr(token.Error())
	}
}

/**
log输出
 */
func (emqx *Emqx) ConnectErr(err error)  {
	fmt.Printf("emq connect err:[%v]\n",err)
}


