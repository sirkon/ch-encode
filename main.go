package main

import (
	"os"

	_ "github.com/mailru/go-clickhouse" // Mail.RU's clickhouse connector
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "test",
			Usage: "don't save generated code and write it into the stdout",
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
