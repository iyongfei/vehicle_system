package tool

import "time"

func StampUnix2Time(timestamp int64) time.Time {
	datetime := time.Unix(timestamp, 0)

	return datetime
}

func TimeNowToUnix() uint32 {
	t := time.Now()
	return uint32(t.Unix())
}

func Str2Time(formatTimeStr string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime.Unix()
}
