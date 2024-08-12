package repo

import (
	"example/cleaner2/domain"
)

type GenaralRepo interface {
	GetAllTasks() []*domain.Task
	GetTaskById(id int) (*domain.Task, error)
	CreateTask(task domain.Task) (string, error)
	UpdateTask(id int, updateBson interface{}) error
	DeleteTask(id int) error
	FilterTask(filter interface{}) []*domain.Task

	
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
	CreateUser(User domain.User) (string, error)
	PromoteUser(username string, updateBson interface{}) error
	DeleteUser(username string) error
	FilterUser(filter interface{}) []*domain.User
	Login(username,password string) (string,error)
}
