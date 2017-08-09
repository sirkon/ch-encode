package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/sirkon/ch-encode/generator"
	"github.com/sirkon/ch-encode/generator/chstuff"
	"github.com/sirkon/ch-encode/generator/gogen"
	"github.com/sirkon/ch-encode/util"
	"github.com/sirkon/gosrcfmt"
	"github.com/sirkon/gotify"
	"github.com/sirkon/message"
	"github.com/urfave/cli"
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

func action(c *cli.Context) error {
	isTesting := c.Bool("test")
	yamlDict := c.String("yaml-dict")
	jsonDict := c.String("json-dict")
	if len(yamlDict) > 0 && len(jsonDict) > 0 {
		message.Criticalf("Choose one dictionary (--yaml-dict or --json-dict)")
	}
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

	for _, table := range c.Args() {
		metas, err := chstuff.RetrieveTableMeta(connect, table)
		if err != nil {
			message.Critical(err)
		}
		fields := []generator.Field{}
		for _, meta := range metas {
			fields = append(fields, chstuff.Meta2Field(meta))
		}

		partWriter := &bytes.Buffer{}
		gen := gogen.New(table, goish, partWriter)
		if err = generator.Generate(gen, fields); err != nil {
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
