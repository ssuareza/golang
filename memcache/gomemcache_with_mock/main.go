package main

import (
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	memcachedAddr = "127.0.0.1:11211"
	expiration    = 300
)

// memcacheClient is an interface used to add compatibiliy with unit tests
type memcacheClient interface {
	Get(key string) (*memcache.Item, error)
	Set(item *memcache.Item) error
}

func main() {
	// key values
	keys := make(map[string]string)
	keys["dog"] = "Spot"
	keys["cat"] = "Nuti"

	client := newMemcacheClient()

	for k, v := range keys {
		// try to pull from memcache
		fetchItem, err := getKey(client, k)

		// check for cache hit
		if err != memcache.ErrCacheMiss {
			if err != nil {
				fmt.Println("Error fetching from memcache", err)
				os.Exit(1)
			} else {
				fmt.Printf("%s: %s (cached)\n", k, fetchItem.Value)
				continue
			}
		}

		// if the key is not cached let's create it!
		err = setKey(client, k, v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s: %s (new)\n", k, v)
	}
}

// newMemcacheClient creates a new memcache client
func newMemcacheClient() *memcache.Client {
	return memcache.New(memcachedAddr)
}

// getKey gets key from memcache
func getKey(client memcacheClient, key string) (*memcache.Item, error) {
	fetchItem, err := client.Get(key)
	if err != nil {
		return nil, err
	}

	return fetchItem, nil
}

// setKey set a new key in memcache
func setKey(client memcacheClient, key string, value string) error {
	setItem := memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: expiration,
	}

	if err := client.Set(&setItem); err != nil {
		fmt.Println("Error setting memcache item", err)
		return fmt.Errorf("Error setting memcache item: %s", err)
	}

	return nil
}
