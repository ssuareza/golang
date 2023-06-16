package main

import (
	"log"
	"yaml2json/pkg/format"
	"yaml2json/pkg/input"
	"yaml2json/pkg/print"
)

func main() {
	// read input from file or stdin
	input, err := input.Read()
	if err != nil {
		log.Fatal(err)
	}

	// convert yaml to json
	j, err := format.YAMLToJSON(input)
	if err != nil {
		log.Fatal(err)
	}

	// print json
	print.JSON(j)
}
