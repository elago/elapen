package controller

import (
	"github.com/elago/ela"
	"github.com/elago/orm"
	"github.com/elago/webapp/model"
	"github.com/gogather/com/log"
)

func F1(ctx ela.RequestContext) {
	orm.PrintModels()
	ctx.Write("hello world")
}

func F2(ctx ela.RequestContext) {
	// u := orm.TestQuery().(model.Users)
	// var u model.Users
	u := model.Users{Id: 8}
	orm.Get(&u)

	log.Pinkf("id: %d\n", u.Id)
	log.Pinkf("username: %s\n", u.Username)
	log.Pinkf("password: %s\n", u.Password)
	ctx.Write("hello function 2")
}
