package main

import (
	"example/cleaner/controller"
	"example/cleaner/repositories"
	genealrouter "example/cleaner/router"
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
	taskCollection := repositories.NewCollection(dbname, taskCollectionName)
	userCollection := repositories.NewCollection(dbname, userCollectionName)
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
	fmt.Println("this is the router that am using",uc,tc)
	
	r := genealrouter.NewGinRouter()
	r.Router(uc,tc)
}
