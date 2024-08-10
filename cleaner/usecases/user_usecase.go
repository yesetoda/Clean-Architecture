package usecases

import (
	"example/cleaner/domain"
	"example/cleaner/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type UserUsecase struct {
	Repository repositories.UserRepository
}

func (tuc *UserUsecase) GetAllUsers() []*domain.User {

	return tuc.Repository.GetAllUsers()
}
func (tuc *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {

	return tuc.Repository.GetUserByUsername(username)
}
func (tuc *UserUsecase) CreateUser(user domain.User) (string, error) {

	return tuc.Repository.CreateUser(user)
}
func (tuc *UserUsecase) PromoteUser(username string, update bson.M) error {

	return tuc.Repository.PromoteUser(username, update)
}
func (tuc *UserUsecase) DeleteUser(username string) error {

	return tuc.Repository.DeleteUser(username)
}
func (tuc *UserUsecase) FilterUser(filter bson.M) []*domain.User {
	return tuc.Repository.FilterUser(filter)
}

func (tuc *UserUsecase) Login(username,password string) (string,error) {
	return tuc.Repository.Login(username,password)
}