package controller

import (
	"example/cleaner/domain"
	"example/cleaner/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GInTaskController struct {
	Taskusecase usecases.TaskUsecase
}

func (tc *GInTaskController) HandleGetAllTasks(ctx *gin.Context) {

	ctx.IndentedJSON(200, tc.Taskusecase.GetAllTasks())
}
func (tc *GInTaskController) HandleGetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}

	task, err := tc.Taskusecase.GetTaskById(intId)
	if err != nil {
		ctx.IndentedJSON(400, "could not find task with such id")
		return
	}
	ctx.IndentedJSON(200, task)
}
func (tc *GInTaskController) HandleCreateTask(ctx *gin.Context) {
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

	msg, err := tc.Taskusecase.CreateTask(task)
	if err != nil {
		ctx.IndentedJSON(400, msg)
		return
	}
	ctx.IndentedJSON(200, msg)
}
func (tc *GInTaskController) HandleUpdateTask(ctx *gin.Context) {
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

	err = tc.Taskusecase.UpdateTask(intId,update)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}
	ctx.IndentedJSON(200, "update successful")
}
func (tc *GInTaskController) HandleDeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}
	err = tc.Taskusecase.DeleteTask(intId)
	if err != nil {
		ctx.IndentedJSON(400, err)
		return
	}
	ctx.IndentedJSON(200, "delete sucessful")
}

func (tc *GInTaskController) HandleFilterTask(ctx *gin.Context) {
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
	result := tc.Taskusecase.FilterTask(filter)
	ctx.IndentedJSON(200, result)
}
