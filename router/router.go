package router

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/controller"
)

func init() {
	ela.Router("/", controller.IndexCtrl)
	ela.Router("/hello1", controller.F1)
	ela.Router("/hello2", controller.F2)
	ela.Router("/panic", controller.F3)
	ela.Router("/hello3", controller.F4)
	ela.Router("@500", controller.Error500)
	ela.Router("@404", controller.Error404)
}
