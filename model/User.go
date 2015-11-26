package model

import (
	"github.com/elago/ela"
)

type User struct {
	ela.Pojo
}

func (this *User) Outerfunc() {

}

func init() {
	ela.RegisterModel(new(User))
}
