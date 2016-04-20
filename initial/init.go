package initial

import (
	"github.com/elago/elapen/global"
	"github.com/elago/elapen/model"
	"github.com/elago/elapen/router"
)

func init() {
	initConfig()

	if !global.Install {
		initDbEngine()
		initGeoIP()
		model.InitModel()
	}

	router.InitRouter()
}
