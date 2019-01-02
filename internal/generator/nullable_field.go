package generator

// Nullable Type implementation
type Nullable struct {
	field     string
	fieldType string
	meta      Field
}

// NewNullable constructor
func NewNullable(field, fieldType string, meta Field) *Nullable {
	return &Nullable{
		field:     field,
		fieldType: fieldType,
		meta:      meta,
	}
}

// FieldName ...
func (n *Nullable) FieldName(gen Generator) string {
	return n.field
}

// FieldTypeName ...
func (n *Nullable) FieldTypeName(gen Generator) string {
	return n.fieldType
}

// TypeName ...
func (n *Nullable) TypeName(gen Generator) string {
	return gen.EasyTypeName(n.field)
}

// ArgName ...
func (n *Nullable) ArgName(gen Generator) string {
	return gen.VarName(n.field)
}

// AccessName ...
func (n *Nullable) AccessName(gen Generator) string {
	return gen.HelperName(n.field)
}

// NativeTypeName Array implementation
func (n *Nullable) NativeTypeName(gen Generator) string {
	return gen.NullableNativeTypeName(n.meta)
}

// Encoding Array implementation
func (n *Nullable) Encoding(source string, gen Generator) error {
	return gen.NullableEncoding(source, n.meta)
}

// Helper ...
func (n *Nullable) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (n *Nullable) TestingTypeName(gen Generator) string {
	return gen.NullableTestingTypeName(n.meta)
}

// TestEncoding ...
func (n *Nullable) TestEncoding(source string, gen Generator) error {
	return gen.NullableTestEncoding(source, n.meta)
}
