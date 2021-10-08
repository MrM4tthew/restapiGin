package routes

import (
	"github.com/gin-gonic/gin"
	"restapiGin/controllers"
)

func addTasks(rg *gin.RouterGroup) {
	route := rg.Group("/tasks")

	route.GET("/", controllers.FindTasks)
	route.GET("/:id", controllers.FindTask)
	route.POST("/", controllers.CreateTask)
	route.PUT("/:id", controllers.UpdateTask)
	route.DELETE("/:id", controllers.DeleteTask)
}
