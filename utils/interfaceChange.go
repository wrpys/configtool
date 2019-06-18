package utils

import (
	"strconv"
	"time"
)

/**
interface转字符串
*/
func ToString(v interface{}) string {
	switch v := v.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	default:
		return "-"
	}
}

/**
interface转字时间
*/
func ToDate(v interface{}) time.Time {
	switch v := v.(type) {
	case time.Time:
		return v.Local()
	default:
		return time.Now()
	}
}
