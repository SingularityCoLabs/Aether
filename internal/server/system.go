package server

import (
	"context"
	"time"

	"connectrpc.com/connect"
	aetherv1 "github.com/SingularityCoLabs/aether/gen/go/aether/v1"
	"github.com/SingularityCoLabs/aether/internal/buildinfo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type systemService struct {
	build       buildinfo.Info
	environment string
	startedAt   time.Time
}

func (s *systemService) GetSystemInfo(
	_ context.Context,
	_ *connect.Request[aetherv1.GetSystemInfoRequest],
) (*connect.Response[aetherv1.GetSystemInfoResponse], error) {
	response := connect.NewResponse(&aetherv1.GetSystemInfoResponse{
		Name:        "aetherd",
		Version:     s.build.Version,
		Commit:      s.build.Commit,
		BuildDate:   s.build.Date,
		Environment: s.environment,
		StartedAt:   timestamppb.New(s.startedAt),
	})
	response.Header().Set("Cache-Control", "no-store")
	return response, nil
}
