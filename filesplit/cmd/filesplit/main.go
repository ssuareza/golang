package main

import (
	"filesplit/pkg/env"
	"filesplit/pkg/file"
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
	client, err := memcache.NewClient(env.MemcachedURL)
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "set":
		// initialize new file
		f, err := file.New(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		// set file
		if err := memcache.SetFile(client, f); err != nil {
			log.Fatal(err)
		}

		fmt.Println("DONE")
	case "get":
		// get file content
		content, err := memcache.GetFile(client, os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		// save to a file
		if err := file.Save(os.Args[2], content); err != nil {
			log.Fatal(err)
		}
	case "delete":
		if err := memcache.DeleteFile(client, os.Args[2]); err != nil {
			log.Fatal(err)
		}

		fmt.Println("DELETED")
	default:
		help()
	}
}
