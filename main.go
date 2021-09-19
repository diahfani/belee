package main

import (
	"Documents/belee/config"
	"Documents/belee/routes"
)

func main() {
	config.DbConfig()
	e := routes.NewRoutes()
	e.Start(":8000")

}
