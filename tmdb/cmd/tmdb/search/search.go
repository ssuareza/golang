package search

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssuareza/tmdb/pkg/themoviedb"
)

func Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "search <movie> [year]",
		Short: "Search movie",
		Long:  `Search movie in TheMovieDB database`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(os.Args) < 3 {
				return fmt.Errorf("%s", "arguments wrong")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			client := themoviedb.NewClient(themoviedb.APIKey)

			var year string
			if len(os.Args) == 4 {
				year = os.Args[3]
			}

			movies, err := client.SearchMovie(os.Args[2], year)
			if err != nil {
				log.Fatal(err)
			}

			if len(movies) == 0 {
				fmt.Println("no results found")
				return
			}

			for _, movie := range movies {
				year := strings.Split(movie.ReleaseDate, "-")[0]

				fmt.Printf("%s", movie.Title)
				if len(year) != 0 {
					fmt.Printf(" (%s)", year)
				}
				fmt.Println()
			}
		},
	}

	return cmd
}
