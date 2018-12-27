package generator

// Decimal128 type implementation
type Decimal128 struct {
	field     string
	fieldType string
	precision int
	scale     int
}

func NewDecimal128(field string, fieldType string, precision int, scale int) *Decimal128 {
	return &Decimal128{field: field, fieldType: fieldType, precision: precision, scale: scale}
}

func (d *Decimal128) FieldName(gen Generator) string {
	return d.field
}

func (d *Decimal128) FieldTypeName(gen Generator) string {
	return d.fieldType
}

func (d *Decimal128) TypeName(gen Generator) string {
	return gen.EasyTypeName(d.field)
}

func (d *Decimal128) ArgName(gen Generator) string {
	return gen.VarName(d.field)
}

func (d *Decimal128) AccessName(gen Generator) string {
	return gen.HelperName(d.field)
}

func (d *Decimal128) NativeTypeName(gen Generator) string {
	panic("implement me")
}

func (d *Decimal128) Encoding(source string, gen Generator) error {
	panic("implement me")
}

func (d *Decimal128) Helper(gen Generator) error {
	panic("implement me")
}

func (d *Decimal128) TestingTypeName(gen Generator) string {
	panic("implement me")
}

func (d *Decimal128) TestEncoding(source string, gen Generator) error {
	panic("implement me")
}
