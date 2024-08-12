package main

import (
	"example/cleaner/controller"
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
	taskCollection := usecases.NewCollection(dbname, taskCollectionName)
	userCollection := usecases.NewCollection(dbname, userCollectionName)
	tc := controller.GInGenaralController{
		Generalusecase: usecases.GeneralUsecase{
			Repository: taskCollection,
		},
	}
	uc := controller.GInGenaralController{
		Generalusecase: usecases.GeneralUsecase{
			Repository: userCollection,
		},
	}
	fmt.Println("this is the router that am using")
	controller.Routers(uc, tc)
}
