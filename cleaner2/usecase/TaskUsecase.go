package usecase

import "example/cleaner2/domain"

type TaskUsecase interface {
	GetAllTasks() []*domain.Task
	GetTaskById(id int) (*domain.Task, error)
	CreateTask(task domain.Task) (string, error)
	UpdateTask(id int, update interface{}) error
	DeleteTask(id int) error
	FilterTask(filter interface{}) []*domain.Task
}

type TaskUsecaseService struct{
	
}