package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/fairyming/binary_formatter/binaryformatter"
)

var Path = flag.String("path", "", "path")

func main() {

	flag.Parse()
	data, err := os.ReadFile(*Path)
	if err != nil {
		panic(err)
	}
	tmpData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		panic(err)
	}

	result, err := binaryformatter.Parse(tmpData)
	if err != nil {
		panic(err)
	}
	fmt.Println(binaryformatter.Dump(result))
}
