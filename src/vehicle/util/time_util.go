package util

import (
	"log"
	"time"
)


/*******************************字符串->时间对象*/
func Str2Time(formatTimeStr string) time.Time{
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime
}
/**字符串->时间戳*/
func Str2Stamp(formatTimeStr string) int64 {
	timeStruct:=Str2Time(formatTimeStr)
	millisecond:=timeStruct.UnixNano()/1e6
	return  millisecond
}

/*********************************时间对象->字符串*/
func Time2Str() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

func SpecificTime2Str(t time.Time) string {
	const shortForm = "2006-01-02 15:04:05"
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

/**时间对象->字符串*/
func Time2StrSimple(t time.Time) string {
	const shortForm = "2006-01-02"
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

func Time2StrFormat() string {
	const shortForm = "20060102150405"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

/*****************************时间对象->时间戳*/
func DataBaseTime2Stamp(t time.Time)int64{
	millisecond:=t.UnixNano()/1e6
	return  millisecond
}


/*时间对象->时间戳 毫秒*/
func Time2Stamp()int64{
	t:=time.Now()
	millisecond:=t.UnixNano()/1e6
	return  millisecond
}

/*时间对象->时间戳*/
func Time2StampUnix()int64{
	t:=time.Now()
	return  t.Unix()
}

func SpecificTime2StampUnix(t time.Time)int64{
	return  t.Unix()
}




/******************************时间戳->字符串*/


func TodayTimeFormat() string {
	today := time.Now().Format("15:04:05")
	return today
}
func Stamp2Str(stamp int64) string{
	timeLayout := "2006-01-02 15:04:05"
	str:=time.Unix(stamp/1000,0).Format(timeLayout)
	return str
}
//返回当天时间
func TodayFormat() string {
	today := time.Now().Format("2006-01-02")
	return today
}

func TodaySpecificFormat() string {
	today := time.Now().Format("2006-01-02 15:04:05")
	return today
}


func Stamp2Time(stamp int64)time.Time{
	stampStr:=Stamp2Str(stamp)
	timer:=Str2Time(stampStr)
	return timer
}

func StampUnix2Time(timestamp int64)time.Time{
	datetime := time.Unix(timestamp, 0)

	return datetime
}

func UnixStamp2Str(stamp int64) string{
	timeLayout := "2006-01-02 15:04:05"
	str:=time.Unix(stamp,0).Format(timeLayout)
	return str
}


/***
2个时间差多少
 */
func TimeSpace(start time.Time,end time.Time)int64{
	startUnix := start.Unix()
	endUnix := end.Unix()
	if endUnix < startUnix{
		return  -1
	}
	return endUnix - startUnix
}

// 传入一个 time.Time 时间，返回该日零点的时间戳
func CurrentDateZeroTime(t *time.Time) (*time.Time) {
	TimeZone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Panicf("fatal: load time zone (Asia/Shanghai) error: %v\n", err)
	}
	year, month, day := t.Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, TimeZone)
	return &date
}

// 传入一个 time.Time 时间，返回该日零点的时间戳
func CurrentDateLastTime(t *time.Time) (*time.Time) {
	TimeZone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Panicf("fatal: load time zone (Asia/Shanghai) error: %v\n", err)
	}
	year, month, day := t.Date()
	date := time.Date(year, month, day, 23, 59, 59, 0, TimeZone)
	return &date
}


// 传入一个 time.Time 时间，返回该日前两天的时间戳
func PerTwoDateTimestamp(t *time.Time) (*time.Time) {
	TimeZone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Panicf("fatal: load time zone (Asia/Shanghai) error: %v\n", err)
	}
	year, month, day := t.Date()
	day--
	day--
	date := time.Date(year, month, day, 0, 0, 0, 0, TimeZone)
	return &date
}

// 传入一个 time.Time 时间，返回次日零点的时间戳
func NextDateTimestamp(t *time.Time) (*time.Time) {
	TimeZone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Panicf("fatal: load time zone (Asia/Shanghai) error: %v\n", err)
	}
	year, month, day := t.Date()
	day++
	date := time.Date(year, month, day, 0, 0, 0, 0, TimeZone)
	return &date
}



//当前时间离明天还有多少秒
func SecondsToTommorow() time.Duration {
	curtime := time.Now()
	year, month, day := curtime.Date()
	today_time := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	tomorrow := today_time.Add(24 * time.Hour)
	timeElapse := tomorrow.Sub(curtime) / time.Second
	return timeElapse
}

//返回两个时间戳间隔之间的日期数组
func AllDaysBetweenTimeStamps(startStamp int,endStamp int) []string {
	allDays := []string{}
	for i:=startStamp; i<=endStamp; i+= 24*3600 *1000 {
		ti := Stamp2Time(int64(i))
		itemtime := ti.Format("2006-01-02")
		allDays = append(allDays,itemtime)
	}
	tiEnd := Stamp2Time(int64(endStamp))
	timeEnd := tiEnd.Format("2006-01-02")
	if allDays[len(allDays)-1] == timeEnd {
		return allDays
	}else {
		allDays = append(allDays,timeEnd)
		return allDays
	}
}

/**
一分钟前
 */

func BeforeOneMinute() time.Time {
	now := time.Now()
	m, _ := time.ParseDuration("-1m")
	return now.Add(m)
}

/**
一小时前
*/
func BeforeOneHour() time.Time {
	now := time.Now()
	m, _ := time.ParseDuration("-1h")
	return now.Add(m)
}




/**
一天前
*/
func BeforeOneDay() time.Time {
	now := time.Now()
	m, _ := time.ParseDuration("-24h")
	return now.Add(m)
}



/**
一分钟后
*/

func AfterOneMinute() time.Time {
	now := time.Now()
	m, _ := time.ParseDuration("1m")
	return now.Add(m)
}



func AddOneMinute() time.Time {
	return time.Now().Add(1*time.Minute)
}



/**
一小时后
*/
func AfterOneHour() time.Time {
	now := time.Now()
	m, _ := time.ParseDuration("1h")
	return now.Add(m)
}

/**
一天后
*/
func AfterOneDay() time.Time {
	now := time.Now()
	m, _ := time.ParseDuration("24h")
	return now.Add(m)
}

/**
计算2个时间差
 */

func SubTime(t time.Time) float64 {
	now := time.Now()

	subM := now.Sub(t)

	return subM.Minutes()
}

//https://www.jianshu.com/p/9d5636d34f17