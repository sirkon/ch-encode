package gogen

import (
	"fmt"

	"github.com/DenisCheremisov/ch-encode/generator"
)

// Types ...
func (gg *GoGen) Types(fields []generator.Field) (err error) {
	for _, field := range fields {
		if err = gg.TypeDef(field); err != nil {
			return
		}
		if err = gg.HelperDef(field); err != nil {
			return
		}
	}
	return gg.RawData("\n")
}

// TypeDef ...
func (gg *GoGen) TypeDef(field generator.Field) (err error) {
	_, err = fmt.Fprintf(gg.dest, "type %s %s;\n", field.TypeName(gg), field.NativeTypeName(gg))
	return
}

// HelperDef ...
func (gg *GoGen) HelperDef(field generator.Field) (err error) {
	if err = field.Helper(gg); err != nil {
		return
	}
	return gg.RawData("\n")
}
