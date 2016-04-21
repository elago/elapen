package model

import "time"

type Project struct {
	Id          int64
	Name        string `xorm:"varchar(200)"`
	IconUrl     string
	AuthorUser  int64
	Description string

	Created time.Time `xorm:"created"`
}
