package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"connectrpc.com/connect"
	aetherv1 "github.com/SingularityCoLabs/aether/gen/go/aether/v1"
	aetherv1connect "github.com/SingularityCoLabs/aether/gen/go/aether/v1/aetherv1connect"
	"github.com/SingularityCoLabs/aether/internal/buildinfo"
	"github.com/spf13/cobra"
)

type rootOptions struct {
	serverURL string
}

func main() {
	if err := newRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	options := &rootOptions{}
	command := &cobra.Command{
		Use:           "aetherctl",
		Short:         "Control Aether through its typed API",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	command.PersistentFlags().StringVar(
		&options.serverURL,
		"server",
		envOrDefault("AETHER_SERVER_URL", "http://localhost:8080"),
		"Aether control-plane URL",
	)
	command.AddCommand(
		newSystemCommand(options),
		&cobra.Command{
			Use:   "version",
			Short: "Print build information",
			Args:  cobra.NoArgs,
			Run: func(_ *cobra.Command, _ []string) {
				info := buildinfo.Current()
				fmt.Printf("aetherctl %s (commit %s, built %s)\n", info.Version, info.Commit, info.Date)
			},
		},
	)
	return command
}

func newSystemCommand(root *rootOptions) *cobra.Command {
	command := &cobra.Command{
		Use:   "system",
		Short: "Inspect the control plane",
	}
	command.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "Read build and runtime metadata",
		Args:  cobra.NoArgs,
		RunE: func(command *cobra.Command, _ []string) error {
			return getSystemInfo(command.Context(), command, root.serverURL)
		},
	})
	return command
}

func getSystemInfo(ctx context.Context, command *cobra.Command, serverURL string) error {
	serverURL = strings.TrimRight(strings.TrimSpace(serverURL), "/")
	if serverURL == "" {
		return errors.New("server URL must not be empty")
	}
	httpClient := &http.Client{Timeout: 15 * time.Second}
	client := aetherv1connect.NewSystemServiceClient(httpClient, serverURL)
	response, err := client.GetSystemInfo(
		ctx,
		connect.NewRequest(&aetherv1.GetSystemInfoRequest{}),
	)
	if err != nil {
		return fmt.Errorf("get system info from %s: %w", serverURL, err)
	}

	output := struct {
		Name        string    `json:"name"`
		Version     string    `json:"version"`
		Commit      string    `json:"commit"`
		BuildDate   string    `json:"buildDate"`
		Environment string    `json:"environment"`
		StartedAt   time.Time `json:"startedAt"`
	}{
		Name:        response.Msg.GetName(),
		Version:     response.Msg.GetVersion(),
		Commit:      response.Msg.GetCommit(),
		BuildDate:   response.Msg.GetBuildDate(),
		Environment: response.Msg.GetEnvironment(),
		StartedAt:   response.Msg.GetStartedAt().AsTime(),
	}
	encoder := json.NewEncoder(command.OutOrStdout())
	encoder.SetIndent("", "  ")
	return encoder.Encode(output)
}

func envOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}
