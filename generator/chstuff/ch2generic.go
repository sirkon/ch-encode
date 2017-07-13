package chstuff

import (
	"fmt"

	"github.com/DenisCheremisov/ch-encode/generator"
)

// Meta2Field translates clickhouse field metainfo into generator Field object
func Meta2Field(meta FieldMeta) (field generator.Field) {
	switch meta.Type {
	case "Int8":
		field = generator.NewInt8(meta.Name, meta.Type)
	case "Int16":
		field = generator.NewInt16(meta.Name, meta.Type)
	case "Int32":
		field = generator.NewInt32(meta.Name, meta.Type)
	case "Int64":
		field = generator.NewInt64(meta.Name, meta.Type)
	case "UInt8":
		field = generator.NewUint8(meta.Name, meta.Type)
	case "UInt16":
		field = generator.NewUint16(meta.Name, meta.Type)
	case "UInt32":
		field = generator.NewUint32(meta.Name, meta.Type)
	case "UInt64":
		field = generator.NewUint64(meta.Name, meta.Type)
	case "Float32":
		field = generator.NewFloat32(meta.Name, meta.Type)
	case "Float64":
		field = generator.NewFloat64(meta.Name, meta.Type)
	case "String":
		field = generator.NewString(meta.Name, meta.Type)
	case "FixedString":
		field = generator.NewFixedString(meta.Name, meta.Type, meta.FixedStringLen)
	case "Date":
		field = generator.NewDate(meta.Name, meta.Type)
	case "DateTime":
		field = generator.NewDateTime(meta.Name, meta.Type)
	case "Enum8":
		field = generator.NewEnum8(meta.Name, meta.Type, meta.EnumData)
	case "Enum16":
		field = generator.NewEnum16(meta.Name, meta.Type, meta.EnumData)
	case "Array":
		field = generator.NewArray(meta.Name, meta.Type, Meta2Field(*meta.ArraySubtype))
	default:
		panic(fmt.Errorf("unsupported clickhouse type %s for field %s", meta.Type, meta.Name))
	}
	return
}
