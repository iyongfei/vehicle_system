package emq_cacha

import (
	"sync"
	"time"
)

const (
	OnLine       = "online"
	UpdateTime   = "update_time"
	DistanceTime = 120
)

var vehicleCache *VehicleCache

var syncLock sync.Mutex

//获取全局单例模式
func GetVehicleCache() *VehicleCache {
	if vehicleCache == nil {
		syncLock.Lock()
		defer syncLock.Unlock()
		vehicleCache = &VehicleCache{
			CacheMap: map[string]map[string]interface{}{},
		}
	}
	return vehicleCache
}

//  vid:{"online":true,time:1312}
type VehicleCache struct {
	CacheMap map[string]map[string]interface{}
}

//添加一个vehicle
func (vehicleCache *VehicleCache) Update(vkey string, flag bool) {
	keyMap := map[string]interface{}{}
	keyMap[OnLine] = flag
	keyMap[UpdateTime] = time.Now()
	vehicleCache.CacheMap[vkey] = keyMap
}

//清除vehicle
func (vehicleCache *VehicleCache) Clean(vkey string) {
	delete(vehicleCache.CacheMap, vkey)
}

//清除vehicle all key

func (vehicleCache *VehicleCache) CleanAllKey() {
	for k, _ := range vehicleCache.CacheMap {
		delete(vehicleCache.CacheMap, k)
	}
}

//判断某key的时间
func (vehicleCache *VehicleCache) JudgeKeyExpire(vkey string) (bool, bool) {

	var exist bool
	var pushFlag bool

	if keyValue, ok := vehicleCache.CacheMap[vkey]; ok {
		exist = true
		if updateTime, updateTimeOk := keyValue[UpdateTime]; updateTimeOk {

			updateTimer := updateTime.(time.Time)
			subTime := SubTime(updateTimer)

			if subTime > DistanceTime {
				//fmt.Println("存在，发送:::", updateTimer.Unix(), "now:::", time.Now().Unix(), "subTime:::", subTime, "vkey::", vkey)
				pushFlag = true
				return exist, pushFlag
			} else {
				pushFlag = false
				return exist, pushFlag
			}
		}
	} else {
		exist = false
		pushFlag = false
	}
	return exist, pushFlag
}

func SubTime(t time.Time) float64 {
	now := time.Now()
	subM := now.Sub(t)
	return subM.Seconds()
}
