package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	memcachedAddr = "127.0.0.1:11211"
	expiration    = 300
)

func main() {
	keys := make(map[string]string)
	keys["dog"] = "Spot"
	keys["cat"] = "Nuti"

	// connect to memcache server
	client := memcache.New(memcachedAddr)

	// create memcache item to store
	for k, v := range keys {
		// try to pull from memcache
		fetchItem, err := client.Get(k)

		// check for cache hit
		if err != memcache.ErrCacheMiss {
			if err != nil {
				fmt.Println("Error fetching from memcache", err)
			} else {
				fmt.Printf("%s: %s (cached)\n", k, fetchItem.Value)
				continue
			}
		}

		setItem := memcache.Item{
			Key:        k,
			Value:      []byte(v),
			Expiration: expiration,
		}

		err = client.Set(&setItem)
		if err != nil {
			fmt.Println("Error setting memcache item", err)
		}
		fmt.Printf("%s: %s (new)\n", k, v)
	}
}
