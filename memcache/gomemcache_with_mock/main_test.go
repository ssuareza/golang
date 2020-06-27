package main

import (
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

type mockMemcacheClient struct{}

// mock get response
func (m *mockMemcacheClient) Get(key string) (*memcache.Item, error) {
	return &memcache.Item{
		Key:        key,
		Value:      []byte("value1"),
		Flags:      0,
		Expiration: 900,
	}, nil
}

// mock set
func (m *mockMemcacheClient) Set(item *memcache.Item) error {
	return nil
}

func TestGet(t *testing.T) {
	keys := make(map[string]string)
	keys["key1"] = "value1"

	client := &mockMemcacheClient{}

	for k, v := range keys {
		resp, _ := getKey(client, k)
		if string(resp.Value) != v {
			t.Errorf("\"%s\" key value expected is %s\n", k, v)
		}
	}
}
