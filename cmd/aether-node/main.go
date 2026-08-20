package main

import (
	"fmt"
	"os"

	"github.com/SingularityCoLabs/aether/internal/buildinfo"
	"github.com/spf13/cobra"
)

func main() {
	if err := newRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	command := &cobra.Command{
		Use:           "aether-node",
		Short:         "Run the Aether managed-node daemon",
		Long:          "The managed-node command surface is established in Phase 0. Enrollment and heartbeat arrive in Phase 2.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	command.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print build information",
		Args:  cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			info := buildinfo.Current()
			fmt.Printf("aether-node %s (commit %s, built %s)\n", info.Version, info.Commit, info.Date)
		},
	})
	return command
}
