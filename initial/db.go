package initial

import (
	"fmt"
	"github.com/elago/elapen/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gogather/com/log"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func initDbEngine() {
	config := global.Config
	engineType, err := config.GetString("_", "engine")
	// if engine not define, it's install mode, skip database engine initial
	if err != nil || global.Install {
		return
	}

	if engineType == "mysql" {
		username := config.GetStringDefault("mysql", "username", "root")
		password := config.GetStringDefault("mysql", "password", "")
		host := config.GetStringDefault("mysql", "host", "127.0.0.1")
		port := config.GetIntDefault("mysql", "port", 3306)
		dbName := "elapen"
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, dbName)

		engine, err = xorm.NewEngine("mysql", dsn)
		if err != nil {
			log.Redln(err)
		} else {
			log.Greenf("use mysql ")
			log.Bluef("%s\n", dsn)
		}
	} else {
		dsn := config.GetStringDefault("sqlite", "path", "./data.sqlite")

		engine, err = xorm.NewEngine("sqlite3", dsn)
		if err != nil {
			log.Redln(err)
		} else {
			log.Greenf("use sqlite ")
			log.Bluef("%s\n", dsn)
		}
	}

	runMode := config.GetStringDefault("_", "runmode", "prod")
	if runMode == "dev" {
		global.Engine.ShowSQL(true)
	}

	global.Engine = engine
}
