package emq

import "vehicle_system/src/vehicle/emq/emq_client"

func init()  {
	emq_client.GetEmqInstance().InitEmqClient()
}

