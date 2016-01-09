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
		// If this fails, we can't handle it with junix.Die because there is no stdout...
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(*path)
	junix.Die(enc, err)

	for _, info := range fileInfos {
		if err := enc.Encode(junix.NewFileInfo(info)); err != nil {
			junix.Die(enc, err)
		}
	}
}

func init() {
	flag.Parse()
	if arg := flag.Arg(0); arg != "" {
		*path = arg
	}
}
