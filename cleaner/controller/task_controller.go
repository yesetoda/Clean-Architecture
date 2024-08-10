package controller

// import (
// 	"context"
// 	"example/cleaner/domain"
// 	"example/cleaner/usecases"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// type TaskController struct {
// 	Taskusecase usecases.TaskUsecase
// }

// func (tc *TaskController) HandleGetAllTasks(ctx context.Context) []*domain.Task {

// 	return tc.Taskusecase.GetAllTasks(ctx)
// }
// func (tc *TaskController) HandleGetTaskById(ctx context.Context, id string) (*domain.Task, error) {

// 	return tc.Taskusecase.GetTaskById(ctx, id)
// }
// func (tc *TaskController) HandleCreateTask(ctx context.Context, task domain.Task) (string, error) {

// 	return tc.Taskusecase.CreateTask(ctx, domain.Task{})
// }
// func (tc *TaskController) HandleUpdateTask(ctx context.Context, id string, update bson.M) error {

// 	return tc.Taskusecase.UpdateTask(ctx, id,update)
// }
// func (tc *TaskController) HandleDeleteTask(ctx context.Context, id string) error {

// 	return tc.Taskusecase.DeleteTask(ctx, id)
// }

// func (tc *TaskController) HandleFilterTask(ctx context.Context, filter bson.M) []*domain.Task {

// 	return tc.Taskusecase.FilterTask(ctx,filter)
// }
