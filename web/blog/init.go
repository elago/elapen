package blog

import (
	"github.com/elago/elapen/web/model"
	"github.com/elago/elapen/web/router"
)

func InitWeb() {
	model.InitModel()
	router.InitRouter()
}
