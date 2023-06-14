package main

import (
	"filesplit/pkg/env"
	"filesplit/pkg/filesplit"
	"filesplit/pkg/memcache"
	"flag"
	"fmt"
	"log"
	"os"
)

func help() {
	fmt.Println("Usage: go run cmd/filesplit/main.go <set|get|delete> <file>")
	flag.PrintDefaults()
	os.Exit(1)
}

func init() {
	// check args
	if len(os.Args) != 3 {
		help()
	}
	if os.Args[1] != "set" && os.Args[1] != "get" && os.Args[1] != "delete" {
		help()
	}
}

func main() {
	// initialize environment variables
	env, err := env.Init()
	if err != nil {
		panic(err)
	}

	// initialize memcache client
	cache, err := memcache.NewClient(env.MemcachedURL)
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "set":
		// split file
		file, err := filesplit.Split(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		// store chunks
		if err := cache.Set(file); err != nil {
			log.Fatal(err)
		}

		fmt.Println("DONE")
	case "get":
		// get file content
		content, err := cache.Get(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		// save to a file
		f, err := os.Create(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = f.Write(content)
		if err != nil {
			log.Fatal(err)
		}
	case "delete":
		fmt.Println("DELETE")
	default:
		help()
	}
}
