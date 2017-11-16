package main

import (
	_ "github.com/mailru/go-clickhouse" // Mail.RU's clickhouse connector
	"github.com/sg3des/argum"
	_ "github.com/sirkon/binenc"  // Binary encoding library go get's "dependency", generated package will need it
	_ "github.com/sirkon/go-diff" // Diff for testing
	"github.com/sirkon/message"
)

// Args program arguments
type Args struct {
	Test   bool `argum:"--test" help:"don't save generated code, just show it in the stdout'"`
	Format struct {
		YAML string `argum:"--yaml-dict" help:"YAML-formatted gotifying dictionary path"`
		JSON string `argum:"--json-dict" help:"JSON-formatted gotifying dictionary path"`
	} `argum:"oneof,req"`
	DateField string   `argum:"--date-field,req" help:"field name to filter by a day"`
	Tables    []string `argum:"req,pos"`
}

func main() {
	var args Args
	argum.MustParse(&args)
	if len(args.Tables) == 0 {
		argum.PrintHelp(1)
	}
	message.Info(args.Tables)
	action(args)
}
