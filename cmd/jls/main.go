package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileInfos, _ := ioutil.ReadDir(".")
	fmt.Println(fileInfos)
}
