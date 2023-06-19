package format

import (
	yaml "gopkg.in/yaml.v2"
)

// YAMLToJSON converts yaml to json.
func YAMLToJSON(m []byte) (interface{}, error) {
	var body interface{}
	yaml.Unmarshal([]byte(m), &body)
	body = JSON(body)
	return body, nil
}

// JSON adds json format.
func JSON(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = JSON(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = JSON(v)
		}
	}
	return i
}
