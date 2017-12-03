package timeutils

import (
	"time"
)

func RoundToMidnight(t time.Time) int64 {
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return midnight.Unix()
}

func RoundToMonth(t time.Time) int64 {
	month := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	return month.Unix()
}

func RoundGivenTimestampToMidnight(timestamp int64) int64 {
	t := time.Unix(timestamp, 0)
	return RoundToMidnight(t)
}

func RoundCurrentTimestampToMidnight() int64 {
	t := time.Now()
	return RoundToMidnight(t)
}

func RoundGivenTimestampToMonth(timestamp int64) int64 {
	t := time.Unix(timestamp, 0)
	return RoundToMonth(t)
}

func RoundCurrentTimestampToMonth() int64 {
	t := time.Now()
	return RoundToMonth(t)
}
