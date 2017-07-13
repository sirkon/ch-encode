package generator

// Uint64 Type implementation
type Uint64 struct {
	field     string
	fieldType string
}

// NewUint64 constructor
func NewUint64(field string, fieldType string) *Uint64 {
	return &Uint64{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (ui64 *Uint64) FieldName(gen Generator) string {
	return ui64.field
}

// FieldTypeName ...
func (ui64 *Uint64) FieldTypeName(gen Generator) string {
	return ui64.fieldType
}

// TypeName ...
func (ui64 *Uint64) TypeName(gen Generator) string {
	return gen.EasyTypeName(ui64.field)
}

// ArgName ...
func (ui64 *Uint64) ArgName(gen Generator) string {
	return gen.VarName(ui64.field)
}

// AccessName ...
func (ui64 *Uint64) AccessName(gen Generator) string {
	return gen.HelperName(ui64.field)
}

// NativeTypeName ...
func (ui64 Uint64) NativeTypeName(gen Generator) string {
	return gen.Uint64NativeTypeName()
}

// Encoding ...
func (ui64 Uint64) Encoding(source string, gen Generator) error {
	return gen.Uint64Encoding(source)
}

// Helper ...
func (ui64 *Uint64) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (ui64 Uint64) TestingTypeName(gen Generator) string {
	return gen.Uint64TestingTypeName()
}

// TestEncoding ...
func (ui64 Uint64) TestEncoding(source string, gen Generator) error {
	return gen.Uint64TestEncoding(source)
}
