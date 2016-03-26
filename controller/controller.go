package controller

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/model"
	"github.com/elago/orm"
	"github.com/gogather/com/log"
	"net/http"
	"runtime"
	// "time"
)

func F1(ctx ela.Context) {
	// ctx.SetStatus(200)
	// ctx.SetHeader("Content-Type", "text/html")
	// ctx.Write("<h1>hello world</h1>")

	// method := ctx.GetMethod()

	// log.Pinkln(method)

	// ctx.SetStatus(404)
	ctx.Data["name"] = "lijun"
	ctx.Data["id"] = -1
	ctx.ServeTemplate("index.html")
}

func F2(ctx ela.Context) {
	// u := orm.TestQuery().(model.Users)
	// var u model.Users
	u := model.Users{Id: 8}
	orm.Get(&u)

	log.Pinkf("id: %d\n", u.Id)
	log.Pinkf("username: %s\n", u.Username)
	log.Pinkf("password: %s\n", u.Password)

	cookie := &http.Cookie{}
	cookie.Name = "ela"
	cookie.Value = "hello Rex Lee"
	cookie.Domain = "127.0.0.1"

	ctx.SetCookie(cookie)

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	log.Blueln(mem.Alloc)
	log.Blueln(mem.TotalAlloc)
	log.Blueln(mem.HeapAlloc)
	log.Blueln(mem.HeapSys)

	ctx.Write("hello function 2")

}
