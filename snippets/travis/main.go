package main

import (
	"context"
	"log"

	"github.com/shuheiktgw/go-travis"
)

func main() {
	ctx := context.Background()
	// authenticate with personal token
	// client := travis.NewClient(travis.ApiComUrl, "YOUR_TOKEN_HERE")

	// authenticate with Github personal access token
	client := travis.NewClient(travis.ApiComUrl, "")
	_, _, err := client.Authentication.UsingGithubToken(context.Background(), "YOUR_TOKEN_HERE")

	// List organizations
	// organizations, _, err := client.Organizations.List(ctx, &travis.OrganizationsOption{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, org := range organizations {
	// 	fmt.Printf("%s\n", *org.Name)
	// }

	// List all the build ids which belongs to the current user
	// builds, _, err := client.Builds.List(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, build := range builds {
	// 	fmt.Printf("%v\n", *build.Id)
	// }

	// List repositories
	// repositories, _, err := client.Repositories.List(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, repo := range repositories {
	// 	fmt.Printf("%v\n", *repo.Name)
	// }

	// Jobs.Cancel will success
	// job := "myjob"
	// _, _, err := client.Jobs.Cancel(context.Background(), job	)

	// Add variable
	envVar := &travis.EnvVarBody{
		Name:   "sebas",
		Value:  "jajaja",
		Public: true,
	}
	_, _, err = client.EnvVars.CreateByRepoSlug(ctx, "Typeform/tfcli", envVar)
	if err != nil {
		log.Fatal(err)
	}
}
