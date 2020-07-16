package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

// Article contains the article
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is an array of Article
var Articles []Article

// New creates a new Article array
func New() []Article {
	return []Article{
		{ID: "1", Title: "Hello1", Desc: "Test Description1", Content: "Hello World1"},
		{ID: "2", Title: "Hello2", Desc: "Test Description2", Content: "Hello World2"},
	}
}

// GetAllArticles returns all articles
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// GetSingleArticle returns one specific article
func GetSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// CreateNewArticle creates a new article
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	// append new article
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	// and print it
	json.NewEncoder(w).Encode(article)
}

// DeleteArticle deletes an article
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.ID == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

// UpdateArticle updates an article
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	id := vars["id"]

	// get body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var modifiedArticle Article
	json.Unmarshal(reqBody, &modifiedArticle)

	// delete
	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

	// and create it again
	Articles = append(Articles, modifiedArticle)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/article", GetAllArticles).Methods("GET")
	router.HandleFunc("/article/{id}", GetSingleArticle).Methods("GET")
	router.HandleFunc("/article", CreateNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", DeleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", UpdateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	Articles = New()
	handleRequests()
}
