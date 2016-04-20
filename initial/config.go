package initial

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/global"
	"github.com/gogather/com"
)

func initConfig() {
	customConfigExist := com.FileExist("custom/app.ini")
	if customConfigExist {
		ela.SetConfig("custom/app.ini")
		global.Install = false
	} else {
		global.Install = true
		configForInstall()
	}

	global.Config = ela.GetConfig()
}

func configForInstall() {
	config := ela.GetConfig()
	config.SetInt("_", "port", 3001)
	config.SetString("_", "runmode", "dev")
}
