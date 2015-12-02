package model

import (
	"github.com/elago/orm"
)

type Users struct {
	Id       int64
	Username string
	password string
}

func (this *Users) Outerfunc() {

}

func init() {
	orm.RegisterModel(new(Users))
}
