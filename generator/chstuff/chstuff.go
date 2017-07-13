package chstuff

import (
	"database/sql"
	"strconv"
	"strings"
)

// FieldMeta field description
type FieldMeta struct {
	Name           string
	Type           string
	EnumData       map[string]int
	FixedStringLen int
	ArraySubtype   *FieldMeta
}

func retreiveField(name, ftype string) (meta FieldMeta, err error) {
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
		meta.ArraySubtype = &submeta
	} else if strings.HasPrefix(ftype, "FixedString(") && strings.HasSuffix(ftype, ")") {
		meta.Type = "FixedString"
		submeta := ftype[len("FixedString(") : len(ftype)-1]
		length, err := strconv.ParseInt(submeta, 10, 64)
		if err != nil {
			return meta, err
		}
		meta.FixedStringLen = int(length)
	} else {
		meta.Type = ftype
	}
	return
}

// RetrieveTableMeta retrieves clickhouse table metainformation
func RetrieveTableMeta(conn *sql.DB, table string) (res []FieldMeta, err error) {
	rows, err := conn.Query("DESC " + table)
	if err != nil {
		return
	}

	var fieldName string
	var fieldType string
	var fieldDefaultType string
	var fieldDefaultValue string

	for rows.Next() {
		if err := rows.Scan(&fieldName, &fieldType, &fieldDefaultType, &fieldDefaultValue); err != nil {
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
