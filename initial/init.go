package initial

import (
	"github.com/elago/elapen/web/blog"
)

func init() {
	initConfig()
	initDbEngine()
	initGeoIP()
	blog.InitWeb()
}
