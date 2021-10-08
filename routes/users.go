package routes

import (
	"github.com/gin-gonic/gin"
	"restapiGin/controllers"
)

func addUsers(rg *gin.RouterGroup) {
	route := rg.Group("/users")

	route.GET("/", controllers.GetUsers)
	route.GET("/:id", controllers.GetUser)
	route.POST("/register", controllers.Register)
}
