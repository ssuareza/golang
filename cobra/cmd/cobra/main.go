package main

import (
	"cobra/pkg/df"
	"cobra/pkg/ls"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "cobra CLI",
		Long:  "cobra CLI",
	}

	cmd.AddCommand(ls.Command())
	cmd.AddCommand(df.Command())

	cmd.Execute()
}
