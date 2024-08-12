package controller

import (
	"example/cleaner2/internal/domain"
	"example/cleaner2/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GInGenaralController struct {
	Generalusecase usecases.GeneralUsecase
	ctx *gin.Context
}

func NewGInGenaralController(usecase *usecases.GeneralUsecase) *GInGenaralController {
	return &GInGenaralController{
		Generalusecase: *usecase,
		ctx: &gin.Context{},
	}
}
func (tc *GInGenaralController) HandleGetAllTasks(ctx *gin.Context) {

	ctx.IndentedJSON(200, tc.Generalusecase.GetAllTasks())
}

func (tc *GInGenaralController) HandleGetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}

	task, err := tc.Generalusecase.GetTaskById(intId)
	if err != nil {
		ctx.IndentedJSON(400, "could not find task with such id")
		return
	}
	ctx.IndentedJSON(200, task)
}
func (tc *GInGenaralController) HandleCreateTask(ctx *gin.Context) {
	id := ctx.Request.FormValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}
	task := domain.Task{
		Id:          intId,
		Title:       ctx.Request.FormValue("title"),
		Description: ctx.Request.FormValue("description"),
		Status:      ctx.Request.FormValue("status"),
		Duedate:     ctx.Request.FormValue("duedate"),
	}

	msg, err := tc.Generalusecase.CreateTask(task)
	if err != nil {
		ctx.IndentedJSON(400, msg)
		return
	}
	ctx.IndentedJSON(200, msg)
}
func (tc *GInGenaralController) HandleUpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}
	updated := bson.M{}
	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	status := ctx.Request.FormValue("status")
	duedate := ctx.Request.FormValue("duedate")
	if len(title) > 0 {
		updated["title"] = title
	}
	if len(description) > 0 {
		updated["description"] = description
	}
	if len(status) > 0 {
		updated["status"] = status
	}
	if len(duedate) > 0 {
		updated["duedate"] = duedate
	}
	update := bson.M{
		"$set": updated,
	}

	err = tc.Generalusecase.UpdateTask(intId, update)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}
	ctx.IndentedJSON(200, "update successful")
}
func (tc *GInGenaralController) HandleDeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}
	err = tc.Generalusecase.DeleteTask(intId)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}
	ctx.IndentedJSON(200, "delete sucessful")
}

func (tc *GInGenaralController) HandleFilterTask(ctx *gin.Context) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.M{}
	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	status := ctx.Request.FormValue("status")
	duedate := ctx.Request.FormValue("duedate")
	if len(title) > 0 {
		filter["title"] = title
	}
	if len(description) > 0 {
		filter["description"] = description
	}
	if len(status) > 0 {
		filter["status"] = status
	}
	if len(duedate) > 0 {
		filter["duedate"] = duedate
	}
	result := tc.Generalusecase.FilterTask(filter)
	ctx.IndentedJSON(200, result)
}

func (tc *GInGenaralController) HandleGetAllUsers(ctx *gin.Context) {

	ctx.IndentedJSON(200, tc.Generalusecase.GetAllUsers())
}
func (tc *GInGenaralController) HandleGetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := tc.Generalusecase.GetUserByUsername(username)
	if err != nil {
		ctx.IndentedJSON(400, "could not find User with such Username")
		return
	}
	ctx.IndentedJSON(200, user)
}
func (tc *GInGenaralController) HandleCreateUser(ctx *gin.Context) {

	User := domain.User{
		Username: ctx.Request.FormValue("username"),
		Password: ctx.Request.FormValue("password"),
	}

	msg, err := tc.Generalusecase.CreateUser(User)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": msg})
		return
	}
	ctx.IndentedJSON(200, gin.H{"message": msg})
}
func (tc *GInGenaralController) HandlePromote(ctx *gin.Context) {
	username := ctx.Param("username")
	update := bson.M{"$set": bson.M{"role": "admin"}}
	err := tc.Generalusecase.PromoteUser(username, update)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(200, "update successful")
}
func (tc *GInGenaralController) HandleDeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")

	err := tc.Generalusecase.DeleteUser(username)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(200, "delete sucessful")
}

func (tc *GInGenaralController) HandleFilterUser(ctx *gin.Context) {
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
	result := tc.Generalusecase.FilterUser(filter)
	ctx.IndentedJSON(200, result)
}

func (tc *GInGenaralController) HandleLogin(ctx *gin.Context) {
	// In a real application, authenticate the user (this is just an example)
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	token, err := tc.Generalusecase.Login(username, password)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "invalid credential"})
		return
	}
	ctx.IndentedJSON(200, gin.H{"token": token})

}
