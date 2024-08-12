package repositories

import (
	"example/cleaner2/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type GeneralRepository interface {
	GetAllTasks() []*domain.Task
	GetTaskById(id int) (*domain.Task, error)
	CreateTask(task domain.Task) (string, error)
	UpdateTask(id int, updateBson bson.M) error
	DeleteTask(id int) error
	FilterTask(filter bson.M) []*domain.Task

	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
	CreateUser(User domain.User) (string, error)
	PromoteUser(username string, updateBson bson.M) error
	DeleteUser(username string) error
	FilterUser(filter bson.M) []*domain.User
	Login(username,password string) (string,error)
}
