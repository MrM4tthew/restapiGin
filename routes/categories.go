package routes

import (
	"github.com/gin-gonic/gin"
	"restapiGin/controllers"
	"restapiGin/middleware"
)

func addCategories(rg *gin.RouterGroup) {
	route := rg.Group("/categories")
	route.GET("/", middleware.TokenAuthMiddleware(), controllers.FindCategories)
	route.GET("/:id", middleware.TokenAuthMiddleware(), controllers.FindCategory)
	route.POST("/", middleware.TokenAuthMiddleware(), controllers.CreateCategory)
}