package models

type Task struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}
