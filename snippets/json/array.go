package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var entry []person
	file, _ := ioutil.ReadFile("array.json")
	json.Unmarshal(file, &entry)

	for _, value := range entry {
		fmt.Println(value)
	}
}
