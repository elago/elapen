package model

import (
	"github.com/elago/elapen/global"
)

func InitModel() {
	engine := global.Engine
	if engine == nil {
		return
	}

	engine.Sync2(new(User))
	engine.Sync2(new(Article))
	engine.Sync2(new(File))
	engine.Sync2(new(Project))
	engine.Sync2(new(Keyword))
	engine.Sync2(new(Artifact))
	engine.Sync2(new(ArticleKeyword))
	engine.Sync2(new(Log))
}
