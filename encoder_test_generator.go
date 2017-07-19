package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/glossina/ch-encode/util"
	"github.com/glossina/gosrcfmt"
)

// EncoderReflectionTest checks if table encoder reflects current table scheme
func EncoderReflectionTest(params util.CHParams, testPath, pkgName, tableName string) (err error) {

	// Retrieve original table
	client := &http.Client{}
	req, err := http.NewRequest("POST", params.URL(), strings.NewReader(fmt.Sprintf("DESC %s FORMAT PrettySpace", tableName)))
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Clickhouse error %d (%s): %s", resp.StatusCode, resp.Status, string(data))
	}
	sampleTable := string(data)

	testTemplate := `
package {{ .pkgName }}

import (
"testing"
"net/http"
"fmt"
"os"
"io/ioutil"
"strings"

    "github.com/DenisCheremisov/ch-encode/util"
    "github.com/sergi/go-diff/diffmatchpatch"
)
` +
		"const sampleTable = `" + sampleTable + "`" + `
func TestEncoderIntegrity(t *testing.T) {
    chparams := util.EnvCHParams()
	client := &http.Client{}
	req, err := http.NewRequest("POST", chparams.URL(), strings.NewReader(fmt.Sprintf("DESC %s FORMAT PrettySpace", "{{ .tableName }}")))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Clickhouse error %d (%s): %s", resp.StatusCode, resp.Status, string(data))
	}
	currentTable := string(data)


    if currentTable == sampleTable {
        return
    }
    dmp := diffmatchpatch.New()
    diffs := dmp.DiffMain(sampleTable, currentTable, false)
    if len(diffs) > 0 {
        fmt.Fprintf(os.Stderr,
            "Encoder table %s is out of date, update clickhouse tables or execute \033[1mgo generate parsers/libs/encoders\033[0m\n",
            "{{ .tableName }}",
        )
        fmt.Fprintln(os.Stderr, "-------------------------------------------------------------")
        fmt.Fprintln(os.Stderr, dmp.DiffPrettyText(diffs))
        t.Fatalf("Encoder and table {{ .tableName }} schemes mismatch")
    }
}
`
	t, err := template.New("encoder_test").Parse(testTemplate)
	if err != nil {
		return
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, map[string]string{
		"tableName": tableName,
		"pkgName":   pkgName,
	})
	if err != nil {
		return
	}

	dest := &bytes.Buffer{}
	gosrcfmt.FormatReader(dest, buf)
	if err = ioutil.WriteFile(testPath, dest.Bytes(), 0644); err != nil {
		return err
	}
	return
}
