package main

import (
	"github.com/go-elaeagnus/elaeagnus"
	"github.com/go-elaeagnus/webapp/model"
)

func f1(ctx elaeagnus.RequestContext) {
	ctx.Write("hello world")
}

func f2(ctx elaeagnus.RequestContext) {
	ctx.Write("hello function 2")
}

func main() {
	user := model.User{}
	user.Get()

	elaeagnus.Router("/hello1", f1)
	elaeagnus.Router("/hello2", f2)
	elaeagnus.Run()
}
