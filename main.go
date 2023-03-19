package main

import (
	"go-web-hello-world-orm/controllers"
	"go-web-hello-world-orm/initializers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.TasksIndex)
	r.GET("/tasks/create", controllers.TasksCreate)
	r.POST("/tasks/create", controllers.TaskCreatePost)
	r.GET("/tasks/:id/edit", controllers.TaskEdit)
	r.POST("/tasks/:id/edit", controllers.TaskEditPost)
	r.GET("/tasks/:id/delete", controllers.TaskDelete)
	r.GET("/tasks/:id/complete", controllers.TaskMarkAsComplete)

	port := os.Getenv("PORT")
	log.Println("Starting on :" + port)
	r.Run(":" + port)
}
