package generator

// Float32 Type implementation
type Float32 struct {
	field     string
	fieldType string
}

// NewFloat32 constructor
func NewFloat32(field string, fieldType string) *Float32 {
	return &Float32{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (f32 *Float32) FieldName(gen Generator) string {
	return f32.field
}

// FieldTypeName ...
func (f32 *Float32) FieldTypeName(gen Generator) string {
	return f32.fieldType
}

// TypeName ...
func (f32 *Float32) TypeName(gen Generator) string {
	return gen.EasyTypeName(f32.field)
}

// ArgName ...
func (f32 *Float32) ArgName(gen Generator) string {
	return gen.VarName(f32.field)
}

// AccessName ...
func (f32 *Float32) AccessName(gen Generator) string {
	return gen.HelperName(f32.field)
}

// NativeTypeName ...
func (f32 Float32) NativeTypeName(gen Generator) string {
	return gen.Float32NativeTypeName()
}

// Encoding ...
func (f32 Float32) Encoding(source string, gen Generator) error {
	return gen.Float32Encoding(source)
}

// Helper ...
func (f32 *Float32) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (f32 Float32) TestingTypeName(gen Generator) string {
	return gen.Float32TestingTypeName()
}

// TestEncoding ...
func (f32 Float32) TestEncoding(source string, gen Generator) error {
	return gen.Float32TestEncoding(source)
}
