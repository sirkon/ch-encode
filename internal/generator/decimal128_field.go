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
	return gen.UneasyTypeName(d.field)
}

func (d *Decimal128) ArgName(gen Generator) string {
	return gen.VarName(d.field)
}

func (d *Decimal128) AccessName(gen Generator) string {
	return gen.HelperName(d.field)
}

func (d *Decimal128) NativeTypeName(gen Generator) string {
	return gen.Dec128NativeTypeName()
}

func (d *Decimal128) Encoding(source string, gen Generator) error {
	return gen.Dec128Encoding(source)
}

func (d *Decimal128) Helper(gen Generator) error {
	return gen.Dec128Helpers(d)
}

func (d *Decimal128) TestingTypeName(gen Generator) string {
	return gen.Dec128TestingTypeName()
}

func (d *Decimal128) TestEncoding(source string, gen Generator) error {
	return gen.Dec128TestEncoding(d.scale, source)
}
