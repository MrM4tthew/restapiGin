package main

import (
	"restapiGin/models"
	"restapiGin/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{}, &models.User{}, &models.Category{})

	// Add foreignKey to table
	db.Model(&models.Task{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Task{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	db.Model(&models.Category{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	//db.DropTable("tasks")
	//db.DropTable("users")
	//db.DropTable("categories")

	r := routes.SetupRoutes(db)
	r.Run()
}
