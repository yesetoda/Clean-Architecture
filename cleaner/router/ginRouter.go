package genealrouter

import (
	"example/cleaner/controller"
	"example/cleaner/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinRouter struct{

}
func NewGinRouter() *GinRouter {
	return &GinRouter{
	}
}

func (*GinRouter) Router(uc controller.GInGenaralController,tc controller.GInGenaralController){
	router := gin.Default()
	router.POST("/signup", uc.HandleCreateUser)
	router.POST("/login", uc.HandleLogin)

	tasks := router.Group("/task")
	tasks.Use(middleware.AuthMiddleware())
	{

		tasks.GET("/", middleware.UserMiddleware(), tc.HandleGetAllTasks)
		tasks.GET("/:id", middleware.UserMiddleware(), tc.HandleGetTaskById)
		tasks.GET("/filter", middleware.UserMiddleware(), tc.HandleFilterTask)
		tasks.POST("/", middleware.AdminMiddleware(), tc.HandleCreateTask)
		tasks.DELETE("/:id", middleware.AdminMiddleware(), tc.HandleDeleteTask)
		tasks.PATCH("/:id", middleware.AdminMiddleware(), tc.HandleUpdateTask)
	}
	users := router.Group("/user")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("/", middleware.AdminMiddleware(), uc.HandleGetAllUsers)
		users.GET("/:username", middleware.AdminMiddleware(), uc.HandleGetUserByUsername)
		users.GET("/filter", middleware.AdminMiddleware(), uc.HandleFilterUser)
		users.DELETE("/:username", middleware.AdminMiddleware(), uc.HandleDeleteUser)
		users.PATCH("/:username", middleware.AdminMiddleware(), uc.HandlePromote)
	}
	fmt.Println("all the routes are defined ")
	router.Run(":8080")
}