package main

import (
	"example/cleaner/Infrastructure/database/dbmongo"
	ginrouter "example/cleaner/Infrastructure/router/routergin"
	"example/cleaner/Infrastructure/router/routergin/controller"
	"example/cleaner/usecases"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	uri, dbname, taskCollectionName, userCollectionName := os.Getenv("MongodbUri"), os.Getenv("MongodbName"), os.Getenv("TaskCollectionName"), os.Getenv("UserCollectionName")
	fmt.Println(uri, dbname, taskCollectionName, userCollectionName)
	taskCollection := dbmongo.NewCollection(dbname, taskCollectionName)
	userCollection := dbmongo.NewCollection(dbname, userCollectionName)
	tc := controller.GInTaskController{
		Taskusecase: usecases.TaskUsecase{
			Repository: taskCollection,
		},
	}
	uc := controller.GInUserController{
		Userusecase: usecases.UserUsecase{
			Repository: userCollection,
		},
	}
	ginrouter.Routers(uc, tc)
}
