package genealrouter

import "example/cleaner2/internal/controller"


type RouterController interface{
	Router(uc controller.GInGenaralController,tc controller.GInGenaralController)
}