package main

import (
	"os"

	"github.com/docopt/docopt-go"
	_ "github.com/mailru/go-clickhouse" // Mail.RU's clickhouse connector
	_ "github.com/sirkon/binenc"        // Binary encoding library go get's "dependency", generated package will need it
	_ "github.com/sirkon/go-diff"       // Diff for testing
	"github.com/sirkon/message"
)

func main() {
	usage := `Generate type safe code to insert into Clickhouse tables

Usage:
  ch-encode [--test] [--yaml-dict <src> | --json-dict <src>] --date-field <date field> <table>...
  ch-encode -h | --help
  ch-encode --version

Options:
  -h --help                 Show this screen.
  --version                 Show version.
  --test                    Don't save generated code, just write it into the stdout
  --yaml-dict <src>         Use this YAML formatted dictionary to generate Goish names
  --json-dict <src>         Use this JSON formatted dictionary to generate Goish names
  --date-field <date field> Use this field as date
`
	arguments, err := docopt.Parse(usage, os.Args[1:], true, version, true)
	if err != nil {
		message.Fatal(err)
	}
	isTesting := arguments["--test"].(bool)
	var yamlDict string
	var jsonDict string
	var dateField string
	if arguments["--yaml-dict"] != nil {
		yamlDict = arguments["--yaml-dict"].(string)
	}
	if arguments["--json-dict"] != nil {
		jsonDict = arguments["--json-dict"].(string)
	}
	action(isTesting, yamlDict, jsonDict, dateField, arguments["<table>"].([]string))
}
