package generator

// NullableArray Type implementation
type NullableArray struct {
	field     string
	fieldType string
	meta      Field
}

// NewNullableArray constructor
func NewNullableArray(field, fieldType string, meta Field) *NullableArray {
	return &NullableArray{
		field:     field,
		fieldType: fieldType,
		meta:      meta,
	}
}

// FieldName ...
func (a *NullableArray) FieldName(gen Generator) string {
	return a.field
}

// FieldTypeName ...
func (a *NullableArray) FieldTypeName(gen Generator) string {
	return a.fieldType
}

// TypeName ...
func (a *NullableArray) TypeName(gen Generator) string {
	return gen.EasyTypeName(a.field)
}

// ArgName ...
func (a *NullableArray) ArgName(gen Generator) string {
	return gen.VarName(a.field)
}

// AccessName ...
func (a *NullableArray) AccessName(gen Generator) string {
	return gen.HelperName(a.field)
}

// NativeTypeName NullableArray implementation
func (a *NullableArray) NativeTypeName(gen Generator) string {
	return gen.NullableArrayNativeTypeName(a.meta)
}

// Encoding NullableArray implementation
func (a *NullableArray) Encoding(source string, gen Generator) error {
	return gen.NullableArrayEncoding(source, a.meta)
}

// Helper ...
func (a *NullableArray) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (a *NullableArray) TestingTypeName(gen Generator) string {
	return gen.NullableArrayTestingTypeName(a.meta)
}

// TestEncoding ...
func (a *NullableArray) TestEncoding(source string, gen Generator) error {
	return gen.NullableArrayTestEncoding(source, a.meta)
}
