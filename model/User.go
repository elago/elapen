package model

import (
	"github.com/go-elaeagnus/elaeagnus"
)

type User struct {
	elaeagnus.Pojo
}

func (this *User) Outerfunc() {

}

func init() {
	elaeagnus.RegisterModel(new(User))
}
