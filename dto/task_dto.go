package dto

import "bar-backend/models"

type CreateTaskDto struct {
	Title       string `json:"title" binding:"required,min=1,max=100"`
	Description string `json:"description" binding:"required,min=1,max=100"`
	DueDate     string `json:"due_date"`
}

// updateはnilを許容するため、ポインタ型にしている
type UpdateTaskDto struct {
	Title       *string            `json:"title" binding:"omitnil,min=1,max=100"`
	Description *string            `json:"description" binding:"omitnil,min=1,max=100"`
	DueDate     *string            `json:"due_date"`
	Status      *models.TaskStatus `json:"status"`
}
