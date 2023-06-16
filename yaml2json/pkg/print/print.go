package print

import (
	"fmt"

	"github.com/hokaccha/go-prettyjson"
	"gopkg.in/yaml.v2"
)

// YAML prints a yaml.
func YAML(in map[interface{}]interface{}) {
	output, _ := yaml.Marshal(in)
	fmt.Printf("%s", output)
}

// JSON prints a json.
func JSON(v interface{}) {
	output, _ := prettyjson.Marshal(v)
	fmt.Println(string(output))
}
