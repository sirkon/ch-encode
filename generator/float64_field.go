package generator

// Float64 Type implementation
type Float64 struct {
	field     string
	fieldType string
}

// NewFloat64 constructor
func NewFloat64(field string, fieldType string) *Float64 {
	return &Float64{
		field:     field,
		fieldType: fieldType,
	}
}

// FieldName ...
func (f64 *Float64) FieldName(gen Generator) string {
	return f64.field
}

// FieldTypeName ...
func (f64 *Float64) FieldTypeName(gen Generator) string {
	return f64.fieldType
}

// TypeName ...
func (f64 *Float64) TypeName(gen Generator) string {
	return gen.EasyTypeName(f64.field)
}

// ArgName ...
func (f64 *Float64) ArgName(gen Generator) string {
	return gen.VarName(f64.field)
}

// AccessName ...
func (f64 *Float64) AccessName(gen Generator) string {
	return gen.HelperName(f64.field)
}

// NativeTypeName ...
func (f64 Float64) NativeTypeName(gen Generator) string {
	return gen.Float64NativeTypeName()
}

// Encoding ...
func (f64 Float64) Encoding(source string, gen Generator) error {
	return gen.Float64Encoding(source)
}

// Helper ...
func (f64 *Float64) Helper(gen Generator) error {
	return nil
}

// TestingTypeName ...
func (f64 Float64) TestingTypeName(gen Generator) string {
	return gen.Float64TestingTypeName()
}

// TestEncoding ...
func (f64 Float64) TestEncoding(source string, gen Generator) error {
	return gen.Float64TestEncoding(source)
}
