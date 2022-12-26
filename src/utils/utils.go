package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type anyType interface {
}

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

// SelectSql 根据传来的对象生成查询sql的条件
func SelectSql(x interface{}) (string, []anyType) {
	sql := "where"
	sqlArr := make([]string, 0)
	keyInfo := reflect.TypeOf(x)
	valueInfo := reflect.ValueOf(x)
	num := keyInfo.NumField()
	sqlValue := make([]anyType, num)
	for i := 0; i < num; i++ {
		key := keyInfo.Field(i)
		value := valueInfo.Field(i)
		if value.IsZero() != true {
			sqlArr = append(sqlArr, " "+key.Name+"=? ")
			sqlValue = append(sqlValue, value)
		}
	}
	sql = sql + strings.Join(sqlArr, "and")
	return sql, sqlValue
}
