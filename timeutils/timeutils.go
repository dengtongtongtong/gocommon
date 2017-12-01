package timeutils

import (
	"time"
)

func RoundGivenTimestampToMidnight(timestamp int64) int64 {
	t := time.Unix(timestamp, 0)
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return midnight.Unix()
}

func RoundCurrentTimestampToMidnight() int64 {
	t := time.Now()
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return midnight.Unix()
}
