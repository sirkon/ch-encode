package generator

// Enum8 Type implementation
type Enum8 struct {
	field     string
	fieldType string
	choices   map[string]int
}

// NewEnum8 constructor
func NewEnum8(field, fieldType string, choices map[string]int) *Enum8 {
	return &Enum8{
		field:     field,
		fieldType: fieldType,
		choices:   choices,
	}
}

// FieldName ...
func (e8 *Enum8) FieldName(gen Generator) string {
	return e8.field
}

// FieldTypeName ...
func (e8 *Enum8) FieldTypeName(gen Generator) string {
	return e8.fieldType
}

// TypeName ...
func (e8 *Enum8) TypeName(gen Generator) string {
	return gen.UneasyTypeName(e8.field)
}

// ArgName ...
func (e8 *Enum8) ArgName(gen Generator) string {
	return gen.VarName(e8.field)
}

// AccessName ...
func (e8 *Enum8) AccessName(gen Generator) string {
	return gen.HelperName(e8.field)
}

// NativeTypeName ...
func (e8 *Enum8) NativeTypeName(gen Generator) string {
	return gen.Int8NativeTypeName()
}

// Encoding *Enum8 implementation
func (e8 *Enum8) Encoding(source string, gen Generator) error {
	return gen.Int8Encoding(source)
}

// Helper ...
func (e8 *Enum8) Helper(gen Generator) error {
	return gen.EnumHelpers(e8, e8.choices)
}

// TestingTypeName ...
func (e8 *Enum8) TestingTypeName(gen Generator) string {
	return gen.EnumTestingTypeName()
}

// TestEncoding ...
func (e8 *Enum8) TestEncoding(source string, gen Generator) error {
	return gen.EnumTestEncoding(source, e8.choices)
}
