package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/sirkon/ch-encode/internal/generator"
	"github.com/sirkon/ch-encode/internal/generator/chstuff"
	"github.com/sirkon/ch-encode/internal/generator/gogen"
	"github.com/sirkon/ch-encode/internal/util"
	"github.com/sirkon/gosrcfmt"
	"github.com/sirkon/gotify"
	"github.com/sirkon/message"
)

func yamlSource(path string) map[string]string {
	res := map[string]string{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		message.Criticalf("Cannot read `%s`: %s", path, err)
	}
	if err := yaml.Unmarshal(data, &res); err != nil {
		message.Criticalf("Cannot parse `%s` as YAML file: %s", path, err)
	}
	return res
}

func jsonSource(path string) map[string]string {
	res := map[string]string{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		message.Criticalf("Cannot read `%s`: %s", path, err)
	}
	if err := json.Unmarshal(data, &res); err != nil {
		message.Criticalf("Cannot parse `%s` as JSON file: %s", path, err)
	}
	return res
}

func action(isTesting bool, yamlDict string, jsonDict string, dateField string, tables []string) error {
	var dict map[string]string
	if len(yamlDict) > 0 {
		dict = yamlSource(yamlDict)
	}
	if len(jsonDict) > 0 {
		dict = jsonSource(jsonDict)
	}

	goish := gotify.New(dict)

	prms := util.EnvCHParams()
	connect, err := sql.Open("clickhouse", prms.DBURL())
	if err != nil {
		panic(err)
	}

	for _, table := range tables {
		metas, err := chstuff.RetrieveTableMeta(connect, table)
		if err != nil {
			message.Critical(err)
		}

		partWriter := &bytes.Buffer{}
		gen := gogen.New(table, goish, partWriter)
		fields := generator.NewFieldSet(gen)
		dateArg := func(gen generator.Generator) string { return "" }
		for _, meta := range metas {
			fieldInfo := chstuff.Meta2Field(meta)
			if meta.Name == dateField {
				dateArg = func(gen generator.Generator) string { return fieldInfo.ArgName(gen) }
			}
			fields.Add(fieldInfo)
		}

		if err = generator.Generate(gen, dateArg, fields); err != nil {
			message.Critical(err)
		}
		writer := &bytes.Buffer{}
		var output io.WriteCloser
		if isTesting {
			output = os.Stdout
		} else {
			output, err = GoModule(goish, table)
			if err != nil {
				message.Critical(err)
			}
		}
		gen.Header(writer)
		io.Copy(writer, partWriter)
		gosrcfmt.FormatReader(output, writer)
		if err = output.Close(); err != nil {
			message.Critical(err)
		}

		// Now go-specific testing part
		err = EncoderReflectionTest(
			prms,
			GoModuleTest(goish, table),
			goish.Package(table),
			table,
		)
		if err != nil {
			message.Critical(err)
		}
		message.Noticef("Table `\033[1m%s\033[0m` encoder generated", table)
	}
	return nil
}
