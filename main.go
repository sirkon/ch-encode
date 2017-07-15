package main

import (
	"os"

	"github.com/urfave/cli"
	_ "github.com/DenisCheremisov/binenc"   // Binary encoding library go get's "dependency", generated package will need it
)

func main() {

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "test",
			Usage: "don't save generated code, just show it in the stdout",
		},
		cli.StringFlag{
			Name:  "yaml-dict",
			Value: "",
			Usage: "YAML-formatted gotifying dictionary",
		},
		cli.StringFlag{
			Name:  "json-dict",
			Value: "",
			Usage: "JSON-formatted gotifying dictionary",
		},
	}
	app.Action = action
	app.Run(os.Args)

}
