package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/template"
)

const (
	baseTemplate = "{{ if .errors }}{{ range .errors }}{{ . }}\n{{ end }}{{ else }}{{ range .columns }}{{ bold .name }}\t{{end}}\n{{ range .data }}{{ range . }}{{ . }}\t{{ end }}\n{{ end }}{{ end }}"

	BoldCode  = "\033[1m"
	ResetCode = "\033[0m"

)

var (
	format string
)

func bold(s string) string {
	return fmt.Sprintf("%s%s%s", BoldCode, s, ResetCode)
}

func main() {
	var jsonObjects map[string]interface{}

	enc := json.NewDecoder(os.Stdin)
	_ = enc.Decode(&jsonObjects)

	t, err := template.New("json-formatter").
		Funcs(template.FuncMap{"bold": bold}).
		Parse(baseTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, jsonObjects)
	if err != nil {
		panic(err)
	}
}

func init() {
	flag.StringVar(&format, "path", ".", "Argument to extract from input")
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		format = args[0]
	}
}
