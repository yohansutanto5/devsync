package util

import "time"

func Date() time.Time {
	return time.Now()
}

func GetTimeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
