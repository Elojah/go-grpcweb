package grpc

import (
	"context"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
)

// Service embed a grpc server.
type Service struct {
	Register func()

	config Config
	lis    net.Listener

	*grpc.Server
}

// Dial connects grpc server and call Register method.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	s.config = cfg

	var err error

	host := strings.Join([]string{s.config.Hostname, s.config.Port}, ":")
	s.lis, err = net.Listen("tcp", host)

	if err != nil {
		return err
	}

	s.Server = grpc.NewServer(
		grpc.ConnectionTimeout(time.Duration(cfg.ConnectionTimeout)*time.Second),
		grpc.NumStreamWorkers(uint32(cfg.NumStreamWorkers)),
		grpc.MaxRecvMsgSize(int(cfg.MaxRecvMsgSize)),
	)

	if s.Register != nil {
		s.Register()
	}

	return nil
}

func (s *Service) Close(ctx context.Context) error {
	if s.Server != nil {
		s.Server.GracefulStop()
	}

	return nil
}

func (s *Service) Serve(ctx context.Context) error {
	if s.Server != nil && s.lis != nil {
		return s.Server.Serve(s.lis)
	}

	return ErrUninitialized{}
}
