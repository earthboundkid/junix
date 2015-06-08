package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	format string
)

func main() {
	var jsonObjects []map[string]interface{}

	enc := json.NewDecoder(os.Stdin)
	_ = enc.Decode(&jsonObjects)

	for _, obj := range jsonObjects {
		if val, ok := obj[format]; ok {
			fmt.Println(val)
		}
	}
}

func init() {
	flag.StringVar(&format, "path", ".", "Argument to extract from input")
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		format = args[0]
	}
}
