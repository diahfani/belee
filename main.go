package main

import (
	"belee/config"

	"belee/routes"
)

func main() {
	config.DbConfig()
	e := routes.NewRoutes()
	e.Start(":8080")

}
