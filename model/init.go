package model

import (
	"github.com/elago/elapen/initial"
	// "github.com/gogather/com/log"
)

func init() {
	engine := initial.GetEngine()
	engine.Sync2(new(User))
}
