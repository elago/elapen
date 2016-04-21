package model

import (
	"time"
)

type User struct {
	Id       int64
	Username string
	Password string `xorm:"varchar(200)"`
	Salt     string
	Email    string `xorm:"varchar(200)"`
	Varified int

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
