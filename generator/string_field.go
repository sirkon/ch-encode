package generator

// String Type implementation
type String struct {
	field     string
	fieldType string
}

// NewString constructor
func NewString(field, fieldType string) *String {
	return &String{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (s *String) FieldName(gen Generator) string {
	return s.field
}

// FieldTypeName ...
func (s *String) FieldTypeName(gen Generator) string {
	return s.fieldType
}

// TypeName ...
func (s *String) TypeName(gen Generator) string {
	return gen.EasyTypeName(s.field)
}

// ArgName ...
func (s *String) ArgName(gen Generator) string {
	return gen.VarName(s.field)
}

// AccessName ...
func (s *String) AccessName(gen Generator) string {
	return gen.HelperName(s.field)
}

// NativeTypeName String implementation
func (s *String) NativeTypeName(gen Generator) string {
	return gen.StringNativeTypeName()
}

// Encoding String implementation
func (s *String) Encoding(source string, gen Generator) error {
	return gen.StringEncoding(source)
}

// Helper ...
func (s *String) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (s *String) TestingTypeName(gen Generator) string {
	return gen.StringTestingTypeName()
}

// TestEncoding ...
func (s *String) TestEncoding(source string, gen Generator) error {
	return gen.StringTestEncoding(source)
}
