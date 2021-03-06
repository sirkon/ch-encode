package gogen

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/sirkon/ch-encode/internal/generator"
)

// EnumHelpers ...
func (gg *GoGen) EnumHelpers(field generator.Field, safeValues map[string]int) error {
	lines := []string{
		fmt.Sprintf("type compl%s struct {", field.TypeName(gg)),
	}
	enumItems := NewEnumItemSlice(safeValues)
	sort.Sort(enumItems)
	for _, item := range enumItems {
		lines = append(lines, fmt.Sprintf("%s %s", gg.goish.Public(item.Key), field.TypeName(gg)))
	}
	lines = append(lines, "}")
	lines = append(lines,
		fmt.Sprintf("var %s compl%s = compl%s {",
			field.AccessName(gg), field.TypeName(gg), field.TypeName(gg)))
	for _, item := range enumItems {
		lines = append(lines, fmt.Sprintf("%s: %s(%d),", gg.goish.Public(item.Key), field.TypeName(gg), item.Value))
	}
	lines = append(lines, "}")
	lines = append(lines, fmt.Sprintf("var %sEnumMapping = map[string]%s {", field.FieldName(gg), field.TypeName(gg)))
	for _, item := range enumItems {
		lines = append(lines, fmt.Sprintf(`"%s": %s.%s,`, item.Key, field.AccessName(gg), gg.goish.Public(item.Key)))
	}
	lines = append(lines, "}")
	lines = append(
		lines,
		fmt.Sprintf(
			"func (ct *compl%s) FromString(key string) (res %s, ok bool) {",
			field.TypeName(gg), field.TypeName(gg)),
	)
	lines = append(lines, fmt.Sprintf("res, ok = %sEnumMapping[key]", field.FieldName(gg)))
	lines = append(lines, "return")
	lines = append(lines, "}")

	return gg.RawData(strings.Join(lines, "\n"))
}

// DateHelpers ...
func (gg *GoGen) DateHelpers(field generator.Field) error {
	gg.useTime()
	lines := []string{
		fmt.Sprintf("type compl%s struct{}", field.TypeName(gg)),
		fmt.Sprintf(`
        func (c compl%s) FromTime(t stdtime.Time) %s {
            return %s(t.Unix()/86400)
        }
        func (c compl%s) FromTimestamp(t int64) %s {
            return %s(t/86400)
        }
        `,
			field.TypeName(gg), field.TypeName(gg), field.TypeName(gg),
			field.TypeName(gg), field.TypeName(gg), field.TypeName(gg)),
		fmt.Sprintf("var %s compl%s", field.AccessName(gg), field.TypeName(gg)),
	}
	return gg.RawData(strings.Join(lines, "\n"))
}

// DateTimeHelpers ...
func (gg *GoGen) DateTimeHelpers(field generator.Field) error {
	gg.useTime()
	lines := []string{
		fmt.Sprintf("type compl%s struct{}", field.TypeName(gg)),
		fmt.Sprintf(`
        func (c compl%s) FromTime(t stdtime.Time) %s {
            return %s(t.Unix())
        }
        func (c compl%s) FromTimestamp(t int64) %s {
            return %s(t)
        }
        `,
			field.TypeName(gg), field.TypeName(gg), field.TypeName(gg),
			field.TypeName(gg), field.TypeName(gg), field.TypeName(gg)),
		fmt.Sprintf("var %s compl%s", field.AccessName(gg), field.TypeName(gg)),
	}
	return gg.RawData(strings.Join(lines, "\n"))
}

// Dec128Helpers ...
func (gg *GoGen) Dec128Helpers(field generator.Field) error {
	format := `
	func {{.FuncName}}(lo, hi uint64) {{.TypeName}} {
		return {{.TypeName}}{
			Lo: lo,
			Hi: hi,
		}
	}

	func {{.FuncName}}Struct(item struct {Lo uint64; Hi uint64}) {{.TypeName}} {
		return {{.TypeName}}{
			Lo: item.Lo,
			Hi: item.Hi,
		}
	}
	`

	gen := template.New("test")
	gen, err := gen.Parse(format)
	if err != nil {
		panic("failed to parse template: " + err.Error())
	}
	err = gen.Execute(gg.dest, map[string]string{
		"FuncName": gg.EasyTypeName(field.AccessName(gg)),
		"TypeName": gg.UneasyTypeName(field.AccessName(gg)),
	})
	return err
}
