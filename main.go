package main

import (
	"CMSGo/config"
	"CMSGo/routes"

	_ "github.com/mattn/go-sqlite3"

	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	db := config.InitDb()

	r := routes.SetupRoutes(db)

	if port == "" {
		r.Run(":5000")
	} else {
		r.Run(":" + port)
	}

}
