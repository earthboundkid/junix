package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

var (
	path string
)

func main() {
	fileInfos, _ := ioutil.ReadDir(path) // TODO: Error handling

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(NewJsonResult(fileInfos))
}

func init() {
	flag.StringVar(&path, "path", ".", "Path to list the contents of")
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		path = args[0]
	}
}
