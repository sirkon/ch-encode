package gogen

import (
	"fmt"
	"text/template"

	"github.com/sirkon/ch-encode/generator"
)

// EncoderInterface ...
func (gg *GoGen) EncoderInterface(field *generator.FieldSet) error {
	text :=
		`
        type {{.encoderName}} interface {
            Encode({{.args}}) error
            InsertCount() int
            ErrorCount() int
        }
        `
	tmpl, err := template.New("interface").Parse(text)
	if err != nil {
		return err
	}
	return tmpl.Execute(gg.dest, map[string]string{
		"encoderName": gg.interfaceEncoderName(),
		"args":        gg.argList(field)},
	)
}

// EncoderDef ...
func (gg *GoGen) EncoderDef(*generator.FieldSet) error {
	text :=
		`
        type {{.encoderName}} struct {
            insertCounter int
         	buffer *bytes.Buffer
         	helper *binenc.Encoder
            zeroes []byte
         	dest   io.Writer
        }

        func New{{.encoderName}}(w io.Writer) *{{.encoderName}} {
         	buffer := &bytes.Buffer{}
         	buffer.Grow(4096)
         	return &{{.encoderName}}{
         		buffer: buffer,
         		helper: binenc.New(),
         		dest:   w,
                zeroes: make([]byte, 64),
         	}
        }

        func (enc *{{.encoderName}}) InsertCount() int {
            return enc.insertCounter
        }
        `
	tmpl, err := template.New("encoder").Parse(text)
	if err != nil {
		return err
	}

	return tmpl.Execute(gg.dest, map[string]string{
		"encoderName": gg.encoderName(),
	})
}

// EncodingMethod ...
func (gg *GoGen) EncodingMethod(fields *generator.FieldSet) (err error) {
	text := "func (enc *%s) Encode(%s) error {\nenc.buffer.Reset();\n"
	if err = gg.RawData(fmt.Sprintf(text, gg.encoderName(), gg.argList(fields))); err != nil {
		return
	}
	if err = gg.constraints(fields); err != nil {
		return
	}
	for _, field := range fields.List() {
		if err = field.Encoding(field.ArgName(gg), gg); err != nil {
			return
		}
		if err = gg.RawData("\n"); err != nil {
			return
		}
	}
	err = gg.RawData(`
        enc.insertCounter++;
        _, err := enc.dest.Write(enc.buffer.Bytes());
        return err}`)
	return
}
