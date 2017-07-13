package generator

// Enum16 Type implementation
type Enum16 struct {
	field     string
	fieldType string
	choices   map[string]int
}

// NewEnum16 constructor
func NewEnum16(field, fieldType string, choices map[string]int) *Enum16 {
	return &Enum16{
		field:     field,
		fieldType: fieldType,
		choices:   choices,
	}
}

// FieldName ...
func (e16 *Enum16) FieldName(gen Generator) string {
	return e16.field
}

// FieldTypeName ...
func (e16 *Enum16) FieldTypeName(gen Generator) string {
	return e16.fieldType
}

// TypeName ...
func (e16 *Enum16) TypeName(gen Generator) string {
	return gen.UneasyTypeName(e16.field)
}

// ArgName ...
func (e16 *Enum16) ArgName(gen Generator) string {
	return gen.VarName(e16.field)
}

// AccessName ...
func (e16 *Enum16) AccessName(gen Generator) string {
	return gen.HelperName(e16.field)
}

// NativeTypeName ...
func (e16 *Enum16) NativeTypeName(gen Generator) string {
	return gen.Int16NativeTypeName()
}

// Encoding *Enum16 implementation
func (e16 *Enum16) Encoding(source string, gen Generator) error {
	return gen.Int16Encoding(source)
}

// Helper ...
func (e16 *Enum16) Helper(gen Generator) error {
	return gen.EnumHelpers(e16, e16.choices)
}

// TestingTypeName ...
func (e16 *Enum16) TestingTypeName(gen Generator) string {
	return gen.EnumTestingTypeName()
}

// TestEncoding ...
func (e16 *Enum16) TestEncoding(source string, gen Generator) error {
	return gen.EnumTestEncoding(source, e16.choices)
}
