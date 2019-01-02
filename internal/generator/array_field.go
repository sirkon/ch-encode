package generator

// Array Type implementation
type Array struct {
	field     string
	fieldType string
	meta      Field
}

// NewArray constructor
func NewArray(field, fieldType string, meta Field) *Array {
	return &Array{
		field:     field,
		fieldType: fieldType,
		meta:      meta,
	}
}

// FieldName ...
func (a *Array) FieldName(gen Generator) string {
	return a.field
}

// FieldTypeName ...
func (a *Array) FieldTypeName(gen Generator) string {
	return a.fieldType
}

// TypeName ...
func (a *Array) TypeName(gen Generator) string {
	return gen.EasyTypeName(a.field)
}

// ArgName ...
func (a *Array) ArgName(gen Generator) string {
	return gen.VarName(a.field)
}

// AccessName ...
func (a *Array) AccessName(gen Generator) string {
	return gen.HelperName(a.field)
}

// NativeTypeName Array implementation
func (a *Array) NativeTypeName(gen Generator) string {
	return gen.ArrayNativeTypeName(a.meta)
}

// Encoding Array implementation
func (a *Array) Encoding(source string, gen Generator) error {
	return gen.ArrayEncoding(source, a.meta)
}

// Helper ...
func (a *Array) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (a *Array) TestingTypeName(gen Generator) string {
	return gen.ArrayTestingTypeName(a.meta)
}

// TestEncoding ...
func (a *Array) TestEncoding(source string, gen Generator) error {
	return gen.ArrayTestEncoding(source, a.meta)
}
