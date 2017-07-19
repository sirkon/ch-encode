package gogen

import (
	"fmt"
	"io"

	"github.com/glossina/gotify"
)

// GoGen encoder code generation for Go language
type GoGen struct {
	goish *gotify.Gotify
	table string
	dest  io.Writer
}

// New Go generator constructor
func New(table string, g *gotify.Gotify, dest io.Writer) *GoGen {
	return &GoGen{
		table: table,
		dest:  dest,
		goish: g,
	}
}

// RawData ...
func (gg *GoGen) RawData(v string) error {
	_, err := gg.dest.Write([]byte(v))
	return err
}

// Header ...
func (gg *GoGen) Header() error {
	pkgname := gg.goish.Package(gg.table)
	_, err := fmt.Fprintf(
		gg.dest,
		`
			package %s

			import (
               stdtime "time"
               "github.com/DenisCheremisov/binenc"
               "bytes"
               "io"
            )
        `,
		pkgname)
	return err
}
