package gogen

import (
	"fmt"
	"io"
	"text/template"

	"github.com/sirkon/ch-encode/internal/generator"
	"github.com/sirkon/gotify"
)

var _ generator.Generator = &GoGen{}

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
func (gg *GoGen) regImport(importAs string, path string) {
	if prev, ok := gg.imports[path]; ok {
		if prev != importAs {
			panic(fmt.Errorf(`attempt to import path "%s" as %s, when it was registered before as %s`, path, importAs, prev))
		}
	}
	gg.imports[path] = importAs
}

func (gg *GoGen) useTime() {
	gg.regImport("stdtime", "time")
}

func (gg *GoGen) useBinEnc() {
	gg.regImport("", "github.com/sirkon/binenc")
}

func (gg *GoGen) useBytes() {
	gg.regImport("", "bytes")
}

func (gg *GoGen) useIO() {
	gg.regImport("", "io")
}

func (gg *GoGen) useUnsafe() {
	gg.regImport("", "unsafe")
}

func (gg *GoGen) useFmt() {
	gg.regImport("", "fmt")
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

const constraintCheck = `
if len({{.First}}) != len({{.Current}}) {
   return fmt.Errorf("length mismatch between {{.First}} and {{.Current}} (%d != %d), lengths must be the same being subfields of {{.Nest}}",
  len({{.First}}), len({{.Current}}))
}
`

// Constraints ...
func (gg *GoGen) constraints(fields *generator.FieldSet) error {
	nests := fields.Nests()
	t := template.New("constraint if")
	tmpl, err := t.Parse(constraintCheck)
	if err != nil {
		return err
	}
	var ctx struct {
		First   string
		Current string
		Nest    string
	}
	for nested, subfields := range nests {
		ctx.Nest = nested
		if len(subfields) <= 1 {
			continue
		}
		gg.useFmt()
		gg.dest.Write([]byte(fmt.Sprintf("\n// constraints on %s nested field", nested)))
		ctx.First = subfields[0].ArgName(gg)
		for _, restItem := range subfields[1:] {
			ctx.Current = restItem.ArgName(gg)
			if err = tmpl.Execute(gg.dest, ctx); err != nil {
				return err
			}
		}
	}
	return nil
}
