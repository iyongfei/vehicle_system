package push

import "time"

//chanLength
func GetPushChanLength(chanLength int) int {
	if chanLength <= 0 {
		return flowChanDefaultLength
	}
	return chanLength
}

//timeout
func GetWriteChanDuration(timer time.Duration) time.Duration {
	if timer <= 0 {
		return timer * time.Second
	}
	return writChanDuration
}
