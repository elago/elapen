package global

import (
	"github.com/elago/ela"
	"github.com/go-xorm/xorm"
	"github.com/oschwald/geoip2-golang"
)

var (
	Install bool
	Config  *ela.Config
	Engine  *xorm.Engine
	GeoIP   *geoip2.Reader
)
