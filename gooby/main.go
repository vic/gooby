package gooby

import (
	"./rbc"
	"github.com/codegangsta/cli"
	"os"
)

func Main() {
	app := cli.NewApp()
	app.Name = "gooby"
	app.Usage = "Rubinius compile and runtime in Go"

	app.Commands = []cli.Command{
		{
			Name:      "compile",
			ShortName: "c",
			Usage:     "Compile RBX bytecode into Go code",
			Action: func(c *cli.Context) {
				rbc.Compile(c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
