package routes

import (
	"gihub.com/EDEN-NN/controllers"
	"gihub.com/EDEN-NN/middlewares"
	"github.com/gin-gonic/gin"
)

var host = "localhost:8080"

func HandleRequests() {
	router := gin.Default()
	api := router.Group("/api").Use(middlewares.Auth())
	{
		api.GET("/index", controllers.Index)
	}
	router.POST("/token", controllers.GenerateToken)
	router.GET("/users", controllers.GetAllUsers).Use(middlewares.Auth())
	router.POST("/users", controllers.RegisterUser)
	router.DELETE("/user/:id", controllers.DeleteUser)
	router.GET("/user/:id", controllers.GetUserByID)
	router.PATCH("/user/:id", controllers.PatchUser)
	router.Run(host)
}
