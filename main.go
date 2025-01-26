package main

import (
	"next-learn-go/infrastructure/database"

	"next-learn-go/router"

	"os"
)

func main() {

	db := database.NewDB()

	e := router.NewRouter(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
