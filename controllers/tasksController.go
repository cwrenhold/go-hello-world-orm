package controllers

import (
	"go-web-hello-world-orm/initializers"
	"go-web-hello-world-orm/models"
	"go-web-hello-world-orm/utils"

	"html/template"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Tasks                []models.Task
	IncompleteTasksCount int
}

type TaskData struct {
	Task models.Task
}

func TasksIndex(c *gin.Context) {
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	incompleteCount := len(utils.Filter(tasks, func(task models.Task) bool {
		return !task.IsComplete
	}))

	pageData := IndexData{
		Tasks:                tasks,
		IncompleteTasksCount: incompleteCount,
	}

	tmpl.Execute(c.Writer, pageData)
}

func TasksCreate(c *gin.Context) {
	tmpl, err := template.ParseFiles("templates/create.html")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, nil)
}

func TaskCreatePost(c *gin.Context) {
	var task models.Task
	c.Bind(&task)

	initializers.DB.Create(&task)

	c.Redirect(302, "/")
}

func TaskEdit(c *gin.Context) {
	var task models.Task
	initializers.DB.First(&task, c.Param("id"))

	tmpl, err := template.ParseFiles("templates/edit.html")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	pageData := TaskData{
		Task: task,
	}

	tmpl.Execute(c.Writer, pageData)
}

func TaskEditPost(c *gin.Context) {
	var task models.Task
	c.Bind(&task)

	var taskToUpdate models.Task
	initializers.DB.First(&taskToUpdate, c.Param("id"))

	taskToUpdate.Description = task.Description

	initializers.DB.Save(&taskToUpdate)

	c.Redirect(302, "/")
}

func TaskDelete(c *gin.Context) {
	var task models.Task
	initializers.DB.First(&task, c.Param("id"))

	initializers.DB.Delete(&task)

	c.Redirect(302, "/")
}

func TaskMarkAsComplete(c *gin.Context) {
	var task models.Task
	initializers.DB.First(&task, c.Param("id"))

	task.IsComplete = true

	initializers.DB.Save(&task)

	c.Redirect(302, "/")
}
