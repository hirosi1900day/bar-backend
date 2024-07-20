package controllers

import (
	"bar-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	FindAll(c *gin.Context)
}

type TaskController struct {
	taskService services.ITaskService
}

func NewTaskController(taskService services.ITaskService) ITaskController {
	return &TaskController{taskService: taskService}
}

func (c *TaskController) FindAll(ctx *gin.Context) {
	tasks, err := c.taskService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": tasks})
}
