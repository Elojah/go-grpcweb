package grpc

import (
	"context"
	"crypto/tls"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	*grpc.ClientConn
}

func (c *Client) Dial(ctx context.Context, cfg ConfigClient) error {
	var err error

	// default to insecure connection except if cert/key provided in config
	secureOption := grpc.WithInsecure()

	if cfg.Cert != "" && cfg.Key != "" {
		cert, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Key)
		if err != nil {
			return err
		}

		secureOption = grpc.WithTransportCredentials(credentials.NewServerTLSFromCert(&cert))
	}

	host := strings.Join([]string{cfg.Hostname, cfg.Port}, ":")

	if c.ClientConn, err = grpc.DialContext(
		ctx,
		host,
		grpc.WithMaxMsgSize(int(cfg.MaxSendMsgSize)),
		secureOption,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.ClientConn.Close()
}
