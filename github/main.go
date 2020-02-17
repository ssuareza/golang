// Example of creating a new repo in Github inside an organization and adding Admin rights to a team (just fill the variables)
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	token := "YOUR_TOKEN_HERE"
	team := "TEAM_NAME"
	organization := "ORGANIZATION"
	repository := "REPO_NAME"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)


	// get team ID
	team, _, err := client.Teams.GetTeamBySlug(ctx, organization, team)
	if err != nil {
		log.Fatal(err)
	teamID := team.GetID()

	// create repo
	r := &github.Repository{
		Name:        github.String(repository),
		Private:     github.Bool(true),
		TeamID:      &teamID,
		Permissions: &map[string]bool{"admin": true},
	}

	_, _, err = client.Repositories.Create(ctx, organization, r)
	if err != nil {
		log.Fatal(err)
	}

	// set permissions
	opts := github.TeamAddTeamRepoOptions{
		Permission: "admin",
	}
	_, err = client.Teams.AddTeamRepoBySlug(ctx, organization, team, organization, repository, &opts)
	if err != nil {
		log.Fatal(err)
	}
}
