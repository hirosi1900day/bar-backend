package services

import (
	"bar-backend/dto"
	"bar-backend/models"
	"bar-backend/repositories"
)

type ITaskService interface {
	FindAll() (*[]models.Task, error)
	FindById(id uint) (*models.Task, error)
	Create(createTaskInput dto.TaskDto) (*models.Task, error)
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

func (s *TaskService) FindById(id uint) (*models.Task, error) {
	return s.taskRepository.FindById(id)
}

func (s *TaskService) Create(createTaskInput dto.TaskDto) (*models.Task, error) {
	task := models.Task{
		ID:          createTaskInput.ID,
		Title:       createTaskInput.Title,
		Status:      createTaskInput.Status,
		Description: createTaskInput.Description,
		DueDate:     createTaskInput.DueDate,
	}
	return s.taskRepository.Create(task)
}
