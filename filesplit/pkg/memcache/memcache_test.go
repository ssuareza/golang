package memcache

import (
	"filesplit/pkg/file"
	"testing"

	memcache "github.com/bradfitz/gomemcache/memcache"
)

type mockClient struct{}

// mock memcache client
func (c *mockClient) Set(item *memcache.Item) error {
	return nil
}
func (c *mockClient) Get(key string) (*memcache.Item, error) {
	return &memcache.Item{}, nil
}
func (c *mockClient) Delete(key string) error {
	return nil
}

func TestSetFile(t *testing.T) {
	// initializes memcache client
	client := &mockClient{}

	// initialize file
	file, err := file.New("../testdata/test.txt")
	if err != nil {
		t.Error(err)
	}

	// set file
	if err := SetFile(client, file); err != nil {
		t.Error(err)
	}
}

func TestGetFile(t *testing.T) {
	// initializes memcache client
	client := &mockClient{}

	// get file
	_, err := GetFile(client, "file.txt")
	if err != nil {
		t.Error(err)
	}
}
