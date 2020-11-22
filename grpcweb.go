package grpcweb

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Service embed a connected redis client.
type Service struct {
	Register func()

	*grpc.Server
	*grpcweb.WrappedGrpcServer
}

// Dial connects client to external redis service.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Duration(cfg.ConnectionTimeout) * time.Second),
		grpc.NumStreamWorkers(uint32(cfg.NumStreamWorkers)),
		grpc.MaxRecvMsgSize(int(cfg.MaxRecvMsgSize)),
	}

	if cfg.Cert != "" || cfg.Key != "" {
		cert, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Key)
		if err != nil {
			return err
		}

		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	}

	s.Server = grpc.NewServer(opts...)

	s.WrappedGrpcServer = grpcweb.WrapServer(s.Server,
		grpcweb.WithOriginFunc(func(origin string) bool {
			fmt.Println("origin:", origin)
			if _, ok := cfg.Origin[origin]; !ok {
				return false
			}

			return true
		}),
	)

	s.Register()

	return nil
}

func (s *Service) Close(ctx context.Context) error {
	s.Server.GracefulStop()

	return nil
}
