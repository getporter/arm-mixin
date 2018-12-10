package main

import (
	"github.com/deislabs/porter-azure/pkg/azure"
	"github.com/spf13/cobra"
)

var (
	commandFile string
)

func buildInstallCommand(m *azure.Mixin) *cobra.Command {
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
