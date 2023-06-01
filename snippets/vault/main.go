package main

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

const (
	vaultAddr   = "http://YOUR_VAULT_ADDR:8200"
	staticToken = "YOUR_STATIC_TOKEN"
)

func main() {
	config := &api.Config{
		Address: vaultAddr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	client.SetToken(staticToken)

	// write
	secret := make(map[string]interface{})
	secret["data"] = map[string]interface{}{
		"value": "world",
		"foo":   "bar",
		"age":   "-1",
	}

	_, err = client.Logical().Write("secret/data/foo", secret)
	if err != nil {
		panic(err)
	}

	// read
	s, err := client.Logical().Read("secret/data/foo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("value:", s.Data["data"].(map[string]interface{})["value"])
	fmt.Println("foo:", s.Data["data"].(map[string]interface{})["foo"])
	fmt.Println("age:", s.Data["data"].(map[string]interface{})["age"])
}
