package grpc

// Config is grpc service config.
type Config struct {
	Hostname string `json:"hostname" env:"GRPC_HOSTNAME" env-default:"0.0.0.0"`
	Port     string `json:"port" env:"PORT" env-default:"8080"`

	ConnectionTimeout uint `json:"connection_timeout" env:"GRPC_CONNECTION_TIMEOUT"`
	NumStreamWorkers  uint `json:"num_stream_workers" env:"GRPC_NUM_STREAM_WORKERS"`
	MaxRecvMsgSize    uint `json:"max_recv_msg_size" env:"GRPC_MAX_RECV_MSG_SIZE"`

	Origin map[string]interface{} `json:"origin" env:"GRPC_ORIGIN"`
}

// Config is grpc client config.
type ConfigClient struct {
	Hostname string `json:"hostname" env:"GRPC_CLIENT_HOSTNAME" env-default:"0.0.0.0"`
	Port     string `json:"port" env:"GRPC_CLIENT_PORT" env-default:"8080"`

	Cert string `json:"cert" env:"GRPC_CLIENT_CERT"`
	Key  string `json:"key" env:"GRPC_CLIENT_KEY"`

	MaxSendMsgSize uint `json:"max_send_msg_size" env:"GRPC_MAX_SEND_MSG_SIZE"`
}
