package generator

// Decimal64 type implementation
type Decimal64 struct {
	field     string
	fieldType string
	precision int
	scale     int
}

func NewDecimal64(field string, fieldType string, precision int, scale int) *Decimal64 {
	return &Decimal64{field: field, fieldType: fieldType, precision: precision, scale: scale}
}

func (d *Decimal64) FieldName(gen Generator) string {
	return d.field
}

func (d *Decimal64) FieldTypeName(gen Generator) string {
	return d.fieldType
}

func (d *Decimal64) TypeName(gen Generator) string {
	return gen.EasyTypeName(d.field)
}

func (d *Decimal64) ArgName(gen Generator) string {
	return gen.VarName(d.field)
}

func (d *Decimal64) AccessName(gen Generator) string {
	return gen.HelperName(d.field)
}

func (d *Decimal64) NativeTypeName(gen Generator) string {
	return gen.Uint64NativeTypeName()
}

func (d *Decimal64) Encoding(source string, gen Generator) error {
	return gen.Uint64Encoding(source)
}

func (d *Decimal64) Helper(gen Generator) error {
	return nil
}

func (d *Decimal64) TestingTypeName(gen Generator) string {
	return gen.Dec64TestingTypeName()
}

func (d *Decimal64) TestEncoding(source string, gen Generator) error {
	return gen.Dec64TestEncoding(d.scale, source)
}
