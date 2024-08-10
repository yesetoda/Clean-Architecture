package main

import (
	"example/cleaner/Infrastructure/database/dbmongo"
	ginrouter "example/cleaner/Infrastructure/router/routergin"
	"example/cleaner/controller"
	"example/cleaner/usecases"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	_, dbname, taskCollectionName, userCollectionName := os.Getenv("MongodbUri"), os.Getenv("MongodbName"), os.Getenv("TaskCollectionName"), os.Getenv("UserCollectionName")
	// fmt.Println(uri, dbname, taskCollectionName, userCollectionName)
	client := dbmongo.GetNewMongoClient()
	taskCollection := dbmongo.NewMongoTaskRepository(client.Database(dbname), taskCollectionName)
	userCollection := dbmongo.NewMongoUserRepository(client.Database(dbname), userCollectionName)
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
