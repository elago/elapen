package router

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/web/controller"
	"github.com/elago/elapen/web/controller/article"
	"github.com/elago/elapen/web/controller/blog"
	"github.com/elago/elapen/web/controller/install"
)

func InitRouter() {

	ela.BeforeController(controller.Before)

	ela.InstallRouter("/install", install.Install)
	ela.InstallRouter("/install", install.InstallSystem)

	ela.Router("/", blog.IndexCtrl)

	ela.Router("/article/:title", article.Article)

	ela.InternalError(controller.Error500)
	ela.NotFountError(controller.Error404)
}
