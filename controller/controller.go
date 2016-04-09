package controller

import (
	"github.com/elago/ela"
)

func F1(ctx ela.Context) {
	ctx.Data["name"] = "lijun"
	ctx.Data["id"] = -1
	
	ctx.ServeTemplate("index.html")
}

func F2(ctx ela.Context) {
	ctx.Write("hello function 2")
}

func F3(ctx ela.Context) {
	panic("panic test")
}

func F4(ctx ela.Context) {
	ctx.ServeTemplate("test.html")
}

func Error500(ctx ela.Context) {
	ctx.ServeError(500, "500.html")
}
