package repositories

import (
	"bar-backend/models"
)

type TaskMemoryRepository struct {
	tasks []models.Task
}

type ITaskRepository interface {
	FindAll() (*[]models.Task, error)
}

func NewTaskMemoryRepository(tasks []models.Task) ITaskRepository {
	return &TaskMemoryRepository{tasks: tasks}
}

func (r *TaskMemoryRepository) FindAll() (*[]models.Task, error) {
	return &r.tasks, nil
}
