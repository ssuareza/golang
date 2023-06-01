package main

import (
	"io/ioutil"
	"os"

	"github.com/flosch/pongo2"
	"gopkg.in/yaml.v2"
)

// Config defines the yaml structure
type Config struct {
	Deploy  map[string]string `yaml: "deploy"`
	Config  map[string]string `yaml: "config"`
	Secrets map[string]string `yaml: "secrets"`
}

func main() {
	// files
	input := "input.yaml"
	template := "template.tpl"
	output := "output.out"

	// context constructor
	yamlFile, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	var inputConfig Config
	err = yaml.Unmarshal(yamlFile, &inputConfig)
	if err != nil {
		panic(err)
	}

	context := pongo2.Context{"value": inputConfig}
	//fmt.Printf("%s\n", context)

	// write output
	err = saveToFile(template, output, context)
	if err != nil {
		panic(err)
	}
}

func saveToFile(template string, file string, context pongo2.Context) error {
	// template
	tpl, err := pongo2.FromFile(template)
	if err != nil {
		return err
	}

	// destination
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// writing
	err = tpl.ExecuteWriter(context, f)
	if err != nil {
		return err
	}

	return nil
}
