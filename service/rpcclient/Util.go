package rpcclient

import "time"

func GetTimeNow() any {
	return time.Now().Format("2006/01/02 15:04:05")
}
