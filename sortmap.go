package main

import (
	"fmt"
	"sort"
)

func main() {
	service := "brownie"
	keys := make([]int, 0)
	templates := templates(service)
	for k := range templates {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for k := range keys {
		kind := templates[k][0]
		tpl := templates[k][1]
		fmt.Println(kind, tpl)
	}
}
func templates(service string) map[int][]string {
	templates := map[int][]string{
		0: []string{"config", "./scripts/templates/config.yaml.tpl"},
		1: []string{"secrets", "./scripts/templates/secrets.yaml.tpl"},
		2: []string{"database", "./" + service + "/database.yaml"},
		3: []string{"deploy", "./" + service + "/deploy.yaml"},
		4: []string{"service", "./" + service + "/service.yaml"},
	}

	return templates
}
