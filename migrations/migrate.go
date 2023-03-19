package main

import (
	"go-web-hello-world-orm/initializers"
	"go-web-hello-world-orm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Task{})
}
