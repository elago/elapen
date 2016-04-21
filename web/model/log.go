package model

import "time"

type Log struct {
	Id      int64
	User    int64
	Ip      string
	Message string

	Created time.Time `xorm:"created"`
}
