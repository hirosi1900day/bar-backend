package controllers

import (
	"bar-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
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

func (c *TaskController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	task, err := c.taskService.FindById(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": task})
}
