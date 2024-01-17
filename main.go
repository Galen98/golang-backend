package main

import (
	"backend/models"
	"backend/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Article{})

	r := routes.SetupRoutes(db)
	r.Run()
}
