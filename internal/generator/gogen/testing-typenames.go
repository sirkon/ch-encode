package gogen

import (
	"github.com/sirkon/ch-encode/internal/generator"
)

// Int8TestingTypeName ...
func (gg *GoGen) Int8TestingTypeName() string {
	return "byte"
}

// Int16TestingTypeName ...
func (gg *GoGen) Int16TestingTypeName() string {
	return "int16"
}

// Int32TestingTypeName ...
func (gg *GoGen) Int32TestingTypeName() string {
	return "int32"
}

// Int64TestingTypeName ...
func (gg *GoGen) Int64TestingTypeName() string {
	return "int64"
}

// Uint8TestingTypeName ...
func (gg *GoGen) Uint8TestingTypeName() string {
	return "byte"
}

// Uint16TestingTypeName ...
func (gg *GoGen) Uint16TestingTypeName() string {
	return "uint16"
}

// Uint32TestingTypeName ...
func (gg *GoGen) Uint32TestingTypeName() string {
	return "uint32"
}

// Uint64TestingTypeName ...
func (gg *GoGen) Uint64TestingTypeName() string {
	return "uint64"
}

func (gg *GoGen) Dec32TestingTypeName() string {
	return "string"
}

func (gg *GoGen) Dec64TestingTypeName() string {
	return "string"
}

func (gg *GoGen) Dec128TestingTypeName() string {
	return "string"
}

// Float32TestingTypeName ...
func (gg *GoGen) Float32TestingTypeName() string {
	return "float32"
}

// Float64TestingTypeName ...
func (gg *GoGen) Float64TestingTypeName() string {
	return "float64"
}

// EnumTestingTypeName ...
func (gg *GoGen) EnumTestingTypeName() string {
	return "string"
}

// DateTestingTypeName ...
func (gg *GoGen) DateTestingTypeName() string {
	return "string"
}

// DateTimeTestingTypeName ...
func (gg *GoGen) DateTimeTestingTypeName() string {
	return "string"
}

// StringTestingTypeName ...
func (gg *GoGen) StringTestingTypeName() string {
	return "string"
}

// FixedStringTestingTypeName ...
func (gg *GoGen) FixedStringTestingTypeName() string {
	return "string"
}

// ArrayTestingTypeName ...
func (gg *GoGen) ArrayTestingTypeName(itemType generator.Field) string {
	return "[]" + itemType.TestingTypeName(gg)
}

// NullableTestingTypeName ...
func (gg *GoGen) NullableTestingTypeName(itemType generator.Field) string {
	return "*" + itemType.TestingTypeName(gg)
}

// NullableStringTestingTypeName ...
func (gg *GoGen) NullableStringTestingTypeName() string {
	return "*string"
}

// NullableArrayTestingTypeName ...
func (gg *GoGen) NullableArrayTestingTypeName(itemType generator.Field) string {
	return gg.ArrayTestingTypeName(itemType)
}
