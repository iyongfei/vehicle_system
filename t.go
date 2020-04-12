package main

import "vehicle_system/src/vehicle/db/redis"

func main() {




	c:=redis.GetRedisInstance()
	c.VFlushDb()

}
