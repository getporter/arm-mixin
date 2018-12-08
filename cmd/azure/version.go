package main

import (
	"github.com/deislabs/porter-azure/pkg/azure"
	"github.com/spf13/cobra"
)

func buildVersionCommand(m *azure.Mixin) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the mixin version",
		Run: func(cmd *cobra.Command, args []string) {
			m.PrintVersion()
		},
	}
}
