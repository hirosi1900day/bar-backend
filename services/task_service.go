package services

import (
	"bar-backend/dto"
	"bar-backend/models"
	"bar-backend/repositories"
)

type ITaskService interface {
	FindAll() (*[]models.Task, error)
	FindById(id uint) (*models.Task, error)
	Create(createTaskInput dto.CreateTaskDto) (*models.Task, error)
	Update(id uint, UpdateTaskDto dto.UpdateTaskDto) (*models.Task, error)
	Delete(id uint) error
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

func (s *TaskService) Create(createTaskInput dto.CreateTaskDto) (*models.Task, error) {
	task := models.Task{
		Title:       createTaskInput.Title,
		Status:      models.Todo,
		Description: createTaskInput.Description,
		DueDate:     createTaskInput.DueDate,
	}
	return s.taskRepository.Create(task)
}

func (s *TaskService) Update(id uint, UpdateTaskDto dto.UpdateTaskDto) (*models.Task, error) {
	task, err := s.taskRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if UpdateTaskDto.Title != nil {
		task.Title = *UpdateTaskDto.Title
	}
	if UpdateTaskDto.Description != nil {
		task.Description = *UpdateTaskDto.Description
	}
	if UpdateTaskDto.DueDate != nil {
		task.DueDate = *UpdateTaskDto.DueDate
	}
	if UpdateTaskDto.Status != nil {
		task.Status = *UpdateTaskDto.Status
	}

	return s.taskRepository.Update(*task)
}

func (s *TaskService) Delete(id uint) error {
	return s.taskRepository.Delete(id)
}
