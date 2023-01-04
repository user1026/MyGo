package DataBase

import (
	"fmt"
	"gin01/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

/*func Init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.UserName, config.Password, config.Port, config.MysqlName)
	database, err := sqlx.Open(config.SqlType, dataSourceName)
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if err != nil {
		fmt.Println("数据库链接错误")
	} else {
		fmt.Println("数据库链接成功")
	}
	Db = database
}*/
type DataApiGroup struct {
	UserApi
}

var DataApi = new(DataApiGroup)

func InitDb() *gorm.DB {
	DbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", config.UserName, config.Password, config.IP, config.Port, config.MysqlName)
	fmt.Println(DbStr)
	db, err := gorm.Open(mysql.Open(DbStr), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库链接错误", err)
	} else {
		fmt.Println("数据库链接成功")
	}
	Db = db
	return db
}
