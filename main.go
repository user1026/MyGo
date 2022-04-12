package main

import (
	"gin01/src/routes"
)

func main() {
	r := routes.Routes()
	r.Run(":9000")
}
