package generator

// DateTime Type implementation
type DateTime struct {
	field     string
	fieldType string
}

// NewDateTime constructor
func NewDateTime(field, fieldType string) *DateTime {
	return &DateTime{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (dt *DateTime) FieldName(gen Generator) string {
	return dt.field
}

// FieldTypeName ...
func (dt *DateTime) FieldTypeName(gen Generator) string {
	return dt.fieldType
}

// TypeName ...
func (dt *DateTime) TypeName(gen Generator) string {
	return gen.UneasyTypeName(dt.field)
}

// ArgName ...
func (dt *DateTime) ArgName(gen Generator) string {
	return gen.VarName(dt.field)
}

// AccessName ...
func (dt *DateTime) AccessName(gen Generator) string {
	return gen.HelperName(dt.field)
}

// NativeTypeName ...
func (dt *DateTime) NativeTypeName(gen Generator) string {
	return gen.Uint32NativeTypeName()
}

// Encoding ...
func (dt *DateTime) Encoding(source string, gen Generator) error {
	return gen.Uint32Encoding(source)
}

// Helper ...
func (dt *DateTime) Helper(gen Generator) error {
	return gen.DateTimeHelpers(dt)
}

// TestingTypeName ...
func (dt *DateTime) TestingTypeName(gen Generator) string {
	return gen.DateTimeTestingTypeName()
}

// TestEncoding ...
func (dt *DateTime) TestEncoding(source string, gen Generator) error {
	return gen.DateTimeTestEncoding(source)
}
