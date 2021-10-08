package main

import (
	"restapiGin/models"
	"restapiGin/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{}, &models.User{}, &models.Category{})
	//db.DropTable("tasks")
	//db.DropTable("users")
	//db.DropTable("categories")

	r := routes.SetupRoutes(db)
	r.Run()
}
