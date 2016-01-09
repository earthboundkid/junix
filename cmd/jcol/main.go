// jcol takes a stream of junix json blobs and pretty prints them as columns
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/carlmjohnson/junix"
)

func main() {
	var r junix.Result

	enc := json.NewDecoder(os.Stdin)
	if err := enc.Decode(&r); err != nil {
		log.Fatal(err)
	}
	if len(r.Errors) > 0 {
		log.Fatal(r.Errors)
	}

	// Make a slice of column names
	cols := make([]string, len(r.Columns))
	for i, col := range r.Columns {
		cols[i] = col.Name
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for _, col := range cols {
		fmt.Fprintf(tw, "%s\t", col)
	}
	fmt.Fprint(tw, "\n")

	for enc.More() {
		var data map[string]interface{}

		enc.Decode(&data)

		for _, col := range cols {
			fmt.Fprintf(tw, "%v\t", data[col])
		}
		fmt.Fprint(tw, "\n")
	}

	if err := tw.Flush(); err != nil {
		log.Fatal(err)
	}
}
