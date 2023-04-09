package main

import (
	"go-auth/database"
	"go-auth/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
