package gauntlt

import (
	"flag"
	"fmt"
	"os"
	// Use Godeps
	"github.com/gauntlt/gauntlt-go/Godeps/_workspace/src/github.com/codegangsta/cli"
)

type filters []string

func (f *filters) String() string {
	return fmt.Sprint(*f)
}

func (f *filters) Set(value string) error {
	*f = append(*f, value)
	return nil
}

var filterFlag filters

func GauntltMain() {
	app := cli.NewApp()
	app.Name = "gauntlt"
	app.Usage = "a framework for rugged testing"
	app.Version = "0.0.1"
	app.ArgsUsage = "arguments"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}
	var dir string

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run gauntlt",
			Action: func(c *cli.Context) {
				println("running gauntlt: ", c.Args().First())
				if flag.NArg() == 0 {
					dir = "examples"
				} else {
					dir = flag.Arg(0)
				}

				filt := []string{}
				for _, f := range filterFlag {
					filt = append(filt, string(f))
				}
				if err := BuildAndRunDir(dir, filt); err != nil {
					panic(err)
				}

			},
		},
	}
	app.Run(os.Args)

}
