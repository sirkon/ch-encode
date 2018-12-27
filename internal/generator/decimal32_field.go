package generator

// Decimal32 type implementation
type Decimal32 struct {
	field     string
	fieldType string
	precision int
	scale     int
}

func NewDecimal32(field string, fieldType string, precision int, scale int) *Decimal32 {
	return &Decimal32{field: field, fieldType: fieldType, precision: precision, scale: scale}
}

func (d *Decimal32) FieldName(gen Generator) string {
	return d.field
}

func (d *Decimal32) FieldTypeName(gen Generator) string {
	return d.fieldType
}

func (d *Decimal32) TypeName(gen Generator) string {
	return gen.EasyTypeName(d.field)
}

func (d *Decimal32) ArgName(gen Generator) string {
	return gen.VarName(d.field)
}

func (d *Decimal32) AccessName(gen Generator) string {
	return gen.HelperName(d.field)
}

func (d *Decimal32) NativeTypeName(gen Generator) string {
	return gen.Uint32NativeTypeName()
}

func (d *Decimal32) Encoding(source string, gen Generator) error {
	return gen.Uint32Encoding(source)
}

func (d *Decimal32) Helper(gen Generator) error {
	panic("implement me")
}

func (d *Decimal32) TestingTypeName(gen Generator) string {
	panic("implement me")
}

func (d *Decimal32) TestEncoding(source string, gen Generator) error {
	panic("implement me")
}
