package controller

import (
	"example/cleaner/domain"
	"example/cleaner/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GInUserController struct {
	Userusecase usecases.UserUsecase
}

func (tc *GInUserController) HandleGetAllUsers(ctx *gin.Context) {

	ctx.IndentedJSON(200, tc.Userusecase.GetAllUsers())
}
func (tc *GInUserController) HandleGetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := tc.Userusecase.GetUserByUsername(username)
	if err != nil {
		ctx.IndentedJSON(400, "could not find User with such Username")
		return
	}
	ctx.IndentedJSON(200, user)
}
func (tc *GInUserController) HandleCreateUser(ctx *gin.Context) {

	User := domain.User{
		Username: ctx.Request.FormValue("username"),
		Password: ctx.Request.FormValue("password"),
	}

	msg, err := tc.Userusecase.CreateUser(User)
	if err != nil {
		ctx.IndentedJSON(400,  gin.H{"error":msg})
		return
	}
	ctx.IndentedJSON(200,  gin.H{"message":msg})
}
func (tc *GInUserController) HandlePromote(ctx *gin.Context) {
	username := ctx.Param("username")
	update := bson.M{"$set": bson.M{"role": "admin"}}
	err := tc.Userusecase.PromoteUser(username, update)
	if err != nil {
		ctx.IndentedJSON(400,  gin.H{"error":err.Error()})
		return
	}
	ctx.IndentedJSON(200, "update successful")
}
func (tc *GInUserController) HandleDeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")

	err := tc.Userusecase.DeleteUser(username)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error":err.Error()})
		return
	}
	ctx.IndentedJSON(200, "delete sucessful")
}

func (tc *GInUserController) HandleFilterUser(ctx *gin.Context) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.M{"role": ctx.Request.FormValue("role")}

	// user := ctx.Request.FormValue("description")
	// status := ctx.Request.FormValue("status")
	// duedate := ctx.Request.FormValue("duedate")
	// if len(title) > 0 {
	// 	filter["title"] = title
	// }
	// if len(description) > 0 {
	// 	filter["description"] = description
	// }
	// if len(status) > 0 {
	// 	filter["status"] = status
	// }
	// if len(duedate) > 0 {
	// 	filter["duedate"] = duedate
	// }
	result := tc.Userusecase.FilterUser(filter)
	ctx.IndentedJSON(200, result)
}

func (tc *GInUserController) HandleLogin(ctx *gin.Context) {
	// In a real application, authenticate the user (this is just an example)
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	token, err := tc.Userusecase.Login(username, password)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "invalid credential"})
		return
	}
	ctx.IndentedJSON(200, gin.H{"token": token})

}
