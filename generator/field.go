package generator

// Field is a generic field representation
type Field interface {
	// FieldName returns raw field name generation as it was in a clickhouse
	FieldName(gen Generator) string

	// FieldTypeName causes a raw field type generation as it was in a clickhouse
	FieldTypeName(gen Generator) string

	// TypeName returns safe type name
	// It will be (Golang)
	//    type Access int32
	// for access Int32 clickhouse field
	// Or (C++)
	//    typedef Access int32_t;
	TypeName(gen Generator) string

	// ArgName returns  argument name to use as a TypeName parameter
	// It will be (Golang)
	//    … access Access …
	// for access clickhouse field
	// Or (C++)
	//    … Access access …
	ArgName(gen Generator) string

	// AccessName returns "namespace" name that keeps helpers.
	// For instance we have
	//   direction Enum8('to'=0, 'from'=1, 'blocked'=2)
	// field in a clickhouse instance table `table`
	// It obviously a bad idea to use codes (0, 1, 2) for encoding direction
	// In this case (C++) the TypeName will be table::DirectionType instead of table::Direction
	// and the access name will be table::Direction.
	// Then the code like
	//   namespace table {
	//     …
	//     typedef char DirectionType;
	//     namespace Direction {
	//       const DirectionType to = 0;
	//       const DirectionType from = 1;
	//       const DirectionType blocked = 2;
	//     }
	//     …
	//   }
	// The Go code is a bit more sophisticated for this piece due to its lack of nested namespaces
	// (will use special public object instead and a bit of boilerplate code to fill the data)
	AccessName(gen Generator) string

	// NativeTypeName returns native type name generation
	NativeTypeName(gen Generator) string

	// Encoding causes an encoding code generation
	Encoding(source string, gen Generator) error

	// Helper generation for the field
	Helper(gen Generator) error

	// Testing stuff

	// TestingTypeName returns a type name used for testing
	TestingTypeName(gen Generator) string

	// TestEncoding causes an encoding code generation used for testing purposes
	TestEncoding(source string, gen Generator) error
}
