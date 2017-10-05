package generator

// Generate generates encoding stuff
func Generate(gen Generator, dateField func(gen Generator) string, fields *FieldSet) (err error) {
	dateFieldName := dateField(gen)
	pieceGens := []func(*FieldSet) error{
		gen.Types,
		gen.EncoderInterface,
		gen.EncoderDef,
		gen.EncodingMethod,
	}
	if dateFieldName != "" {
		pieceGens = append(pieceGens,
			gen.DateFilterEncoderDef,
			func(f *FieldSet) error { return gen.DateFilterEncodingMethod(dateField(gen), f) },
		)
	}
	pieceGens = append(pieceGens,
		gen.VoidEncoderDef,
		gen.VoidEncodingMethod,
		gen.TestDef,
		gen.TestEncoderDef,
		gen.TestEncodingMethod,
	)
	for _, piece := range pieceGens {
		if err = piece(fields); err != nil {
			return
		}
	}

	return
}
