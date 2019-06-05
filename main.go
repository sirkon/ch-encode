package main

import (
	"os"

	cli "github.com/jawher/mow.cli"
	_ "github.com/mailru/go-clickhouse" // Mail.RU's clickhouse connector
	_ "github.com/sirkon/binenc"        // Binary encoding library go get's "dependency", generated package will need it
	_ "github.com/sirkon/go-diff"       // Diff for testing
	"github.com/sirkon/message"
)

const (
	version = "0.1.0"
)

func chDefaultParams() string {
	res := os.Getenv("CLICKHOUSE")
	if len(res) > 0 {
		return res
	}
	return "default@localhost:8123/default"
}

func main() {
	app := cli.App("ch-encode", "Go code generator for Clickhouse data insert")
	var (
		test      = app.BoolOpt("test", false, "Don't save generated code, just write it into the stdout")
		yamlDict  = app.StringOpt("yaml-dict", "", "Use this YAML formatted dictionary to generate Goish names")
		jsonDict  = app.StringOpt("json-dict", "", "User this JSON formatted dictionary to generate Goish names")
		dateField = app.StringOpt("date-field", "", "Use this field as a date")
		chConn    = app.StringOpt("clickhouse", chDefaultParams(), "Clickhouse connection params")
		tables    = app.StringsArg("TABLES", nil, "List of clickhouse tables")
	)
	app.Spec = `[--test] [--yaml-dict|--json-dict] [--date-field] [--clickhouse] TABLES...`
	app.Action = func() {
		if err := action(*test, *yamlDict, *jsonDict, *dateField, *tables, *chConn); err != nil {
			message.Fatal(err)
		}

	}

	if err := app.Run(os.Args); err != nil {
		message.Fatal(err)
	}
}
