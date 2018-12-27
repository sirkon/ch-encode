package generator

// Int64 Type implementation
type Int64 struct {
	field     string
	fieldType string
}

// NewInt64 constructor
func NewInt64(field string, fieldType string) *Int64 {
	return &Int64{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (i64 *Int64) FieldName(gen Generator) string {
	return i64.field
}

// FieldTypeName ...
func (i64 *Int64) FieldTypeName(gen Generator) string {
	return i64.fieldType
}

// TypeName ...
func (i64 *Int64) TypeName(gen Generator) string {
	return gen.EasyTypeName(i64.field)
}

// ArgName ...
func (i64 *Int64) ArgName(gen Generator) string {
	return gen.VarName(i64.field)
}

// AccessName ...
func (i64 *Int64) AccessName(gen Generator) string {
	return gen.HelperName(i64.field)
}

// NativeTypeName ...
func (i64 Int64) NativeTypeName(gen Generator) string {
	return gen.Int64NativeTypeName()
}

// Encoding ...
func (i64 Int64) Encoding(source string, gen Generator) error {
	return gen.Int64Encoding(source)
}

// Helper ...
func (i64 *Int64) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (i64 Int64) TestingTypeName(gen Generator) string {
	return gen.Int64TestingTypeName()
}

// TestEncoding ...
func (i64 Int64) TestEncoding(source string, gen Generator) error {
	return gen.Int64TestEncoding(source)
}
