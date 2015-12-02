package controller

import (
	"github.com/elago/ela"
	"github.com/elago/orm"
)

func F1(ctx ela.RequestContext) {
	orm.PrintModels()
	ctx.Write("hello world")
}

func F2(ctx ela.RequestContext) {
	orm.TestQuery()
	ctx.Write("hello function 2")
}
