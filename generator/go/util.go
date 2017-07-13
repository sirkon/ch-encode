package gogen

import (
	"fmt"
	"strings"

	"github.com/DenisCheremisov/ch-encode/generator"
)

func (gg *GoGen) interfaceEncoderName() string {
	return fmt.Sprintf("%sEncoder", gg.goish.Public(gg.table))
}

func (gg *GoGen) encoderName() string {
	return fmt.Sprintf("%sRawEncoder", gg.goish.Public(gg.table))
}

func (gg *GoGen) filterEncoderName() string {
	return fmt.Sprintf("%sRawEncoderDateFilter", gg.goish.Public(gg.table))
}

func (gg *GoGen) voidEncoderName() string {
	return fmt.Sprintf("New%sRawEncoderVoid", gg.goish.Public(gg.table))
}

func (gg *GoGen) testingEncoderName() string {
	return fmt.Sprintf("Testing%sEncoder", gg.goish.Public(gg.table))
}

func (gg *GoGen) testingResultName() string {
	return fmt.Sprintf("Testing%sResult", gg.goish.Public(gg.table))
}

func (gg *GoGen) argList(fields []generator.Field) string {
	res := []string{}
	for _, field := range fields {
		res = append(res, fmt.Sprintf("%s %s", field.ArgName(gg), field.TypeName(gg)))
	}
	return strings.Join(res, ", ")
}

// EnumItem represents enumX key→value mapping
type EnumItem struct {
	Key   string
	Value int
}

// EnumItemSlice represents
type EnumItemSlice []EnumItem

// NewEnumItemSlice constructor
func NewEnumItemSlice(input map[string]int) (res EnumItemSlice) {
	for key, value := range input {
		res = append(res, EnumItem{Key: key, Value: value})
	}
	return res
}

//// Now implement sort.Interface for EnumItemSlice

// Len for sort.Interface
func (eis EnumItemSlice) Len() int { return len(eis) }

// Less for sort.Interface
func (eis EnumItemSlice) Less(i, j int) bool { return eis[i].Value < eis[j].Value }

// Swap for sort.Interface
func (eis EnumItemSlice) Swap(i, j int) {
	eis[i], eis[j] = eis[j], eis[i]
}
