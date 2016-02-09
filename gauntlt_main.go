package gauntlt

import (
	"os"
	// Use Godeps
	"github.com/gauntlt/gauntlt-go/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func GauntltMain() {
	app := cli.NewApp()
	app.Name = "gauntlt"
	app.Usage = "a framework for rugged testing"
	app.Version = "0.0.1"
	app.ArgsUsage = "arguments"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run gauntlt",
			Action: func(c *cli.Context) {
				println("running gauntlt: ", c.Args().First())
			},
		},
	}
	app.Run(os.Args)

}
