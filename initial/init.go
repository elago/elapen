package initial

import "github.com/elago/elapen/model"

func init() {
	initConfig()
	initDbEngine()
	initGeoIP()
	model.InitModel()
}
