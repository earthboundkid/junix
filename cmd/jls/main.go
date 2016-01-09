package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	path = flag.String("path", ".", "Path to list the contents of")
)

func main() {
	fileInfos, _ := ioutil.ReadDir(*path) // TODO: Error handling

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(NewJsonResult(fileInfos)); err != nil {
		log.Fatal(err)
	}
}

func init() {
	flag.Parse()
	if arg := flag.Arg(0); arg != "" {
		*path = arg
	}
}
