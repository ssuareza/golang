package main

import (
	"errors"
	"log"
)

func main() {
	log.SetFlags(0)

	err := foo(1)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func foo(i int) error {
	if 4 != i {
		return errors.New("4 is not equal to i")
	}
	return nil
}
