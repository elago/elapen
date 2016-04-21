package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/elago/ela"
	_ "github.com/elago/elapen/web/model"
	_ "github.com/elago/elapen/web/router"
)

var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "This command should will start web service",
	Description: `Serv provide access auth for repositories`,
	Action:      webServ,
}

func webServ(c *cli.Context) {
	ela.Run()
}
