package redis

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/logger"
)

var ConnectOnce sync.Once
var SessionRedisClient *redis.Client


type RedisConn struct {}
var RedisConnInstance *RedisConn
/*单例*/
func GetRedisInstance() *RedisConn{
	ConnectOnce.Do(NewRedisConn)
	return RedisConnInstance
}


func NewRedisConn()  {
	if RedisConnInstance == nil{
		RedisConnInstance = new(RedisConn)
	}
}
func (redisConn *RedisConn) InitDataBase()(*redis.Client,error){
	var err error
	SessionRedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.SessionRedisAddr,
		Password: conf.SessionRedisPassword,
		DB:       conf.SessionRedisDB,
	})
	_, err = SessionRedisClient.Ping().Result()
	if err != nil {
		logger.Logger.Error("redis open error:%v", err)
		logger.Logger.Print("redis open error:%v", err)
		log.Fatalf("redis open error:%v\n", err)
	}
	return SessionRedisClient,err
}

func (m *RedisConn)  GetRedisDB() (rd_con *redis.Client,err error) {
	if SessionRedisClient == nil{
		SessionRedisClient = redis.NewClient(&redis.Options{
			Addr:     conf.SessionRedisAddr,
			Password: conf.SessionRedisPassword,
			DB:       conf.SessionRedisDB,
		})
		_, err = SessionRedisClient.Ping().Result()
		if err != nil {
			logger.Logger.Error("redis open error:%v", err)
			logger.Logger.Print("redis open error:%v", err)
			log.Fatalf("redis open error:%v", err)
		}
		return SessionRedisClient,err
	}
	return SessionRedisClient,err
}


