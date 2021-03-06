package gogen

import (
	"strings"
	"text/template"

	"github.com/sirkon/ch-encode/internal/generator"
)

// DateFilterEncoderDef ...
func (gg *GoGen) DateFilterEncoderDef(field *generator.FieldSet) error {
	text :=
		`
		type {{.encoder}} struct {
			dayNo uint16
			encoder {{.interface}}
		}

		func New{{.encoder}}(dayNo uint16, enc {{.interface}}) *{{.encoder}} {
			return &{{.encoder}}{
				dayNo: dayNo,
				encoder: enc,
			}
		}

        func (enc *{{.encoder}}) InsertCount() int {
            return enc.encoder.InsertCount();
        }
        func (enc *{{.encoder}}) ErrorCount() int {
            return enc.encoder.ErrorCount();
        }
        `
	tmpl, err := template.New("filterer").Parse(text)
	if err != nil {
		return err
	}
	return tmpl.Execute(gg.dest, map[string]string{
		"encoder":   gg.filterEncoderName(),
		"interface": gg.interfaceEncoderName(),
	})
}

// DateFilterEncodingMethod ...
func (gg *GoGen) DateFilterEncodingMethod(dateArg string, fields *generator.FieldSet) error {
	text :=
		`
        func (enc *{{.encoder}}) Encode({{.args}}) error {
            if enc.dayNo != uint16({{.dateArg}}) {
                return nil
            }
            return enc.encoder.Encode({{.app}})
        }
        `
	tmpl, err := template.New("filterer").Parse(text)
	if err != nil {
		return err
	}
	app := []string{}
	for _, field := range fields.List() {
		app = append(app, field.ArgName(gg))
	}
	return tmpl.Execute(gg.dest, map[string]string{
		"encoder": gg.filterEncoderName(),
		"args":    gg.argList(fields),
		"app":     strings.Join(app, ", "),
		"dateArg": dateArg,
	})
}
