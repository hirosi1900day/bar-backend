package dto

type TaskDto struct {
	Title       string `json:"title" binding:"required,min=1,max=100"`
	Description string `json:"description" binding:"required,min=1,max=100"`
	DueDate     string `json:"due_date"`
}
