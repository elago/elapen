package initial

import (
	"github.com/elago/elapen/global"
	"github.com/go-xorm/xorm"
	"github.com/gogather/com/log"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func initDbEngine() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./data.sqlite")
	if err != nil {
		log.Redln(err)
	}

	global.Engine = engine
}
