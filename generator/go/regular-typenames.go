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

// VarName ...
func (gg *GoGen) VarName(name string) string {
	return gg.goish.Private(name)
}
