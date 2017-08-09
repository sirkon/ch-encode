package generator

// Generate generates encoding stuff
func Generate(gen Generator, fields *FieldSet) (err error) {
	pieceGens := []func(*FieldSet) error{
		gen.Types,
		gen.EncoderInterface,
		gen.EncoderDef,
		gen.EncodingMethod,
		gen.FilterEncoderDef,
		gen.FilterEncodingMethod,
		gen.VoidEncoderDef,
		gen.VoidEncodingMethod,
		gen.TestDef,
		gen.TestEncoderDef,
		gen.TestEncodingMethod,
	}
	for _, piece := range pieceGens {
		if err = piece(fields); err != nil {
			return
		}
	}

	return
}
