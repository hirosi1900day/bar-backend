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
	Create(task models.Task) (*models.Task, error)
	Update(task models.Task) (*models.Task, error)
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

func (r *TaskMemoryRepository) Create(task models.Task) (*models.Task, error) {
	task.ID = uint(len(r.tasks) + 1)
	r.tasks = append(r.tasks, task)
	return &task, nil
}

func (r *TaskMemoryRepository) Update(task models.Task) (*models.Task, error) {
	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = task
			return &task, nil
		}
	}
	return nil, errors.New("record not found")
}
