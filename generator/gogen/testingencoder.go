package gogen

import (
	"fmt"
	"text/template"

	"github.com/sirkon/ch-encode/generator"
)

// TestDef ...
func (gg *GoGen) TestDef(fields *generator.FieldSet) error {
	text :=
		`
        type {{.record}} struct {
        {{ range .fields }}    {{.Name}} {{.Type}}
        {{ end}}
        }
        `
	type item struct {
		Name string
		Type string
	}
	items := make([]item, len(fields.List()))
	for i, field := range fields.List() {
		items[i].Name = field.AccessName(gg)
		items[i].Type = field.TestingTypeName(gg)
	}
	tmpl, err := template.New("testing").Parse(text)
	if err != nil {
		return err
	}
	return tmpl.Execute(gg.dest, map[string]interface{}{
		"record": gg.testingResultName(),
		"fields": items,
	})
}

// TestEncoderDef ...
func (gg *GoGen) TestEncoderDef(*generator.FieldSet) error {
	text :=
		`
        type {{.encoder}} struct {
            Result []{{.record}}
        }

        func New{{.encoder}}() *{{.encoder}} {
            return &{{.encoder}}{
                Result: make([]{{.record}}, 0, 3),
            }
        }

        func (enc *{{.encoder}}) InsertCount() int {
           return len(enc.Result);
        }
        func (enc *{{.encoder}}) ErrorCount() int {
           return 0;
        }
        `
	tmpl, err := template.New("testingEncoder").Parse(text)
	if err != nil {
		return err
	}
	return tmpl.Execute(gg.dest, map[string]string{
		"record":  gg.testingResultName(),
		"encoder": gg.testingEncoderName(),
	})
}

// TestEncodingMethod ...
func (gg *GoGen) TestEncodingMethod(fields *generator.FieldSet) (err error) {
	if err = gg.RawData(fmt.Sprintf("func (enc *%s)Encode(%s) error{", gg.testingEncoderName(), gg.argList(fields))); err != nil {
		return
	}
	if err = gg.constraints(fields); err != nil {
		return
	}
	if err = gg.RawData(fmt.Sprintf("enc.Result = append(enc.Result, %s{\n", gg.testingResultName())); err != nil {
		return
	}
	for _, field := range fields.List() {
		if _, err = fmt.Fprintf(gg.dest, "%s:", field.AccessName(gg)); err != nil {
			return
		}
		if err = field.TestEncoding(field.ArgName(gg), gg); err != nil {
			return
		}
		if err = gg.RawData("\n"); err != nil {
			return
		}
	}
	return gg.RawData("\n}); return nil}\n")
}
