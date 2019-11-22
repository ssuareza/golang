package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetResponse matchs the request json response
type GetResponse struct {
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
}

// ToString print response fields
func (r *GetResponse) ToString() string {
	s := fmt.Sprintf("GET Response\nOrigin: %s\nURL: %s\n", r.Origin, r.URL)

	for k, v := range r.Headers {
		s += fmt.Sprintf("Header: %s = %s\n", k, v)
	}

	return s
}

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// json
	respContent := &GetResponse{}
	json.Unmarshal(content, respContent)
	fmt.Println(respContent.ToString())

}
