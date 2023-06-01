package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/flosch/pongo2"
)

// Manifest is the kubernetes yaml object to import
type Manifest struct {
	io.Reader
}

func main() {
	input := strings.NewReader(`
	deploy:
  replicas: 2
  anothervalue: 44
config:
  api_key: 'AAAAA'
  db_user: 'BBBBB'
secrets:
  api_secret: 'CCCC'
  db_pass: 'DDDD'
	`)
	template := "template.tpl"

	c, err := NewManifest(input, template)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", c)
}

// NewManifest creates a new Config
func NewManifest(input io.Reader, template string) (Manifest, error) {
	// input
	inputBytes, err := ioutil.ReadAll(input)
	if err != nil {
		return Manifest{}, nil
	}

	var c Manifest
	err = json.Unmarshal(inputBytes, &c)
	if err != nil {
		return Manifest{}, nil
	}

	context := pongo2.Context{"var": c}

	// template
	tpl, err := pongo2.FromFile(template)
	if err != nil {
		return Manifest{}, err
	}

	// render
	var writer io.Writer
	err = tpl.ExecuteWriter(context, writer)
	if err != nil {
		return Manifest{}, err
	}

	return Manifest{}, nil
}

// Render returns the final manifest
/*func (c *Config) Render() (io.Reader, error) {
	return nil, nil
}*/
