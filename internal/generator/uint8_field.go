package generator

// Uint8 Type implementation
type Uint8 struct {
	field     string
	fieldType string
}

// NewUint8 constructor
func NewUint8(field string, fieldType string) *Uint8 {
	return &Uint8{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (ui8 *Uint8) FieldName(gen Generator) string {
	return ui8.field
}

// FieldTypeName ...
func (ui8 *Uint8) FieldTypeName(gen Generator) string {
	return ui8.fieldType
}

// TypeName ...
func (ui8 *Uint8) TypeName(gen Generator) string {
	return gen.EasyTypeName(ui8.field)
}

// ArgName ...
func (ui8 *Uint8) ArgName(gen Generator) string {
	return gen.VarName(ui8.field)
}

// AccessName ...
func (ui8 *Uint8) AccessName(gen Generator) string {
	return gen.HelperName(ui8.field)
}

// NativeTypeName ...
func (ui8 Uint8) NativeTypeName(gen Generator) string {
	return gen.Uint8NativeTypeName()
}

// Encoding ...
func (ui8 Uint8) Encoding(source string, gen Generator) error {
	return gen.Uint8Encoding(source)
}

// Helper ...
func (ui8 *Uint8) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (ui8 Uint8) TestingTypeName(gen Generator) string {
	return gen.Uint8TestingTypeName()
}

// TestEncoding ...
func (ui8 Uint8) TestEncoding(source string, gen Generator) error {
	return gen.Uint8TestEncoding(source)
}
