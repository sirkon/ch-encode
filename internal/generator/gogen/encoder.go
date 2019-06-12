package gogen

import (
	"fmt"
	"text/template"

	"github.com/sirkon/ch-encode/internal/generator"
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
func (gg *GoGen) EncoderDef(fs *generator.FieldSet) error {
	text :=
		`
        type {{.encoderName}} struct {
            insertCounter int
         	buffer *bytes.Buffer
         	helper *binenc.Encoder
            zeroes []byte
         	dest   io.Writer
        }

		func checkSchema(curCols []chinsert.Column) error {
			sample := []chinsert.Column{
                {{ range .columns }}{Name: {{.Name}}, Type: {{.Type}}},
                {{ end }}}
            for i, col := range sample {
                if i >= len(curCols) {
					return fmt.Errorf("column %s(%s) is missing in the current schema for table {{.table}}, outdated encoder?", col.Name, col.Type)
                }
                var colNum string
                switch i {
				case 0:
					colNum = "1st"
				case 1:
					colNum = "2nd"
				case 2:
					colNum = "3rd"
				default:
					colNum = fmt.Sprintf("%dth", i+1)
				}
                if col.Name != curCols[i].Name || col.Type != curCols[i].Type {
					return fmt.Errorf("table {{.table}} encoder expects for a %s column to be %s(%s), while it is %s(%s) in actual schema", colNum, col.Name, col.Type, curCols[i].Name, curCols[i].Type) 
				}
            }
			if len(curCols) > len(sample) {
				mismatch := curCols[len(sample)]
				return fmt.Errorf("unexpected extra column %s(%s) in the current schema for table {{.table}}, outdated encoder?", mismatch.Name, mismatch.Type)
			}
			return nil
		}

        func New{{.encoderName}}(w chinsert.WriterWithSchemaCheck) (*{{.encoderName}},error) {
            columns, err := w.Schema()
            if err != nil {
				return nil,err
			}
			if err := checkSchema(columns); err != nil {
				return nil, err
			}
         	buffer := &bytes.Buffer{}
         	buffer.Grow(4096)
         	return &{{.encoderName}}{
         		buffer: buffer,
         		helper: binenc.New(),
         		dest:   w,
                zeroes: make([]byte, 64),
         	}, nil
        }

        func (enc *{{.encoderName}}) InsertCount() int {
            return enc.insertCounter
        }
        `
	gg.regImport("", "github.com/sirkon/ch-insert")
	gg.regImport("", "fmt")
	tmpl, err := template.New("encoder").Parse(text)
	if err != nil {
		return err
	}

	var fields []struct {
		Name string
		Type string
	}
	for _, f := range fs.List() {
		fields = append(fields, struct {
			Name string
			Type string
		}{Name: "`" + f.FieldName(gg) + "`", Type: "`" + f.FieldTypeName(gg) + "`"})
	}

	return tmpl.Execute(gg.dest, map[string]interface{}{
		"encoderName": gg.encoderName(),
		"columns":     fields,
		"table":       gg.table,
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
