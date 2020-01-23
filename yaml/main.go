package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ssuareza/yaml2json"
)

func main() {
	f, err := ioutil.ReadFile("./example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	yamlEntries := strings.Split(string(f), "---")
	for _, yamlEntry := range yamlEntries {
		// fmt.Println(yamlEntry)
		j, err := yaml2json.YamlToJSON([]byte(fmt.Sprint(yamlEntry)))
		if err != nil {
			log.Fatal(err)
		}

		j2, err := json.Marshal(j)
		if err != nil {
			log.Fatal(err)
		}
		if fmt.Sprintf("%s", j2) == "null" {
			fmt.Println("ES NULLLLL")
		}
		fmt.Printf("%s\n", j2)

		valido := json.Valid(j2)
		fmt.Println(valido)
	}
}
