package chstuff

import (
	"fmt"
	"strconv"
	"strings"
)

// Enum{8,16}('arg1'=val1,....,'argN'=valN) arguments decompositor
func decomposeEnumArgs(enum string) (res map[string]int, err error) {
	res = map[string]int{}
	makeError := func(appendix string) {
		err = fmt.Errorf("%s is not an enum"+appendix, enum)
	}

	// Must be either Enum8 or Enum16
	if !(strings.HasPrefix(enum, "Enum8") || strings.HasPrefix(enum, "Enum16")) {
		makeError("")
		return
	}

	// Locate data between ( and )
	pos := strings.IndexByte(enum, '(')
	if pos < 0 {
		makeError("")
		return
	}
	data := enum[pos+1:]
	pos = strings.IndexByte(data, ')')
	if pos < 0 {
		makeError("")
		return
	}
	data = data[:pos]

	// Extract arguments
	for _, litCouple := range strings.Split(data, ",") {
		couple := strings.Split(litCouple, "=")
		if len(couple) != 2 {
			makeError(". The most probable reason is enum key value having unnacceptable characters. Try to use alnums only")
			return
		}
		key := strings.Trim(couple[0], " ")
		key = strings.TrimPrefix(key, "'")
		key = strings.TrimSuffix(key, "'")
		num, err := strconv.Atoi(strings.TrimSpace(couple[1]))
		if err != nil {
			return res, fmt.Errorf("Cannot parse numeric value `%s` for %s", couple[1], key)
		}
		res[key] = int(num)
	}
	return
}
