package initial

import (
	"github.com/elago/elapen/model"
	"github.com/elago/elapen/router"
)

func init() {
	initConfig()
	initDbEngine()
	initGeoIP()
	model.InitModel()
	router.InitRouter()
}
