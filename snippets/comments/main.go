package main

import (
	"fmt"
	"regexp"
)

func main() {
	example := `# first comment
first line
# second comment
# third comment
second line
# forth comment
`

	// example = regexp.MustCompile("(?m)^[\r\n]#.*$").ReplaceAllString(example, "")
	example = regexp.MustCompile("(?m)^#.*$").ReplaceAllString(example, "")

	fmt.Printf("%v", example)
}
