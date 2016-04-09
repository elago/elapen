package model

import (
	"time"
)

type Article struct {
	Id       int64
	Title    string `xorm:"varchar(255)"`
	Uri      string `xorm:"varchar(255)"`
	Keywords string `xorm:"varchar(1000)"`
	Abstract string `xorm:"text"`
	Content  string `xorm:"text"`
	Author   int64
	Count    int64
	Status   int

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
