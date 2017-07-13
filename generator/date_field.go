package generator

// Date Type implementation
type Date struct {
	field     string
	fieldType string
}

// NewDate constructor
func NewDate(field, fieldType string) *Date {
	return &Date{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (d *Date) FieldName(gen Generator) string {
	return d.field
}

// FieldTypeName ...
func (d *Date) FieldTypeName(gen Generator) string {
	return d.fieldType
}

// TypeName ...
func (d *Date) TypeName(gen Generator) string {
	return gen.UneasyTypeName(d.field)
}

// ArgName ...
func (d *Date) ArgName(gen Generator) string {
	return gen.VarName(d.field)
}

// AccessName ...
func (d *Date) AccessName(gen Generator) string {
	return gen.HelperName(d.field)
}

// NativeTypeName Date implementation
func (d *Date) NativeTypeName(gen Generator) string {
	return gen.Uint16NativeTypeName()
}

// Encoding ...
func (d *Date) Encoding(source string, gen Generator) error {
	return gen.Uint16Encoding(source)
}

// Helper ...
func (d *Date) Helper(gen Generator) error {
	return gen.DateHelpers(d)
}

// TestingTypeName ...
func (d *Date) TestingTypeName(gen Generator) string {
	return gen.DateTestingTypeName()
}

// TestEncoding ...
func (d *Date) TestEncoding(source string, gen Generator) error {
	return gen.DateTestEncoding(source)
}
