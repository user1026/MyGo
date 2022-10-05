package utils

import (
	"fmt"
	"time"
)

func GetNowTime() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return t
}
func GetNowYMD(JoinStr string) string {
	str := fmt.Sprintf("2006%s01%s02", JoinStr, JoinStr)
	t := time.Now().Format(str)
	return t
}
func GetNowHMS(JoinStr string) string {
	str := fmt.Sprintf("15%s04%s05", JoinStr, JoinStr)
	t := time.Now().Format(str)
	return t
}
