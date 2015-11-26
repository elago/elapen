package model

import (
	// "github.com/elago/ela"
	"github.com/elago/orm"
)

type User struct {
	Id       int64  "int(11)"
	Username string "char(50)"
	password string "char(128)"
}

func (this *User) Outerfunc() {

}

func init() {
	orm.RegisterModel(new(User))
}
