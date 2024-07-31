package controllers

import (
	"bar-backend/dto"
	"bar-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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

func (c *TaskController) Create(ctx *gin.Context) {
	var CreateTaskDto dto.CreateTaskDto
	if err := ctx.ShouldBindJSON(&CreateTaskDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := c.taskService.Create(CreateTaskDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": task})
}

func (c *TaskController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input = dto.UpdateTaskDto{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateItem, err := c.taskService.Update(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": updateItem})
}

func (c *TaskController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.taskService.Delete(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
