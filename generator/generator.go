package generator

// Generator abstraction
type Generator interface {
	RawData(string) error
	Header() error

	////////////////////////////////////
	////////////// Production purposes

	// Types generates type definitions for all fields
	Types([]Field) error

	// TypeDef generates type definition for the field
	TypeDef(Field) error

	// HelperDef generates type helpers for the field
	HelperDef(Field) error

	// EncoderInterface generates encoding interface
	EncoderInterface([]Field) error

	// EncoderDef generates production encoder
	EncoderDef([]Field) error

	// EncodeMethod generates
	EncodingMethod([]Field) error

	////////////////////////////////////
	////////////// Production date filter

	// FilterEncoderDef generates encoder with filter
	FilterEncoderDef([]Field) error

	// FilterEncodingMethod generates
	FilterEncodingMethod([]Field) error

	////////////////////////////////////
	////////////// Void purposes

	// VoidEncoderDef generates encoder that does nothing
	VoidEncoderDef([]Field) error

	//VoidEncodingMethod
	VoidEncodingMethod([]Field) error

	////////////////////////////////////
	////////////// Testing purposes

	// TestDef generates record representation for testing purposes
	TestDef([]Field) error

	// TestEncoderDef generates encoder aimed for testing
	TestEncoderDef([]Field) error

	// TestEncodingMethod Encode method generator
	TestEncodingMethod([]Field) error

	EasyTypeName(string) string
	UneasyTypeName(string) string
	HelperName(string) string
	VarName(string) string

	Int8NativeTypeName() string
	Int16NativeTypeName() string
	Int32NativeTypeName() string
	Int64NativeTypeName() string
	Uint8NativeTypeName() string
	Uint16NativeTypeName() string
	Uint32NativeTypeName() string
	Uint64NativeTypeName() string
	Float32NativeTypeName() string
	Float64NativeTypeName() string
	StringNativeTypeName() string
	FixedStringNativeTypeName() string
	ArrayNativeTypeName(itemType Field) string

	Int8TestingTypeName() string
	Int16TestingTypeName() string
	Int32TestingTypeName() string
	Int64TestingTypeName() string
	Uint8TestingTypeName() string
	Uint16TestingTypeName() string
	Uint32TestingTypeName() string
	Uint64TestingTypeName() string
	Float32TestingTypeName() string
	Float64TestingTypeName() string
	EnumTestingTypeName() string
	DateTestingTypeName() string
	DateTimeTestingTypeName() string
	StringTestingTypeName() string
	FixedStringTestingTypeName() string
	ArrayTestingTypeName(itemType Field) string

	Int8Encoding(string) error
	Int16Encoding(string) error
	Int32Encoding(string) error
	Int64Encoding(string) error
	Uint8Encoding(string) error
	Uint16Encoding(string) error
	Uint32Encoding(string) error
	Uint64Encoding(string) error
	Float32Encoding(string) error
	Float64Encoding(string) error
	DateEncoding(string) error
	DateTimeEncoding(string) error
	StringEncoding(string) error
	FixedStringEncoding(string, int) error
	ArrayEncoding(string, Field) error

	Int8TestEncoding(string) error
	Int16TestEncoding(string) error
	Int32TestEncoding(string) error
	Int64TestEncoding(string) error
	Uint8TestEncoding(string) error
	Uint16TestEncoding(string) error
	Uint32TestEncoding(string) error
	Uint64TestEncoding(string) error
	Float32TestEncoding(string) error
	Float64TestEncoding(string) error
	EnumTestEncoding(string, map[string]int) error
	DateTestEncoding(string) error
	DateTimeTestEncoding(string) error
	StringTestEncoding(string) error
	FixedStringTestEncoding(string, int) error
	ArrayTestEncoding(string, Field) error

	EnumHelpers(Field, map[string]int) error
	DateHelpers(Field) error
	DateTimeHelpers(Field) error
}
