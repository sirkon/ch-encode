/*
* THE FILE WAS GENERATED WITH logparsergen --source=testingch.script --package=util
* DO NOT TOUCH IT!
 */
package util

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"
)

type components struct {
	rest []byte
	Auth struct {
		Valid bool
		Data  []byte
	}
	Host   []byte
	Port   uint16
	DBName []byte
}

func (p *components) Parse(line []byte) (bool, error) {
	p.rest = line
	var pos int
	var tmp []byte
	var restAuth []byte
	restAuth = p.rest
	if pos = bytes.IndexByte(p.rest, '@'); pos < 0 {
		p.Auth.Valid = false
		p.rest = restAuth
		goto outAuth
	}
	p.Auth.Data = p.rest[:pos]
	p.rest = p.rest[pos+1:]
	p.Auth.Valid = true
outAuth:
	if pos = bytes.IndexByte(p.rest, ':'); pos < 0 {
		return false, fmt.Errorf("`[1m%s[0m` is not a prefix of \033[1m%s\033[0m", string(':'), string(p.rest))
	}
	p.Host = p.rest[:pos]
	p.rest = p.rest[pos+1:]
	if pos = bytes.IndexByte(p.rest, '/'); pos < 0 {
		if value, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&p.rest)), 10, 16); err == nil {
			p.Port = uint16(value)
		} else {
			return false, fmt.Errorf("Cannot convert `[1m%s[0m` into uint16 (field Port)", string(p.rest))
		}
		p.rest = p.rest[len(p.rest):]
	} else {
		tmp = p.rest[:pos]
		if value, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 16); err == nil {
			p.Port = uint16(value)
		} else {
			return false, fmt.Errorf("Cannot convert `[1m%s[0m` into uint16 (field Port)", string(p.rest[:pos]))
		}
		p.rest = p.rest[pos+1:]
	}
	p.DBName = p.rest
	p.rest = p.rest[len(p.rest):]
	return true, nil
}

func (p *components) GetAuthData() (res []byte) {
	if p.Auth.Valid {
		res = p.Auth.Data
	}
	return
}
