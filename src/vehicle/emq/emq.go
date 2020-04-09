package emq

import "vehicle_system/src/vehicle/emq/emq_client"

func Setup()  {
	emq_client.GetEmqInstance().InitEmqClient()
}

