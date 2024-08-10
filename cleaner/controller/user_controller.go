package controller

// import (
// 	"context"
// 	"example/cleaner/domain"
// 	"example/cleaner/usecases"

// 	"go.mongodb.org/mongo-driver/bson"
// )


// type UserController struct {
// 	Userusecase usecases.UserUsecase
// 	Contxt      context.Context
// }

// func (tc *UserController) HandleGetAllUsers(ctx context.Context) []*domain.User {
// 	return tc.Userusecase.GetAllUsers(ctx)
// }
// func (tc *UserController) HandleGetUserById(ctx context.Context) (*domain.User, error) {
// 	var id string

// 	return tc.Userusecase.GetUserById(ctx, id)
// }
// func (tc *UserController) HandleCreateUser(ctx context.Context) (string, error) {

// 	return tc.Userusecase.CreateUser(ctx, domain.User{})
// }
// func (tc *UserController) HandleUpdateUser(ctx context.Context) error {
// 	var id string
// 	var update bson.M

// 	return tc.Userusecase.UpdateUser(ctx, id, update)
// }
// func (tc *UserController) HandleDeleteUser(ctx context.Context) error {
// 	id := ""
// 	return tc.Userusecase.DeleteUser(ctx, id)
// }

// func (tc *UserController) HandleFilterUser(ctx context.Context) []*domain.User {
// 	filter := bson.M{}
// 	return tc.Userusecase.FilterUser(ctx, filter)
// }
// func (tc *UserController) HandleLogin(ctx context.Context) []*domain.User {
// 	filter := bson.M{}
// 	return tc.Userusecase.FilterUser(ctx, filter)
// }

// func (tc *UserController) HandleSignUp(ctx context.Context) []*domain.User {
// 	filter := bson.M{}

// 	return tc.Userusecase.FilterUser(ctx, filter)
// }
