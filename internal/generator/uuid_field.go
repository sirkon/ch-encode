package generator

// NewUUIDField генерация поля с типом UUID
func NewUUIDField(field string, fieldType string) *UUIDField {
	return &UUIDField{field: field, fieldType: fieldType}
}

var _ Field = &UUIDField{}

// UUIDField ...
type UUIDField struct {
	field     string
	fieldType string
}

// FieldName ...
func (f *UUIDField) FieldName(gen Generator) string {
	return f.field
}

// FieldTypeName ...
func (f *UUIDField) FieldTypeName(gen Generator) string {
	return f.fieldType
}

// TypeName ...
func (f *UUIDField) TypeName(gen Generator) string {
	return gen.UneasyTypeName(f.field)
}

// ArgName ...
func (f *UUIDField) ArgName(gen Generator) string {
	return gen.VarName(f.field)
}

// AccessName ...
func (f *UUIDField) AccessName(gen Generator) string {
	return gen.HelperName(f.field)
}

// NativeTypeName ...
func (f *UUIDField) NativeTypeName(gen Generator) string {
	return gen.UUIDNativeTypeName()
}

// Encoding ...
func (f *UUIDField) Encoding(source string, gen Generator) error {
	return gen.UUIDEncoding(source)
}

// Helper ...
func (f *UUIDField) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (f *UUIDField) TestingTypeName(gen Generator) string {
	return gen.UUIDTestingTypeName()
}

// TestEncoding ...
func (f *UUIDField) TestEncoding(source string, gen Generator) error {
	return gen.UUIDTestEncoding(source)
}
