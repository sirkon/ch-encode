package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/DenisCheremisov/ch-encode/generator"
	"github.com/DenisCheremisov/ch-encode/generator/chstuff"
	"github.com/DenisCheremisov/ch-encode/util"
	"github.com/DenisCheremisov/gotify"
	"github.com/DenisCheremisov/message"
	"github.com/go-yaml/yaml"
	"github.com/urfave/cli"

	"github.com/DenisCheremisov/ch-encode/generator/go"
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

		cmd := exec.Command("goimports")
		cmd.Stderr = os.Stderr
		writer, err := cmd.StdinPipe()
		if err != nil {
			message.Critical(err)
		}
		var output io.WriteCloser
		if !isTesting {
			output, err = GoModule(goish, table)
			cmd.Stdout = output
			if err != nil {
				message.Critical(err)
			}
		} else {
			output = os.Stdout
			cmd.Stdout = os.Stdout
		}

		gen := gogen.New(table, goish, writer)
		if err = cmd.Start(); err != nil {
			message.Critical(err)
		}
		if err = generator.Generate(gen, fields); err != nil {
			message.Critical(err)
		}
		if err = writer.Close(); err != nil {
			message.Critical(err)
		}
		if err = cmd.Wait(); err != nil {
			message.Critical(err)
		}
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
	}
	message.Infof("bos_access table encoder generated")
	return nil
}
