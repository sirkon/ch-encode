package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/glossina/gotify"
)

// GoModule creates
func GoModule(goish *gotify.Gotify, table string) (writer io.WriteCloser, err error) {
	packagepath := filepath.Join(goish.Goimports(table))
	fi, err := os.Stat(packagepath)
	if err != nil {
		err = os.MkdirAll(packagepath, 0777)
		if err != nil {
			return
		}
	} else {
		if !fi.IsDir() {
			err = fmt.Errorf("%s exists and is not a folder", packagepath)
			return
		}
	}
	path := filepath.Join(packagepath, goish.Goimports(table)+".go")
	writer, err = os.Create(path)
	return
}

// GoModuleTest creates
func GoModuleTest(goish *gotify.Gotify, table string) (path string) {
	packagepath := filepath.Join(goish.Goimports(table))
	fi, err := os.Stat(packagepath)
	if err != nil {
		err = os.MkdirAll(packagepath, 0777)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if !fi.IsDir() {
			log.Fatalf("%s exists and is not a folder", packagepath)
		}
	}
	path = filepath.Join(packagepath, goish.Goimports(table)+"_test.go")
	return
}
