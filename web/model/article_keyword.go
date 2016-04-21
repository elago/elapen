package model

import (
	"time"
)

type ArticleKeyword struct {
	Id        int64
	KeywordId int64
	ArticleId int64

	Created time.Time `xorm:"created"`
}
