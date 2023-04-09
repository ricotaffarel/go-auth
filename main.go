package main

import (
	"go-auth/database"
	"go-auth/router"
	"os"
)

func main() {
	database.StartDB()
	var PORT = os.Getenv("PORT")
	router.StartApp().Run(":" + PORT)
}
