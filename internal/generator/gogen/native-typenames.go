package gogen

import (
	"github.com/sirkon/ch-encode/internal/generator"
)

// Int8NativeTypeName ...
func (gg *GoGen) Int8NativeTypeName() string {
	return "int8"
}

// Int16NativeTypeName ...
func (gg *GoGen) Int16NativeTypeName() string {
	return "int16"
}

// Int32NativeTypeName ...
func (gg *GoGen) Int32NativeTypeName() string {
	return "int32"
}

// Int64NativeTypeName ...
func (gg *GoGen) Int64NativeTypeName() string {
	return "int64"
}

// Uint8NativeTypeName ...
func (gg *GoGen) Uint8NativeTypeName() string {
	return "byte"
}

// Uint16NativeTypeName ...
func (gg *GoGen) Uint16NativeTypeName() string {
	return "uint16"
}

// Uint32NativeTypeName ...
func (gg *GoGen) Uint32NativeTypeName() string {
	return "uint32"
}

// Uint64NativeTypeName ...
func (gg *GoGen) Uint64NativeTypeName() string {
	return "uint64"
}

// Dec128NativeTypeName ...
func (gg *GoGen) Dec128NativeTypeName() string {
	return `struct {
			Lo uint64
			Hi uint64
		}`
}

// Float32NativeTypeName ...
func (gg *GoGen) Float32NativeTypeName() string {
	return "float32"
}

// Float64NativeTypeName ...
func (gg *GoGen) Float64NativeTypeName() string {
	return "float64"
}

// StringNativeTypeName ...
func (gg *GoGen) StringNativeTypeName() string {
	return "[]byte"
}

// FixedStringNativeTypeName ...
func (gg *GoGen) FixedStringNativeTypeName() string {
	return "[]byte"
}

// UUIDNativeTypeName ...
func (gg *GoGen) UUIDNativeTypeName() string {
	gg.useGoogleUUID()
	return "googleUUID.UUID"
}

// ArrayNativeTypeName ...
func (gg *GoGen) ArrayNativeTypeName(itemType generator.Field) string {
	return "[]" + itemType.NativeTypeName(gg)
}

// NullableNativeTypeName ...
func (gg *GoGen) NullableNativeTypeName(itemType generator.Field) string {
	return "*" + itemType.NativeTypeName(gg)
}

// NullableStringNativeTypeName ...
func (gg *GoGen) NullableStringNativeTypeName() string {
	return gg.StringNativeTypeName()
}

// NullableArrayNativeTypeName ...
func (gg *GoGen) NullableArrayNativeTypeName(itemType generator.Field) string {
	return gg.ArrayNativeTypeName(itemType)
}
