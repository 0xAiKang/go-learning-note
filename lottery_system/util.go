package main

import "time"

// 实现字符串类型和时间类型的转换
func ParseStringToTime(tm string) (int64, error) {
	loc, _ := time.LoadLocation("Local")
	resultTime, err := time.ParseInLocation("2006-01-02 15:04:05", tm, loc)

	return resultTime.Unix(), err
}