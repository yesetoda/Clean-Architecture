package genealrouter

import (
	"example/cleaner2/internal/controller"
	"example/cleaner2/internal/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinRouter struct{
	usercontroller *controller.GInGenaralController
	taskcontroller *controller.GInGenaralController
	auth middleware.GeneralAuth


}
func NewGinRouter(uc *controller.GInGenaralController,tc *controller.GInGenaralController,auth middleware.GeneralAuth) *GinRouter {
	return &GinRouter{
		usercontroller: uc,
		taskcontroller: tc,
		auth: auth,
		// userAuth:
	}
}

func (gr *GinRouter) Route(GMW,AMW,UMW gin.HandlerFunc){
	router := gin.Default()
	router.POST("/signup", gr.usercontroller.HandleCreateUser)
	router.POST("/login", gr.usercontroller.HandleLogin)

	tasks := router.Group("/task")
	
	tasks.Use(GMW)
	{

		tasks.GET("/", UMW, gr.taskcontroller.HandleGetAllTasks)
		tasks.GET("/:id", UMW, gr.taskcontroller.HandleGetTaskById)
		tasks.GET("/filter", UMW, gr.taskcontroller.HandleFilterTask)
		tasks.POST("/", AMW, gr.taskcontroller.HandleCreateTask)
		tasks.DELETE("/:id", AMW, gr.taskcontroller.HandleDeleteTask)
		tasks.PATCH("/:id", AMW, gr.taskcontroller.HandleUpdateTask)
	}
	users := router.Group("/user")
	users.Use(GMW)
	{
		users.GET("/", AMW, gr.usercontroller.HandleGetAllUsers)
		users.GET("/:username", AMW, gr.usercontroller.HandleGetUserByUsername)
		users.GET("/filter", AMW, gr.usercontroller.HandleFilterUser)
		users.DELETE("/:username", AMW, gr.usercontroller.HandleDeleteUser)
		users.PATCH("/:username", AMW, gr.usercontroller.HandlePromote)
	}
	fmt.Println("all the routes are defined ")
	router.Run(":8080")
}