package rename

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssuareza/tmdb/pkg/themoviedb"
)

var (
	errMoveMissing         = errors.New(`with "--move" you should define a destination path`)
	errNotPossibleToSearch = errors.New("not possible to perform a search")
)

func Command() *cobra.Command {
	var move bool
	var subtitles bool
	var cmd = &cobra.Command{
		Use:   "rename <movie-file>",
		Short: "Rename movie file",
		Long: `Rename movie file based on TheMovieDB database.

Example:
  tmdb rename Joker.2019.720p.BluRay.x264-[YTS.LT].avi --move /media/Movies --include-subtitles
  File renamed to "Joker (2019).avi"
  File moved to "/media/Movies/Joker (2019)/Joker (2019).avi"`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(os.Args) < 4 {
				return fmt.Errorf("%s", "you should pass a file")
			}
			if _, err := os.Stat(os.Args[2]); err != nil {
				return fmt.Errorf("file \"%s\" does not exist", os.Args[2])
			}

			if move {
				if len(os.Args) != 6 {
					return errMoveMissing
				}
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			file := os.Args[2]

			results, err := search(clean(filepath.Base(file)))
			if len(results) == 0 {
				fmt.Println("no matches found")
				os.Exit(0)
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			// rename
			newName := fmt.Sprintf("%s (%s)%s", results[0].Title, strings.Split(results[0].ReleaseDate, "-")[0], filepath.Ext(file))
			if err := os.Rename(file, filepath.Dir(file)+"/"+newName); err != nil {
				fmt.Printf("not possible to rename file to \"%s\"", newName)
				os.Exit(0)
			}
			fmt.Printf("file renamed to \"%s\"\n", newName)

			// and move
			var dstDir string
			if move {
				dstDir = filepath.Clean(os.Args[4]) + "/" + newName[0:len(newName)-len(filepath.Ext(file))]
				dst := dstDir + "/" + newName
				if err := moveFile(filepath.Dir(file)+"/"+newName, dstDir); err != nil {
					fmt.Println(err)
					os.Exit(0)
				}

				fmt.Printf("file moved to \"%s\"\n", dst)
			}

			// include subitles
			if subtitles {
				subs, err := findFilesByExtension(filepath.Dir(file), ".srt")
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}

				if len(subs) > 0 {
					for _, sub := range subs {
						if err := moveFile(sub, dstDir); err != nil {
							fmt.Printf(`not possible to move subtitle "%s"`, sub)
							os.Exit(1)
						}
					}
				}
			}
		},
	}

	cmd.Flags().BoolVarP(&move, "move", "m", false, "move file to another destination")
	cmd.Flags().BoolVarP(&subtitles, "include-subtitles", "s", false, "include subtitles")
	return cmd
}

// clean removes dots, dashes, parentheses and brackets from filename.
func clean(file string) string {
	ext := filepath.Ext(file)
	name := file[0 : len(file)-len(ext)]

	clean := strings.ReplaceAll(name, ".", " ")
	clean = strings.ReplaceAll(clean, "-", " ")
	clean = strings.ReplaceAll(clean, "(", " ")
	clean = strings.ReplaceAll(clean, ")", " ")
	clean = strings.ReplaceAll(clean, "[", " ")
	clean = strings.ReplaceAll(clean, "]", " ")
	clean = strings.ReplaceAll(clean, "  ", " ")
	return strings.TrimSpace(clean)
}

// search searches for a movie name in TheMovieDB database.
func search(name string) ([]themoviedb.Movie, error) {
	// getting year if exists
	r, _ := regexp.Compile("[1-2][0-9][0-9][0-9]")
	year := r.FindString(name)

	var (
		found bool
		words []string
		query string
	)

	for found || len(name) != 0 {
		query = strings.Replace(strings.TrimSpace(name), " ", "%20", -1)
		client := themoviedb.NewClient(themoviedb.APIKey)
		movies, err := client.SearchMovie(query, year)
		if err != nil {
			return []themoviedb.Movie{}, errNotPossibleToSearch
		}
		if len(movies) != 0 {
			return movies, nil
		}

		// prepare next iteration
		words = strings.Fields(name)
		if len(words) == 0 {
			name = ""
		} else {
			name = strings.Replace(name, words[len(words)-1], "", -1)
		}
	}

	return []themoviedb.Movie{}, nil
}

// moveFile moves a file to another destination.
func moveFile(file, dst string) error {
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return fmt.Errorf("not possible to create \"%s\"", dst)
	}

	/*
	   Move files with os.Rename() give error "invalid cross-device link" for Docker container with Volumes.
	   We fix this creating a temporary (os.Create) file and making a copy (io.Copy).
	*/
	inputFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}

	outputFile, err := os.Create(filepath.Join(dst, filepath.Base(file)))
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}

	// The copy was successful, so now delete the original file
	err = os.Remove(file)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}

	return nil
}

// findFilesByExtension returns all the files by extension.
func findFilesByExtension(path, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil

}
