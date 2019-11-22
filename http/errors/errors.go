package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var router = NewRouter()

// here we define the responses based in the status code
func init() {
	router.Register(200, func(r *http.Response) {
		defer r.Body.Close()
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	})

	router.Register(400, func(r *http.Response) {
		log.Fatalln("Not found:", r.Request.URL.String())
	})
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	router.Process(resp)
}
