package controller

import (
	"github.com/elago/ela"
)

func IndexCtrl(ctx ela.Context) {
	ctx.ServeTemplate("index.html")
}

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

func Error500(ctx ela.Context, err error) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Write("服务器内部错误，以下是错误内容\n"+err.Error())
	// ctx.ServeError(500, "500.html")
}

func Error404(ctx ela.Context, err error) {
	ctx.ServeError(404, "404.html")
}
