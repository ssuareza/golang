package main

import (
	"github.com/spf13/cobra"
	"github.com/ssuareza/tmdb/cmd/tmdb/rename"
	"github.com/ssuareza/tmdb/cmd/tmdb/search"
)

func main() {
	cmd := &cobra.Command{
		Use:   "tmdb",
		Short: "The Movie Database CLI",
		Long:  "CLI to search movies and rename files based in TheMovieDB database.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
			}
		},
	}

	cmd.AddCommand(search.Command())
	cmd.AddCommand(rename.Command())

	cmd.Execute()
}
