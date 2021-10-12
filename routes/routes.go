package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"restapiGin/controllers"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	//Refresh route
	r.POST("/token/refresh", controllers.Refresh)

	//Route group for api
	api := r.Group("/api")

	addUsers(api)
	addTasks(api)
	addCategories(api)

	return r
}
