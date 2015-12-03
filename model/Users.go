package model

import (
	"github.com/elago/orm"
)

type Users struct {
	Id       int64
	Username string
	Password string
}

func init() {
	orm.RegisterModel(new(Users))
}
