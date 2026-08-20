package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SingularityCoLabs/aether/internal/app/controlplane"
	"github.com/SingularityCoLabs/aether/internal/buildinfo"
	"github.com/SingularityCoLabs/aether/internal/config"
	"github.com/SingularityCoLabs/aether/internal/observability"
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
		Use:           "aetherd",
		Short:         "Run the Aether control plane",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(_ *cobra.Command, _ []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}
			logger := observability.NewLogger(os.Stdout, cfg.LogLevel)
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer stop()
			return controlplane.Run(ctx, cfg, logger)
		},
	}
	command.AddCommand(versionCommand(), healthcheckCommand())
	return command
}

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print build information",
		Args:  cobra.NoArgs,
		Run: func(_ *cobra.Command, _ []string) {
			info := buildinfo.Current()
			fmt.Printf("aetherd %s (commit %s, built %s)\n", info.Version, info.Commit, info.Date)
		},
	}
}

func healthcheckCommand() *cobra.Command {
	var target string
	command := &cobra.Command{
		Use:   "healthcheck",
		Short: "Exit successfully when aetherd is ready",
		Args:  cobra.NoArgs,
		RunE: func(command *cobra.Command, _ []string) error {
			request, err := http.NewRequestWithContext(command.Context(), http.MethodGet, target, nil)
			if err != nil {
				return err
			}
			client := &http.Client{Timeout: 3 * time.Second}
			response, err := client.Do(request)
			if err != nil {
				return err
			}
			defer response.Body.Close()
			if response.StatusCode != http.StatusOK {
				return errors.New(response.Status)
			}
			return nil
		},
	}
	command.Flags().StringVar(&target, "url", "http://localhost:8080/readyz", "readiness URL")
	return command
}
