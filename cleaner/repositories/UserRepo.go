package repositories

import (
	"example/cleaner/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
	CreateUser(User domain.User) (string, error)
	PromoteUser(username string, updateBson bson.M) error
	DeleteUser(username string) error
	FilterUser(filter bson.M) []*domain.User
	Login(username,password string) (string,error)
}
