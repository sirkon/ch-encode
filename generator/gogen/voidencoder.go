package gogen

import (
	"fmt"

	"github.com/sirkon/ch-encode/generator"
)

// VoidEncoderDef ...
func (gg *GoGen) VoidEncoderDef(*generator.FieldSet) error {
	_, err := fmt.Fprintf(gg.dest, "\ntype %s bool\n", gg.voidEncoderName())
	return err
}

// VoidEncodingMethod ...
func (gg *GoGen) VoidEncodingMethod(fields *generator.FieldSet) error {
	_, err := fmt.Fprintf(
		gg.dest,
		`
        func (enc %s) Encode(%s) error {return nil;}
        func (enc %s) InsertCount() int {return 0;}
        func (enc %s) ErrorCount() int {return 0;}
        `,
		gg.voidEncoderName(), gg.argList(fields), gg.voidEncoderName(), gg.voidEncoderName())
	return err
}
