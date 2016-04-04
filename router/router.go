package router

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/controller"
)

func init() {
	ela.Router("/hello1", controller.F1)
	ela.Router("/hello2", controller.F2)
	ela.Router("/panic", controller.F3)
}
