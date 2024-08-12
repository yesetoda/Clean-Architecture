package usecase

import "example/cleaner2/domain"

type UserUsecase interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
	CreateUser(user domain.User) (string, error)
	PromoteUser(username string, update interface{}) error
	DeleteUser(username string) error
	FilterUser(filter interface{}) []*domain.User
	Login(username, password string) (string, error)
}
