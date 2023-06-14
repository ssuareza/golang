package memcache

import (
	"filesplit/pkg/filesplit"
	"fmt"
	"strings"

	memcache "github.com/bradfitz/gomemcache/memcache"
)

type Client struct {
	db *memcache.Client
}

func NewClient(server string) (*Client, error) {
	client := memcache.New(server)

	if err := client.Ping(); err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

// SetChunks sets the file in memcached.
func (c *Client) Set(file *filesplit.File) error {
	// set index
	index := fmt.Sprint(file.Index)
	if err := c.db.Set(&memcache.Item{Key: file.Name, Value: []byte(index)}); err != nil {
		return err
	}

	// set chunks
	for _, chunk := range file.Chunks {
		if err := c.db.Set(&memcache.Item{Key: chunk.Key, Value: chunk.Value}); err != nil {
			return err
		}
	}

	return nil
}

// Get gets the file from memcached.
func (c *Client) Get(file string) ([]byte, error) {
	// get index
	i, err := c.db.Get(file)
	if err != nil {
		return nil, err
	}
	index := strings.Split(string(i.Value), " ")

	// merge chunks from index
	var fileBytes []byte
	for _, chunk := range index {
		c, err := c.db.Get(chunk)
		if err != nil {
			return nil, err
		}
		fileBytes = append(fileBytes, c.Value...)
	}

	return fileBytes, nil
}
