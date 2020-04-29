package emq

import "vehicle_system/src/vehicle/emq/emq_client"

func Setup() {
	go startEmq()
}

func startEmq() {
	emq_client.GetEmqInstance().InitEmqClient()
}
