package controllers

import (
	"go-web-hello-world-orm/initializers"
	"go-web-hello-world-orm/models"

	"html/template"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Tasks []models.Task
}

func TasksIndex(c *gin.Context) {
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	pageData := IndexData{
		Tasks: tasks,
	}

	tmpl.Execute(c.Writer, pageData)
}
