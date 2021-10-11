package routes

import (
	"github.com/gin-gonic/gin"
	"restapiGin/controllers"
	"restapiGin/middleware"
)

func addTasks(rg *gin.RouterGroup) {
	route := rg.Group("/tasks")

	route.GET("/", middleware.TokenAuthMiddleware(), controllers.FindTasks)
	route.GET("/:id", middleware.TokenAuthMiddleware(), controllers.FindTask)
	route.POST("/", middleware.TokenAuthMiddleware(), controllers.CreateTask)
	route.PUT("/:id", middleware.TokenAuthMiddleware(), controllers.UpdateTask)
	route.DELETE("/:id", middleware.TokenAuthMiddleware(), controllers.DeleteTask)
}
