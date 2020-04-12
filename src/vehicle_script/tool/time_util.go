package tool

import "time"

func StampUnix2Time(timestamp int64)time.Time{
	datetime := time.Unix(timestamp, 0)

	return datetime
}

func TimeNowToUnix()uint32{
	t:=time.Now()
	return  uint32(t.Unix())
}
