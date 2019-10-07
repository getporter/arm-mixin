package main

import (
	"github.com/deislabs/porter-arm/pkg/arm"
	"github.com/spf13/cobra"
)

func buildVersionCommand(m *arm.Mixin) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the mixin version",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.PrintVersion(opts)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&opts.RawFormat, "output", "o", string(version.DefaultVersionFormat),
		"Specify an output format.  Allowed values: json, plaintext")

	return cmd
}
