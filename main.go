package main

import (
	"go-server/database/connection"
	"go-server/initializes"
	"go-server/start/routes"
)

func init() {
	initializes.LoadVariable()
	connection.ConnectToDb()
}

func main() {
	r := routes.InitRoutes()
	r.Run() // listen and serve on 0.0.0.0:8080
}
