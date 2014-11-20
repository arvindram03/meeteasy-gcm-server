package helpers

import "time"

func ParseTime(timeString string) time.Time {
	time, _ := time.Parse(time.RFC822, timeString)
	return time
}
