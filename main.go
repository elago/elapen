package main

import (
	"github.com/codegangsta/cli"
	"github.com/elago/elapen/cmd"
	_ "github.com/elago/elapen/initial"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Elapen"
	app.Usage = "Writting my blog with Elapen!"

	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}

	app.Run(os.Args)
}
