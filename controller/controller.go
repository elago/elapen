package controller

import (
	"github.com/elago/ela"
	"github.com/elago/orm"
	"github.com/elago/webapp/model"
	"github.com/gogather/com/log"
	// "net/http"
	"runtime"
)

func F1(ctx ela.RequestContext) {
	// ctx.SetStatus(200)
	// ctx.SetHeader("Content-Type", "text/html")
	// ctx.Write("<h1>hello world</h1>")

	ctx.SetStatus(404)
	ctx.Data["name"] = "lijun"
	ctx.Data["id"] = -1
	ctx.ServeTemplate("index.html")
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

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	log.Blueln(mem.Alloc)
	log.Blueln(mem.TotalAlloc)
	log.Blueln(mem.HeapAlloc)
	log.Blueln(mem.HeapSys)

}
