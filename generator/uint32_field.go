package generator

// Uint32 Type implementation
type Uint32 struct {
	field     string
	fieldType string
}

// NewUint32 constructor
func NewUint32(field string, fieldType string) *Uint32 {
	return &Uint32{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (ui32 *Uint32) FieldName(gen Generator) string {
	return ui32.field
}

// FieldTypeName ...
func (ui32 *Uint32) FieldTypeName(gen Generator) string {
	return ui32.fieldType
}

// TypeName ...
func (ui32 *Uint32) TypeName(gen Generator) string {
	return gen.EasyTypeName(ui32.field)
}

// ArgName ...
func (ui32 *Uint32) ArgName(gen Generator) string {
	return gen.VarName(ui32.field)
}

// AccessName ...
func (ui32 *Uint32) AccessName(gen Generator) string {
	return gen.HelperName(ui32.field)
}

// NativeTypeName ...
func (ui32 Uint32) NativeTypeName(gen Generator) string {
	return gen.Uint32NativeTypeName()
}

// Encoding ...
func (ui32 Uint32) Encoding(source string, gen Generator) error {
	return gen.Uint32Encoding(source)
}

// Helper ...
func (ui32 *Uint32) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (ui32 Uint32) TestingTypeName(gen Generator) string {
	return gen.Uint32TestingTypeName()
}

// TestEncoding ...
func (ui32 Uint32) TestEncoding(source string, gen Generator) error {
	return gen.Uint32TestEncoding(source)
}
