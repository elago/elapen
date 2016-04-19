package initial

import (
	"github.com/elago/ela"
	"github.com/elago/elapen/global"
)

func initConfig() {
	ela.SetConfig("conf/app.ini")
	global.Config = ela.GetConfig()
}
