package gogen

// EasyTypeName ...
func (gg *GoGen) EasyTypeName(name string) string {
	return gg.goish.Public(name)
}

// UneasyTypeName ...
func (gg *GoGen) UneasyTypeName(name string) string {
	return gg.goish.Public(name + "_type")
}

// HelperName ...
func (gg *GoGen) HelperName(name string) string {
	return gg.goish.Public(name)
}

var reservedKeywords = map[string]struct{}{
	"break": {}, "case": {}, "chan": {}, "const": {}, "continue": {}, "default": {}, "defer": {},
	"else": {}, "fallthrough": {}, "for": {}, "func": {}, "go": {}, "goto": {}, "if": {},
	"import": {}, "interface": {}, "map": {}, "package": {}, "range": {}, "return": {},
	"select": {}, "struct": {}, "switch": {}, "type": {}, "var": {},
}

// VarName ...
func (gg *GoGen) VarName(name string) string {
	if _, ok := reservedKeywords[name]; ok {
		return gg.goish.Private("thisIsReallyStrangeIfYouTookSuchANameForAField" + name)
	}
	return gg.goish.Private(name)
}
