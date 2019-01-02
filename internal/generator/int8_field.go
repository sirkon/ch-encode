package generator

// Int8 Type implementation
type Int8 struct {
	field     string
	fieldType string
}

// NewInt8 constructor
func NewInt8(field string, fieldType string) *Int8 {
	return &Int8{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (i8 *Int8) FieldName(gen Generator) string {
	return i8.field
}

// FieldTypeName ...
func (i8 *Int8) FieldTypeName(gen Generator) string {
	return i8.fieldType
}

// TypeName ...
func (i8 *Int8) TypeName(gen Generator) string {
	return gen.EasyTypeName(i8.field)
}

// ArgName ...
func (i8 *Int8) ArgName(gen Generator) string {
	return gen.VarName(i8.field)
}

// AccessName ...
func (i8 *Int8) AccessName(gen Generator) string {
	return gen.HelperName(i8.field)
}

// NativeTypeName ...
func (i8 Int8) NativeTypeName(gen Generator) string {
	return gen.Int8NativeTypeName()
}

// Encoding ...
func (i8 Int8) Encoding(source string, gen Generator) error {
	return gen.Int8Encoding(source)
}

// Helper ...
func (i8 *Int8) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (i8 Int8) TestingTypeName(gen Generator) string {
	return gen.Int8TestingTypeName()
}

// TestEncoding ...
func (i8 Int8) TestEncoding(source string, gen Generator) error {
	return gen.Int8TestEncoding(source)
}
