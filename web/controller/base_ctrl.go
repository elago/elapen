package controller

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/global"
)

func Before(ctx ela.Context) {
	if global.Install {
		ctx.Redirect("/install")
	}
}

func Error500(ctx ela.Context, err error) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Write("服务器内部错误，以下是错误内容\n" + err.Error())
	// ctx.ServeError(500, "500.html")
}

func Error404(ctx ela.Context, err error) {
	ctx.ServeError(404, "404.html")
}
