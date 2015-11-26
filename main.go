package main

import (
	"github.com/elago/ela"
	"github.com/elago/webapp/model"
)

func f1(ctx ela.RequestContext) {
	ctx.Write("hello world")
}

func f2(ctx ela.RequestContext) {
	ctx.Write("hello function 2")
}

func main() {
	user := model.User{}
	user.Get()

	ela.Router("/hello1", f1)
	ela.Router("/hello2", f2)
	ela.Run()
}
