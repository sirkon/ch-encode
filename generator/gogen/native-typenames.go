package gogen

import "github.com/glossina/ch-encode/generator"

// Int8NativeTypeName ...
func (gg *GoGen) Int8NativeTypeName() string {
	return "byte"
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

// ArrayNativeTypeName ...
func (gg *GoGen) ArrayNativeTypeName(itemType generator.Field) string {
	return "[]" + itemType.NativeTypeName(gg)
}
