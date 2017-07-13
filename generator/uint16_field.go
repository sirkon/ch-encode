package generator

// Uint16 Type implementation
type Uint16 struct {
	field     string
	fieldType string
}

// NewUint16 constructor
func NewUint16(field string, fieldType string) *Uint16 {
	return &Uint16{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (ui16 *Uint16) FieldName(gen Generator) string {
	return ui16.field
}

// FieldTypeName ...
func (ui16 *Uint16) FieldTypeName(gen Generator) string {
	return ui16.fieldType
}

// TypeName ...
func (ui16 *Uint16) TypeName(gen Generator) string {
	return gen.EasyTypeName(ui16.field)
}

// ArgName ...
func (ui16 *Uint16) ArgName(gen Generator) string {
	return gen.VarName(ui16.field)
}

// AccessName ...
func (ui16 *Uint16) AccessName(gen Generator) string {
	return gen.HelperName(ui16.field)
}

// NativeTypeName ...
func (ui16 Uint16) NativeTypeName(gen Generator) string {
	return gen.Uint16NativeTypeName()
}

// Encoding ...
func (ui16 Uint16) Encoding(source string, gen Generator) error {
	return gen.Uint16Encoding(source)
}

// Helper ...
func (ui16 *Uint16) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (ui16 Uint16) TestingTypeName(gen Generator) string {
	return gen.Uint16TestingTypeName()
}

// TestEncoding ...
func (ui16 Uint16) TestEncoding(source string, gen Generator) error {
	return gen.Uint16TestEncoding(source)
}
