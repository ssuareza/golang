// Example of listing repositories in Github
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github" // with go modules disabled
	"golang.org/x/oauth2"
)

func main() {
	token := "ADD_YOUR_TOKEN_HERE"
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		fmt.Printf("%s\n", *repo.Name)
	}
}
