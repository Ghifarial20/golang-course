package main

import (
	"go-jwt/database"
	routes "go-jwt/router"
)

func main() {
	database.StartDB()
	r := routes.StartApp()
	r.Run(":8080")
}
