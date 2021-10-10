package main

import (
	"fmt"
	"restapiGin/environment"
)

func main() {

	//db := models.SetupDB()
	//db.AutoMigrate(&models.Task{}, &models.User{}, &models.Category{})
	//
	//// Add foreignKey to table
	//db.Model(&models.Task{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	//db.Model(&models.Task{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	//db.Model(&models.Category{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	//
	////db.DropTable("tasks")
	////db.DropTable("users")
	////db.DropTable("categories")
	//
	//r := routes.SetupRoutes(db)
	//r.Run()

	//tokenString, err := service.GenerateJWT()
	//
	//if err != nil {
	//	fmt.Println("Error generating token string")
	//}
	//
	//fmt.Println(tokenString)

	viperenv := environment.ViperEnvVariable("STRONGEST_AVENGER")

	fmt.Printf("viper : %s = %s \n", "STRONGEST_AVENGER", viperenv)
}
