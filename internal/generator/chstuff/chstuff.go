package chstuff

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/sirkon/ch-encode/internal/util"
)

// FieldMeta field description
type FieldMeta struct {
	Name           string
	Type           string
	EnumData       map[string]int
	FixedStringLen int
	Decimal        struct {
		Bits      int // разрядность в битах для данного Decimal
		Precision int // количество цифр в записи числа
		Scale     int // количество цифр в дробной части
	}
	Subtype *FieldMeta
}

func retreiveField(name, ftype string) (meta FieldMeta, err error) {
	decExtractor := util.Extractor{}
	meta.Name = name
	if strings.HasPrefix(ftype, "Enum8(") && strings.HasSuffix(ftype, ")") {
		meta.Type = "Enum8"
		meta.EnumData, err = decomposeEnumArgs(ftype)
		if err != nil {
			return
		}
	} else if strings.HasPrefix(ftype, "Enum16(") && strings.HasSuffix(ftype, ")") {
		meta.Type = "Enum16"
		meta.EnumData, err = decomposeEnumArgs(ftype)
		if err != nil {
			return
		}
	} else if strings.HasPrefix(ftype, "Array(") && strings.HasSuffix(ftype, ")") {
		meta.Type = "Array"
		submeta, err := retreiveField("", ftype[6:len(ftype)-1])
		if err != nil {
			return meta, err
		}
		meta.Subtype = &submeta
	} else if strings.HasPrefix(ftype, "FixedString(") && strings.HasSuffix(ftype, ")") {
		meta.Type = "FixedString"
		submeta := ftype[len("FixedString(") : len(ftype)-1]
		length, err := strconv.ParseInt(submeta, 10, 64)
		if err != nil {
			return meta, err
		}
		meta.FixedStringLen = int(length)
	} else if strings.HasPrefix(ftype, "Nullable(") && strings.HasSuffix(ftype, ")") {
		meta.Type = "Nullable"
		submeta, err := retreiveField(name, ftype[9:len(ftype)-1])
		if err != nil {
			return meta, err
		}
		meta.Subtype = &submeta
	} else if ok, _ := decExtractor.Extract(ftype); ok {
		meta.Type = "Decimal"
		meta.Decimal.Precision = decExtractor.Precision
		meta.Decimal.Scale = decExtractor.Scale
		switch {
		case decExtractor.Precision <= 9:
			meta.Decimal.Bits = 32
			meta.Type += "32"
		case decExtractor.Precision <= 18:
			meta.Decimal.Bits = 64
			meta.Type += "64"
		case decExtractor.Precision <= 38:
			meta.Decimal.Bits = 128
			meta.Type += "128"
		}
	} else {
		meta.Type = ftype
	}
	return
}

// RetrieveTableMeta retrieves clickhouse table metainformation
func RetrieveTableMeta(conn *sql.DB, table string) (res []FieldMeta, err error) {
	rows, err := conn.Query("SELECT name, type FROM system.columns WHERE table = ?", table)
	if err != nil {
		return
	}

	var fieldName string
	var fieldType string

	for rows.Next() {
		if err := rows.Scan(&fieldName, &fieldType); err != nil {
			return res, err
		}
		meta, err := retreiveField(fieldName, fieldType)
		if err != nil {
			return nil, err
		}
		res = append(res, meta)
	}
	if rows.Err() != nil {
		return res, rows.Err()
	}
	return
}
