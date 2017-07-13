package generator

// Int16 Type implementation
type Int16 struct {
	field     string
	fieldType string
}

// NewInt16 constructor
func NewInt16(field string, fieldType string) *Int16 {
	return &Int16{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (i16 *Int16) FieldName(gen Generator) string {
	return i16.field
}

// FieldTypeName ...
func (i16 *Int16) FieldTypeName(gen Generator) string {
	return i16.fieldType
}

// TypeName ...
func (i16 *Int16) TypeName(gen Generator) string {
	return gen.EasyTypeName(i16.field)
}

// ArgName ...
func (i16 *Int16) ArgName(gen Generator) string {
	return gen.VarName(i16.field)
}

// AccessName ...
func (i16 *Int16) AccessName(gen Generator) string {
	return gen.HelperName(i16.field)
}

// NativeTypeName ...
func (i16 Int16) NativeTypeName(gen Generator) string {
	return gen.Int16NativeTypeName()
}

// Encoding ...
func (i16 Int16) Encoding(source string, gen Generator) error {
	return gen.Int16Encoding(source)
}

// Helper ...
func (i16 *Int16) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (i16 Int16) TestingTypeName(gen Generator) string {
	return gen.Int16TestingTypeName()
}

// TestEncoding ...
func (i16 Int16) TestEncoding(source string, gen Generator) error {
	return gen.Int16TestEncoding(source)
}
