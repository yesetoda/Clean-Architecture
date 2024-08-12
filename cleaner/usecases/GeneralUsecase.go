package usecases

import (
	"example/cleaner/domain"
	"example/cleaner/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type GeneralUsecase struct {
	Repository repositories.GeneralRepository
}

func (guc *GeneralUsecase) GetAllTasks() []*domain.Task {

	return guc.Repository.GetAllTasks()
}
func (guc *GeneralUsecase) GetTaskById(id int) (*domain.Task, error) {

	return guc.Repository.GetTaskById(id)
}
func (guc *GeneralUsecase) CreateTask(task domain.Task) (string, error) {

	return guc.Repository.CreateTask(task)
}
func (guc *GeneralUsecase) UpdateTask(id int, update bson.M) error {

	return guc.Repository.UpdateTask(id, update)
}
func (guc *GeneralUsecase) DeleteTask(id int) error {

	return guc.Repository.DeleteTask(id)
}
func (guc *GeneralUsecase) FilterTask(filter bson.M) []*domain.Task {
	return guc.Repository.FilterTask(filter)
}



func (guc *GeneralUsecase) GetAllUsers() []*domain.User {

	return guc.Repository.GetAllUsers()
}
func (guc *GeneralUsecase) GetUserByUsername(username string) (*domain.User, error) {

	return guc.Repository.GetUserByUsername(username)
}
func (guc *GeneralUsecase) CreateUser(user domain.User) (string, error) {

	return guc.Repository.CreateUser(user)
}
func (guc *GeneralUsecase) PromoteUser(username string, update bson.M) error {

	return guc.Repository.PromoteUser(username, update)
}
func (guc *GeneralUsecase) DeleteUser(username string) error {

	return guc.Repository.DeleteUser(username)
}
func (guc *GeneralUsecase) FilterUser(filter bson.M) []*domain.User {
	return guc.Repository.FilterUser(filter)
}

func (guc *GeneralUsecase) Login(username,password string) (string,error) {
	return guc.Repository.Login(username,password)
}