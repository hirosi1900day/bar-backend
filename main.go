package main

import (
	"bar-backend/controllers"
	"bar-backend/models"
	"bar-backend/repositories"
	"bar-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	tasks := []models.Task{
		{ID: 1, Title: "title1", Status: "status1", Description: "description", DueDate: "2021-01-01"},
		{ID: 2, Title: "title2", Status: "status2", Description: "description", DueDate: "2021-01-02"},
		{ID: 3, Title: "title3", Status: "status3", Description: "description", DueDate: "2021-01-03"},
	}
	task_repository := repositories.NewTaskMemoryRepository(tasks)
	task_service := services.NewTaskService(task_repository)
	task_controller := controllers.NewTaskController(task_service)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/tasks", task_controller.FindAll)
	r.GET("/tasks/:id", task_controller.FindById)
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
