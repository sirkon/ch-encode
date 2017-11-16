package main

import (
	"os"

	_ "github.com/mailru/go-clickhouse" // Mail.RU's clickhouse connector
	_ "github.com/sirkon/binenc"        // Binary encoding library go get's "dependency", generated package will need it
	_ "github.com/sirkon/go-diff"       // Diff for testing
	"github.com/urfave/cli"
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
		cli.StringFlag{ // This is ugly solution for the lack of requirement flag
			Name:  "date-field",
			Value: "date",
			Usage: "Field name to filter by a day",
		},
	}
	app.Action = action
	app.Run(os.Args)

}
