package DataBase

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init() {
	database, err := sqlx.Open("mysql", "user:123456@tcp(localhost:9906)/user")
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if err != nil {
		fmt.Println("数据库链接错误")

	} else {
		fmt.Println("数据库链接成功")
	}
	Db = database
}
