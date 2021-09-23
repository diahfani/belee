package main

import (
	"final_project/belee/config"

	"final_project/belee/routes"
)

func main() {
	config.DbConfig()
	e := routes.NewRoutes()
	e.Start(":8000")

}
