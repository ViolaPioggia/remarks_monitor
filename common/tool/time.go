package tool

import "time"

// 获取下一天的零点时间
func GetNextExecutionTime() time.Time {
	now := time.Now()
	nextDay := now.Add(24 * time.Hour)
	nextExecutionTime := time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())
	return nextExecutionTime
}
