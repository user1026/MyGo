package main

import (
	"fmt"
	"gin01/src/utils"
)

func main() {
	fmt.Println(utils.GetNowTimeByOptions(true, "p", false, ""))
	//DataBase.Init()
	//r := routes.Routes()
	//r.Run(":9000")
	//defer DataBase.Db.Close()
}
