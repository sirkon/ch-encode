package generator

// Int32 Type implementation
type Int32 struct {
	field     string
	fieldType string
}

// NewInt32 constructor
func NewInt32(field string, fieldType string) *Int32 {
	return &Int32{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (i32 *Int32) FieldName(gen Generator) string {
	return i32.field
}

// FieldTypeName ...
func (i32 *Int32) FieldTypeName(gen Generator) string {
	return i32.fieldType
}

// TypeName ...
func (i32 *Int32) TypeName(gen Generator) string {
	return gen.EasyTypeName(i32.field)
}

// ArgName ...
func (i32 *Int32) ArgName(gen Generator) string {
	return gen.VarName(i32.field)
}

// AccessName ...
func (i32 *Int32) AccessName(gen Generator) string {
	return gen.HelperName(i32.field)
}

// NativeTypeName ...
func (i32 Int32) NativeTypeName(gen Generator) string {
	return gen.Int32NativeTypeName()
}

// Encoding ...
func (i32 Int32) Encoding(source string, gen Generator) error {
	return gen.Int32Encoding(source)
}

// Helper ...
func (i32 *Int32) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (i32 Int32) TestingTypeName(gen Generator) string {
	return gen.Int32TestingTypeName()
}

// TestEncoding ...
func (i32 Int32) TestEncoding(source string, gen Generator) error {
	return gen.Int32TestEncoding(source)
}
