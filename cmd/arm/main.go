package main

import (
	"fmt"
	"io"
	"os"

	"get.porter.sh/mixin/arm/pkg/arm"
	"github.com/spf13/cobra"
)

func main() {
	cmd, err := buildRootCommand(os.Stdin)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(1)
	}
}

func buildRootCommand(in io.Reader) (*cobra.Command, error) {
	m, err := arm.New()
	if err != nil {
		return nil, err
	}
	m.In = in
	cmd := &cobra.Command{
		Use:  "arm",
		Long: "An ARM mixin for porter 👩🏽‍✈️",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Enable swapping out stdout/stderr for testing
			m.Out = cmd.OutOrStdout()
			m.Err = cmd.OutOrStderr()
		},
		SilenceUsage: true,
	}

	cmd.PersistentFlags().BoolVar(&m.Debug, "debug", false, "Enable debug logging")

	cmd.AddCommand(buildVersionCommand(m))
	cmd.AddCommand(buildSchemaCommand(m))
	cmd.AddCommand(buildBuildCommand(m))
	cmd.AddCommand(buildInstallCommand(m))
	cmd.AddCommand(buildUninstallCommand(m))

	return cmd, nil
}
