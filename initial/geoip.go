package initial

import (
	"github.com/elago/elapen/global"
	"github.com/oschwald/geoip2-golang"
	"log"
)

var db *geoip2.Reader

func initGeoIP() {
	var err error
	config := global.Config
	path, err := config.GetString("GeoIP", "path")
	if err != nil {
		db, err = geoip2.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		global.GeoIP = db
	}
}
