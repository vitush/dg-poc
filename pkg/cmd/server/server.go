package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/vitush/go-grpc-dg-poc/pkg/protocol/grpc"
	dmap "github.com/vitush/go-grpc-dg-poc/pkg/service"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "port", "8080", "gRPC port to bind")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	v1API := dmap.NewDMapServiceServer()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
