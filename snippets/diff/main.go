package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"bytes"
	)

func main() {
    // per comment, better to not read an entire file into memory
    // this is simply a trivial example.
    f1, err1 := ioutil.ReadFile("same1.txt")

    if err1 != nil {
        log.Fatal(err1)
    }

    f2, err2 := ioutil.ReadFile("same2.txt")

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println(bytes.Equal(f1, f2)) // Per comment, this is significantly more performant.
}