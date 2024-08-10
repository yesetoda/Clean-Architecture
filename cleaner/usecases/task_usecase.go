package usecases

import (
	"example/cleaner/domain"
	"example/cleaner/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type TaskUsecase struct {
	Repository repositories.TaskRepository
}

func (tuc *TaskUsecase) GetAllTasks() []*domain.Task {

	return tuc.Repository.GetAllTasks()
}
func (tuc *TaskUsecase) GetTaskById(id int) (*domain.Task, error) {

	return tuc.Repository.GetTaskById(id)
}
func (tuc *TaskUsecase) CreateTask(task domain.Task) (string, error) {

	return tuc.Repository.CreateTask(task)
}
func (tuc *TaskUsecase) UpdateTask(id int, update bson.M) error {

	return tuc.Repository.UpdateTask(id, update)
}
func (tuc *TaskUsecase) DeleteTask(id int) error {

	return tuc.Repository.DeleteTask(id)
}
func (tuc *TaskUsecase) FilterTask(filter bson.M) []*domain.Task {
	return tuc.Repository.FilterTask(filter)
}
