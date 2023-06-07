package ls

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func ls() {
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", string(out))
}

func Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ls",
		Short: "List directory contents",
		Long: `For each operand that names a file of a type other than directory, ls displays
			   its name as well as any requested, associated information.  For each operand
			   that names a file of type directory, ls displays the names of files contained
			   within that directory, as well as any requested, associated information`,
		Run: func(cmd *cobra.Command, args []string) {
			ls()
		},
	}

	return cmd
}
