package model

import "time"

type Artifact struct {
	Id        int64
	FileId    int64
	ProjectId int64
	CommitId  string
	From      int

	Created time.Time `xorm:"created"`
}
