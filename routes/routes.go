package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//type routes struct {
//	router *gin.Engine
//}

//func SetupRoutes(db *gorm.DB) routes {
//	r := routes{
//		router: gin.Default(),
//	}
//
//	r.router.Use(func(c *gin.Context) {
//		c.Set("db", db)
//	})
//
//	//r := gin.Default()
//	api := r.router.Group("/api")
//
//	r.addUsers(api)
//	r.addTasks(api)
//
//	return r
//}

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	api := r.Group("/api")

	addUsers(api)
	addTasks(api)

	return r
}
