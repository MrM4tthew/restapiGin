package routes

import (
	"github.com/gin-gonic/gin"
	"restapiGin/controllers"
	"restapiGin/middleware"
)

func addUsers(rg *gin.RouterGroup) {
	route := rg.Group("/users")

	route.GET("/", middleware.TokenAuthMiddleware(), controllers.GetUsers)
	route.GET("/:id", middleware.TokenAuthMiddleware(), controllers.GetUser)
	route.POST("/register", controllers.Register)
	route.POST("/login", controllers.Login)
	route.POST("/logout", middleware.TokenAuthMiddleware(), controllers.Logout)
}
