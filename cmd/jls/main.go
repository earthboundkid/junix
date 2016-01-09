package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/carlmjohnson/junix"
)

var (
	path = flag.String("path", ".", "Path to list the contents of")
)

func main() {
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(junix.FileInfoColumns); err != nil {
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(*path)
	if err != nil {
		log.Fatal(err) // TODO: Error handling
	}

	for _, info := range fileInfos {
		if err := enc.Encode(junix.NewFileInfo(info)); err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	flag.Parse()
	if arg := flag.Arg(0); arg != "" {
		*path = arg
	}
}
