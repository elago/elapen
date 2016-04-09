package initial

import (
	"github.com/go-xorm/xorm"
	"github.com/gogather/com/log"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./data.sqlite")
	if err != nil {
		log.Redln(err)
	}
}

func GetEngine() *xorm.Engine {
	return engine
}
