package controller

import (
	"github.com/elago/ela"
)

func F1(ctx ela.RequestContext) {
	ctx.Write("hello world")
}

func F2(ctx ela.RequestContext) {
	ctx.Write("hello function 2")
}
