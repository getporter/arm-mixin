package main

import (
	"get.porter.sh/mixin/arm/pkg/arm"
	"github.com/spf13/cobra"
)

var (
	commandFile string
)

func buildInstallCommand(m *arm.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Execute the install functionality of this mixin",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return m.LoadConfigFromEnvironment()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Install()
		},
	}
	return cmd
}
