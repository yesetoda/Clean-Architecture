package main

import (
	"example/cleaner2/internal/controller"
	"example/cleaner2/internal/middleware"
	"example/cleaner2/internal/repositories"
	genealrouter "example/cleaner2/internal/router"
	"example/cleaner2/internal/usecases"
	"example/cleaner2/pkg"
	"os"
)
func main() {
	pkg.LoadEnv(".env")
	hasher := pkg.NewHasher()
	uri, dbname, taskCollectionName, userCollectionName := os.Getenv("MongodbUri"), os.Getenv("MongodbName"), os.Getenv("TaskCollectionName"), os.Getenv("UserCollectionName")
	TaskRepo := repositories.NewMongoRepository(uri, dbname, taskCollectionName, hasher)
	UserRepo := repositories.NewMongoRepository(uri, dbname, userCollectionName, hasher)
	userusecase := usecases.NewGeneralUsecase(UserRepo)
	taskusecase := usecases.NewGeneralUsecase(TaskRepo)
	tc := controller.NewGInGenaralController(taskusecase)
	uc := controller.NewGInGenaralController(userusecase)
	auth := middleware.NewJWTAuth()

	// info: is it only possible to do this using the taskcontroller tc? since its a new context maybe
	GMW := tc.AuthMiddlewareGIn(auth)
	AMW := tc.AdminMiddlewareGin(auth)
	UMW := tc.UserMiddlewareGin(auth)
	r := genealrouter.NewGinRouter(uc, tc, auth)
	r.Route(GMW, AMW, UMW)
}
