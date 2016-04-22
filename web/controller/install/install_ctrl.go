package install

import (
	"github.com/elago/ela"
)

func Install(ctx ela.Context) {
	ctx.ServeTemplate("install/install.html")
}

func InstallSystem(ctx ela.Context) {

}
