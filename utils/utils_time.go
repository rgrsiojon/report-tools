package utils

import "time"

type UtilTime struct{}

func (UtilTime) AppendTime(data []*time.Time, due *time.Time) []*time.Time {
	if due != nil {
		return append(data, due)
	}
	return data
}

func (UtilTime) CompareTwoTime(a, b *time.Time) bool {
	if a != nil && b != nil && a.Local().Format("2006-01-02") != b.Local().Format("2006-01-02") {
		return false
	}
	return true
}
