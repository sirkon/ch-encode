set-version:
	git tag | tail -1 | xargs printf 'package main\n\nvar version="%s"\n' > version.go