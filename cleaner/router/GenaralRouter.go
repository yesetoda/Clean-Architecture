package genealrouter

import "example/cleaner/controller"

type RouterController interface{
	Router(uc controller.GInGenaralController,tc controller.GInGenaralController)
}