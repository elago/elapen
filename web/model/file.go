package model

import "time"

type File struct {
	Id         int64
	Filename   string `xorm:"varchar(200)"`
	Path       string
	Mime       string
	UploadUser int64
	Store      int

	Created time.Time `xorm:"created"`
}
