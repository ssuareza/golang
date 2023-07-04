package main

import (
	"kubecli/pkg/ckill"

	"kubecli/pkg/config"

	"kubecli/pkg/exec"

	"kubecli/pkg/get"

	"kubecli/pkg/logs"

	"github.com/ansel1/merry"
	"github.com/spf13/cobra"
)

var versionString = "dev"

func main() {
	cmd := &cobra.Command{
		Use:     "kubecli",
		Short:   config.Style.Title(`kube CLI`),
		Version: versionString,
		PreRun: func(cmd *cobra.Command, args []string) {
			merry.SetStackCaptureEnabled(config.Config.Debug)
		},
	}

	cmd.PersistentFlags().BoolVar(&config.Config.Debug, "debug", false, "Enable debug output")
	cmd.AddCommand(logs.Command())
	cmd.AddCommand(exec.Command())
	cmd.AddCommand(ckill.Command())
	cmd.AddCommand(get.Command())

	cmd.Execute()
}
