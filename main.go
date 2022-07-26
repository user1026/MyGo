package main

import (
	"gin01/src/DataBase"
	"gin01/src/routes"
)

func main() {
	DataBase.Init()
	r := routes.Routes()
	r.Run(":9000")
	defer DataBase.Db.Close()
}
