/*
* THE FILE WAS GENERATED WITH logparsergen --source=testingch.script --package=util
* DO NOT TOUCH IT!
 */
package util

import (
	"bytes"
)

type userpass struct {
	rest     []byte
	User     []byte
	Password []byte
}

func (p *userpass) Parse(line []byte) (bool, error) {
	p.rest = line
	var pos int
	if pos = bytes.IndexByte(p.rest, ':'); pos < 0 {
		p.User = p.rest
		p.rest = p.rest[len(p.rest):]
	} else {
		p.User = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	}
	p.Password = p.rest
	p.rest = p.rest[len(p.rest):]
	return true, nil
}
