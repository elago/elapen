package router

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/controller"
)

func InitRouter() {

	ela.BeforeController(controller.Before)

	ela.InstallRouter("/install", controller.Install)
	ela.Router("/", controller.IndexCtrl)
	ela.Router("/hello1", controller.F1)
	ela.Router("/hello2", controller.F2)
	ela.Router("/panic", controller.F3)
	ela.Router("/hello3", controller.F4)
	ela.Router("/article/:title", controller.F5)

	ela.InternalError(controller.Error500)
	ela.NotFountError(controller.Error404)
}
