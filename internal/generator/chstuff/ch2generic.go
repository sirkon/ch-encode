package chstuff

import (
	"fmt"

	"github.com/sirkon/ch-encode/internal/generator"
)

// Meta2Field translates clickhouse field metainfo into generator Field object
func Meta2Field(meta FieldMeta) (field generator.Field) {
	switch meta.Type {
	case "Int8":
		field = generator.NewInt8(meta.Name, meta.RawType)
	case "Int16":
		field = generator.NewInt16(meta.Name, meta.RawType)
	case "Int32":
		field = generator.NewInt32(meta.Name, meta.RawType)
	case "Int64":
		field = generator.NewInt64(meta.Name, meta.RawType)
	case "UInt8":
		field = generator.NewUint8(meta.Name, meta.RawType)
	case "UInt16":
		field = generator.NewUint16(meta.Name, meta.RawType)
	case "UInt32":
		field = generator.NewUint32(meta.Name, meta.RawType)
	case "UInt64":
		field = generator.NewUint64(meta.Name, meta.RawType)
	case "Float32":
		field = generator.NewFloat32(meta.Name, meta.RawType)
	case "Float64":
		field = generator.NewFloat64(meta.Name, meta.RawType)
	case "String":
		field = generator.NewString(meta.Name, meta.RawType)
	case "FixedString":
		field = generator.NewFixedString(meta.Name, meta.RawType, meta.FixedStringLen)
	case "UUID":
		field = generator.NewUUIDField(meta.Name, meta.RawType)
	case "Date":
		field = generator.NewDate(meta.Name, meta.RawType)
	case "DateTime":
		field = generator.NewDateTime(meta.Name, meta.RawType)
	case "Enum8":
		field = generator.NewEnum8(meta.Name, meta.RawType, meta.EnumData)
	case "Enum16":
		field = generator.NewEnum16(meta.Name, meta.RawType, meta.EnumData)
	case "Array":
		field = generator.NewArray(meta.Name, meta.RawType, Meta2Field(*meta.Subtype))
	case "Nullable":
		if meta.Subtype == nil {
			panic(fmt.Errorf("integrity error, nullable type must be nullable type of T, not just nullable"))
		}
		switch meta.Subtype.Type {
		case "String":
			field = generator.NewNullableString(meta.Name, meta.RawType)
		case "Array":
			field = generator.NewNullableArray(meta.Name, meta.RawType, Meta2Field(*meta.Subtype))
		default:
			field = generator.NewNullable(meta.Name, meta.RawType, Meta2Field(*meta.Subtype))
		}
	case "Decimal32":
		field = generator.NewDecimal32(meta.Name, meta.RawType, meta.Decimal.Precision, meta.Decimal.Scale)
	case "Decimal64":
		field = generator.NewDecimal64(meta.Name, meta.RawType, meta.Decimal.Precision, meta.Decimal.Scale)
	case "Decimal128":
		field = generator.NewDecimal128(meta.Name, meta.RawType, meta.Decimal.Precision, meta.Decimal.Scale)
	default:
		panic(fmt.Errorf("unsupported clickhouse type %s for field %s", meta.RawType, meta.Name))
	}
	return
}
