package services

import (
	"bar-backend/models"
	"bar-backend/repositories"
)

type ITaskService interface {
	FindAll() (*[]models.Task, error)
}

type TaskService struct {
	taskRepository repositories.ITaskRepository
}

func NewTaskService(taskRepository repositories.ITaskRepository) ITaskService {
	return &TaskService{taskRepository: taskRepository}
}

func (s *TaskService) FindAll() (*[]models.Task, error) {
	return s.taskRepository.FindAll()
}
