package main

import (
	"github.com/Typeform/tfcli/cmd/tf/config"
	"github.com/ansel1/merry"
	"github.com/spf13/cobra"
	"github.com/ssuareza/goland/cli/df"
	"github.com/ssuareza/goland/cli/ls"
)

var versionString = "dev"

func main() {
	cmd := &cobra.Command{
		Use:     "cli",
		Short:   config.Style.Title(`ssuareza CLI \(ᵔᵕᵔ)/`),
		Version: versionString,
		PreRun: func(cmd *cobra.Command, args []string) {
			merry.SetStackCaptureEnabled(config.Config.Debug)
		},
	}

	cmd.PersistentFlags().BoolVar(&config.Config.Debug, "debug", false, "Enable debug output")
	cmd.AddCommand(ls.Command())
	cmd.AddCommand(df.Command())

	cmd.Execute()
}
