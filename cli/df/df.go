package df

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func df() {
	cmd := exec.Command("df", "-h")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ls",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
					  love by spf13 and friends in Go.
					  Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			df()
		},
	}

	return cmd
}
