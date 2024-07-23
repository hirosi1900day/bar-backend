package models

// Status 型の定義
type TaskStatus int

// Enum の定義
const (
	Todo TaskStatus = iota
	InProgress
	Done
)

type Task struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Status      TaskStatus `json:"status"`
	Description string     `json:"description"`
	DueDate     string     `json:"due_date"`
}
