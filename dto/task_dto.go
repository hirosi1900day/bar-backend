package dto

import "bar-backend/models"

type TaskDto struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title" binding:"required,min=1,max=100"`
	Status      models.TaskStatus `json:"status" binding:"required,taskstatus"`
	Description string            `json:"description"`
	DueDate     string            `json:"due_date"`
}
