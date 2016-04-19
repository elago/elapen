package router

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/controller"
	"github.com/elago/elapen/global"
)

func InitRouter() {
	if global.Install {
		ela.Router("/install", controller.Error404)
	}

	ela.Router("/", controller.IndexCtrl)
	ela.Router("/hello1", controller.F1)
	ela.Router("/hello2", controller.F2)
	ela.Router("/panic", controller.F3)
	ela.Router("/hello3", controller.F4)
	ela.Router("/article/:title", controller.F5)
	ela.Router("@500", controller.Error500)
	ela.Router("@404", controller.Error404)
}
