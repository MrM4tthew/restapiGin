package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"restapiGin/controllers"
	"restapiGin/middleware"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.Use(middleware.CORSMiddleware())
	//r.Use(cors.Default())

	//Refresh route
	r.POST("/token/refresh", controllers.Refresh)

	//Route group for api
	api := r.Group("/api")

	addUsers(api)
	addTasks(api)
	addCategories(api)

	return r
}
