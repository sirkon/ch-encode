package generator

import "strings"

// FieldSet represents whole set of table fields
type FieldSet struct {
	g      Generator
	fields []Field
	nests  map[string][]Field
}

// NewFieldSet constructor
func NewFieldSet(g Generator) *FieldSet {
	return &FieldSet{
		g:     g,
		nests: map[string][]Field{},
	}
}

// Add adds field to the field set
func (fs *FieldSet) Add(field Field) {
	data := strings.Split(field.FieldName(fs.g), ".")
	fs.fields = append(fs.fields, field)
	if len(data) > 1 {
		fieldName := data[0]
		fs.nests[fieldName] = append(fs.nests[fieldName], field)
	}
}

// List returns list of fields
func (fs *FieldSet) List() []Field {
	return fs.fields
}

// Nests returns dictionary of { nested field → list of its subfields }
func (fs *FieldSet) Nests() map[string][]Field {
	return fs.nests
}
