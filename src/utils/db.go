package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Db() *sqlx.DB {
	database, err := sqlx.Open("mysql", "user:123456@tcp(127.0.0.1:9906)/user")
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if err != nil {
		fmt.Println("数据库链接错误")
		return nil
	}
	return database
}
