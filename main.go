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

	port := os.Getenv("PORT")
	log.Println("Starting on :" + port)
	r.Run(":" + port)
}
