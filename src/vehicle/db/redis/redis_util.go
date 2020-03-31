package redis

import (
	"fmt"
	"time"
	"vehicle_system/src/vehicle/util"
)
/**
获取所有key
 */
func (rConn *RedisConn) Vkeys() (result []string,err error){
	sessionRedis,err := rConn.GetRedisDB()
	if err!=nil{
		return nil,fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}

	return  sessionRedis.Keys("*").Val(),nil
}

/**
某个key是否存在
 */
func (rConn *RedisConn) KeyExist(key string)(error,int64){
	sessionRedis,err := rConn.GetRedisDB()
	if err!=nil{
		return fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error()),0
	}
	return nil,sessionRedis.Exists(key).Val()
}

/**需要时候在补充*/

/***************************************
设置
设置值会覆盖
 */
func (rConn *RedisConn) VSet(key string, value interface{}, expiration time.Duration)  error{
	sessionRedis,err := rConn.GetRedisDB()
	if err!=nil{
		return fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	err = sessionRedis.Set(key, value, 0).Err()
	//sessionRedis.Expire(key, expiration)
	if err!=nil{
		return fmt.Errorf("%s vset key:%s,value:%s err:%v",util.RunFuncName(),key,value,err.Error())
	}
	return nil
}
/**
设置过期时间 setnex
*/
func (rConn *RedisConn) VSetNex(key string, value interface{}, expiration time.Duration)  error{
	sessionRedis,err := rConn.GetRedisDB()
	if err!=nil{
		return fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	sessionRedis.SetNX(key, value, 0)
	sessionRedis.Expire(key, expiration)
	//err = sessionRedis.SetNX(key,value,expiration).Err()
	return nil
}

/**
获取值
 */
func (rConn *RedisConn) VGet(key string)  (interface{},error){
	sessionRedis,err := rConn.GetRedisDB()
	if err!=nil{
		return "",fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	ret,err := sessionRedis.Get(key).Result()
	if err!=nil{
		return "",fmt.Errorf("%s vget key:%s err:%v",util.RunFuncName(),key,err.Error())
	}

	return ret,err
}



/**
删除
 */
func (rConn *RedisConn) VDel(key ...string)  error{
	sessionRedis,err :=rConn.GetRedisDB()
	if err!=nil{
		return fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	err = sessionRedis.Del(key...).Err()
	if err!=nil{
		return fmt.Errorf("%s vdel key:%s err:%v",util.RunFuncName(),key,err.Error())
	}
	return err
}
/****************************************************************
hash get设置
 */
func (rConn RedisConn) VHGet(key, field string)  (string,error){
	sessionRedis,err  := rConn.GetRedisDB()
	if err!=nil{
		return "" ,fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	oldSessionkey:=sessionRedis.HGet(key,field).Val()
	return  oldSessionkey,nil
}

/**
hash set设置
*/
func (rConn RedisConn) VHSet(key, field string, value interface{})  error{
	sessionRedis,err :=rConn.GetRedisDB()
	if err!=nil{
		return fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	err = sessionRedis.HSet(key,field,value).Err()

	if err!=nil{
		return fmt.Errorf("%s vhashset key:%s,field:%s,value:%s err:%v",util.RunFuncName(),key,field,value,err.Error())
	}

	return err
}





/**********************************************
list操作
 */

func (rConn RedisConn) VLPush(key string, value interface{})  error{
	sessionRedis,err :=rConn.GetRedisDB()
	if err!=nil{
		return fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	err = sessionRedis.LPush(key,value).Err()
	if err!=nil{
		return fmt.Errorf("%s vlpush key:%s,value:%s,err:%v",util.RunFuncName(),key,value,err.Error())
	}
	return err
}

func (rConn RedisConn) VRPush(key string, value interface{})  error{
	sessionRedis,err :=rConn.GetRedisDB()
	if err!=nil{
		return fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	err = sessionRedis.RPush(key,value).Err()
	if err!=nil{
		return fmt.Errorf("%s vrpush key:%s,value:%s,err:%v",util.RunFuncName(),key,value,err.Error())
	}
	return err
}



func (rConn RedisConn) VLPop(key string)(string,error)  {
	sessionRedis,err :=rConn.GetRedisDB()
	if err!=nil{
		return "",fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	result,err :=sessionRedis.LPop(key).Result()
	if err!=nil{
		return "",fmt.Errorf("%s vlpop key:%s,err:%v",util.RunFuncName(),key,err.Error())
	}
	return result,err
}


func (rConn RedisConn) VRPop(key string)(string,error)  {
	sessionRedis,err :=rConn.GetRedisDB()
	if err!=nil{
		return "",fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	result,err :=sessionRedis.RPop(key).Result()
	if err!=nil{
		return "",fmt.Errorf("%s vrpop key:%s,err:%v",util.RunFuncName(),key,err.Error())
	}
	return result,err
}


func (rConn RedisConn) VBLPop(key string,timeout time.Duration)([]string,error)  {
	sessionRedis,err :=rConn.GetRedisDB()
	if err!=nil{
		return nil,fmt.Errorf("%s open redis db err:%v",util.RunFuncName(),err.Error())
	}
	result,err :=sessionRedis.BLPop(timeout,key).Result()
	if err!=nil{
		return nil,fmt.Errorf("%s vblpop key:%s,err:%v",util.RunFuncName(),key,err.Error())
	}
	return result,err
}
