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
		log.Fatalln(err)
	}
	fmt.Printf("%s", string(out))
}

func Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "df",
		Short: "Display free disk space",
		Long: `The df utility displays statistics about the amount 
		of free disk space on the specified filesystem or on the filesystem 
		of which file is a part.`,
		Run: func(cmd *cobra.Command, args []string) {
			df()
		},
	}

	return cmd
}
