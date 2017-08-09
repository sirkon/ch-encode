package gogen

import (
	"fmt"
	"io"
	"text/template"

	"github.com/sirkon/gotify"
)

// GoGen encoder code generation for Go language
type GoGen struct {
	goish   *gotify.Gotify
	table   string
	dest    io.Writer
	imports map[string]string // path -> access name
}

// New Go generator constructor
func New(table string, g *gotify.Gotify, dest io.Writer) *GoGen {
	res := &GoGen{
		table:   table,
		dest:    dest,
		goish:   g,
		imports: map[string]string{},
	}
	res.useBytes()
	res.useBinEnc()
	res.useIO()
	return res
}

// RawData ...
func (gg *GoGen) RawData(v string) error {
	_, err := gg.dest.Write([]byte(v))
	return err
}

// regImport registers new import path
func (gg *GoGen) regImport(path, access string) {
	if prev, ok := gg.imports[path]; ok {
		if prev != access {
			panic(fmt.Errorf(`attempt to import path "%s" as %s, when it was registered before as %s`, path, access, prev))
		}
	}
	gg.imports[path] = access
}

func (gg *GoGen) useTime() {
	gg.regImport("time", "stdtime")
}

func (gg *GoGen) useBinEnc() {
	gg.regImport("github.com/sirkon/binenc", "")
}

func (gg *GoGen) useBytes() {
	gg.regImport("bytes", "")
}

func (gg *GoGen) useIO() {
	gg.regImport("io", "")
}

func (gg *GoGen) useUnsafe() {
	gg.regImport("unsafe", "")
}

func (gg *GoGen) useFmt() {
	gg.regImport("fmt", "")
}

const headerTemplate = `
package {{ .Package }}

import (
    {{ range $path, $access := .Imports }}{{ $access }} "{{ $path }}" {{ printf "\n" }}{{ end }})
`

// Header ...
func (gg *GoGen) Header(dest io.Writer) error {
	pkgname := gg.goish.Package(gg.table)
	t := template.New("encoder header")
	tmpl, err := t.Parse(headerTemplate)
	if err != nil {
		return err
	}
	var ctx struct {
		Package string
		Imports map[string]string
	}
	ctx.Package = pkgname
	ctx.Imports = gg.imports
	if err := tmpl.Execute(dest, ctx); err != nil {
		return err
	}
	return nil
}
