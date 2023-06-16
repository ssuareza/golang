package memcache

import (
	"filesplit/pkg/file"
	"fmt"
	"strings"

	memcache "github.com/bradfitz/gomemcache/memcache"
)

// Client is an interface used to add compatibiliy with unit tests.
type Client interface {
	Get(key string) (*memcache.Item, error)
	Set(item *memcache.Item) error
	Delete(key string) error
}

// NewClient creates a new memcache client.
func NewClient(server string) (Client, error) {
	client := memcache.New(server)

	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

// SetFile sets a file in memcached.
func SetFile(c Client, f *file.File) error {
	// set index
	index := fmt.Sprint(f.Index)
	fmt.Println(f.Name)
	fmt.Println(f.Index)
	if err := c.Set(&memcache.Item{Key: f.Name, Value: []byte(index)}); err != nil {
		return err
	}

	// set chunks
	for _, chunk := range f.Chunks {
		if err := c.Set(&memcache.Item{Key: chunk.Key, Value: chunk.Value}); err != nil {
			return err
		}
	}

	return nil
}

// GetFile gets the file from memcached.
func GetFile(c Client, file string) ([]byte, error) {
	// get index
	i, err := c.Get(file)
	if err != nil {
		return nil, err
	}
	index := strings.Split(string(i.Value), " ")

	// merge chunks from index
	var fileBytes []byte
	for _, chunk := range index {
		c, err := c.Get(chunk)
		if err != nil {
			return nil, err
		}
		fileBytes = append(fileBytes, c.Value...)
	}

	return fileBytes, nil
}

// Delete deletes the file from memcached.
func DeleteFile(c Client, file string) error {
	// get index
	i, err := c.Get(file)
	if err != nil {
		return err
	}
	index := strings.Split(string(i.Value), " ")

	// delete chunks from index
	for _, chunk := range index {
		err := c.Delete(chunk)
		if err != nil {
			return err
		}
	}

	// delete index
	if err := c.Delete(file); err != nil {
		return err
	}

	return nil
}
