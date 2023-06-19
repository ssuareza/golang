package main

import (
	"fmt"
	"os"
	"yaml2json/pkg/format"
	"yaml2json/pkg/input"
	"yaml2json/pkg/print"
)

func main() {
	// read input from file or stdin
	input, err := input.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// convert yaml to json
	j, err := format.YAMLToJSON(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print json
	print.JSON(j)
}
