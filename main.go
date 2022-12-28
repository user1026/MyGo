package main

import (
	"fmt"
	"gin01/src/DataBase"
	"gin01/src/global"
	"gin01/src/routes"
)

func main() {

	global.GLOBAL_DB = DataBase.InitDb()

	DB := global.GLOBAL_DB.DB()
	defer func() {
		if err := DB.Close(); err != nil {
			fmt.Println("数据库关闭失败")
		}
	}()
	r := routes.Routes()
	r.Run(":9000")

}
