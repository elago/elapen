package router

import (
	"github.com/elago/ela"
	"github.com/elago/webapp/controller"
)

func init() {
	ela.Router("/hello1", controller.F1)
	ela.Router("/hello2", controller.F2)
}
