package model

import (
	"time"
)

type Keyword struct {
	Id   int64
	Name string `xorm:varchar(50)`

	Created time.Time `xorm:"created"`
}
