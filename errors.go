package junix

import (
	"encoding/json"
	"fmt"
	"os"
)

func Die(enc *json.Encoder, err error) {
	if err == nil {
		return
	}
	errStr := err.Error()
	fmt.Fprintln(os.Stderr, errStr)
	_ = enc.Encode(Result{
		Errors: []string{errStr},
	})
	os.Exit(1)
}
