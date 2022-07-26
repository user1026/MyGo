package main

import (
<<<<<<< HEAD
	"gin01/src/DataBase"
=======
>>>>>>> 8305f918e55a1711e3d120a7f69d55a17b0d9af5
	"gin01/src/routes"
)

func main() {
<<<<<<< HEAD
	DataBase.Init()
	r := routes.Routes()
	r.Run(":9000")
	defer DataBase.Db.Close()
=======
	r := routes.Routes()
	r.Run(":9000")
>>>>>>> 8305f918e55a1711e3d120a7f69d55a17b0d9af5
}
