package generator

// FixedString Type implementation
type FixedString struct {
	field        string
	fieldType    string
	stringLength int
}

// NewFixedString constructor
func NewFixedString(field, fieldType string, length int) *FixedString {
	return &FixedString{
		field:        field,
		fieldType:    fieldType,
		stringLength: length,
	}
}

// FieldName ...
func (fs *FixedString) FieldName(gen Generator) string {
	return fs.field
}

// FieldTypeName ...
func (fs *FixedString) FieldTypeName(gen Generator) string {
	return fs.fieldType
}

// TypeName ...
func (fs *FixedString) TypeName(gen Generator) string {
	return gen.EasyTypeName(fs.field)
}

// ArgName ...
func (fs *FixedString) ArgName(gen Generator) string {
	return gen.VarName(fs.field)
}

// AccessName ...
func (fs *FixedString) AccessName(gen Generator) string {
	return gen.HelperName(fs.field)
}

// NativeTypeName ...
func (fs *FixedString) NativeTypeName(gen Generator) string {
	return gen.FixedStringNativeTypeName()
}

// Encoding ...
func (fs *FixedString) Encoding(source string, gen Generator) error {
	return gen.FixedStringEncoding(source, fs.stringLength)
}

// Helper ...
func (fs *FixedString) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (fs *FixedString) TestingTypeName(gen Generator) string {
	return gen.FixedStringTestingTypeName()
}

// TestEncoding ...
func (fs *FixedString) TestEncoding(source string, gen Generator) error {
	return gen.FixedStringTestEncoding(source, fs.stringLength)
}
