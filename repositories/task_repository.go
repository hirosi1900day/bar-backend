package repositories

import (
	"bar-backend/models"
	"errors"
)

type TaskMemoryRepository struct {
	tasks []models.Task
}

type ITaskRepository interface {
	FindAll() (*[]models.Task, error)
	FindById(id uint) (*models.Task, error)
}

func NewTaskMemoryRepository(tasks []models.Task) ITaskRepository {
	return &TaskMemoryRepository{tasks: tasks}
}

func (r *TaskMemoryRepository) FindAll() (*[]models.Task, error) {
	return &r.tasks, nil
}

func (r *TaskMemoryRepository) FindById(id uint) (*models.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("record not found")
}
