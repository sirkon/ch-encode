package generator

// NullableString Type implementation
type NullableString struct {
	field     string
	fieldType string
}

// NewNullableString constructor
func NewNullableString(field, fieldType string) *NullableString {
	return &NullableString{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (s *NullableString) FieldName(gen Generator) string {
	return s.field
}

// FieldTypeName ...
func (s *NullableString) FieldTypeName(gen Generator) string {
	return s.fieldType
}

// TypeName ...
func (s *NullableString) TypeName(gen Generator) string {
	return gen.EasyTypeName(s.field)
}

// ArgName ...
func (s *NullableString) ArgName(gen Generator) string {
	return gen.VarName(s.field)
}

// AccessName ...
func (s *NullableString) AccessName(gen Generator) string {
	return gen.HelperName(s.field)
}

// NativeTypeName NullableString implementation
func (s *NullableString) NativeTypeName(gen Generator) string {
	return gen.NullableStringNativeTypeName()
}

// Encoding NullableString implementation
func (s *NullableString) Encoding(source string, gen Generator) error {
	return gen.NullableStringEncoding(source)
}

// Helper ...
func (s *NullableString) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (s *NullableString) TestingTypeName(gen Generator) string {
	return gen.NullableStringTestingTypeName()
}

// TestEncoding ...
func (s *NullableString) TestEncoding(source string, gen Generator) error {
	return gen.NullableStringTestEncoding(source)
}
