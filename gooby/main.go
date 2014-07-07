package gooby

import (
	"./compiler"
	"./runtime"
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
				compiler.CompileRbc(c.Args().First())
			},
		},
		{
			Name:      "interpret",
			ShortName: "i",
			Usage:     "Interpret RBX bytecode in Go runtime",
			Action: func(c *cli.Context) {
				runtime.InterpretRbc(c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
