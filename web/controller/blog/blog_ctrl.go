package blog

import (
	"github.com/elago/ela"
)

func IndexCtrl(ctx ela.Context) {
	ctx.ServeTemplate("index.html")
}
