package model

import (
	"github.com/elago/orm"
	"time"
)

type UserLog struct {
	Id         int64
	CreateTime time.Time
}

func init() {
	orm.RegisterModel(new(UserLog))
}
