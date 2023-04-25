package grpcweb

// Config is grpc service config.
type Config struct {
	Cert string `json:"cert" env:"GRPCWEB_CERT"`
	Key  string `json:"key" env:"GRPCWEB_KEY"`

	ConnectionTimeout uint `json:"connection_timeout" env:"GRPCWEB_CONNECTION_TIMEOUT"`
	NumStreamWorkers  uint `json:"num_stream_workers" env:"GRPCWEB_NUM_STREAM_WORKERS"`
	MaxRecvMsgSize    uint `json:"max_recv_msg_size" env:"GRPCWEB_MAX_RECV_MSG_SIZE"`

	WebSocket bool `json:"web_socket" env:"GRPCWEB_WEB_SOCKET"`

	Origin map[string]interface{} `json:"origin" env:"GRPCWEB_ORIGIN"`

	ForceStop bool `json:"force_stop" env:"FORCE_STOP"`
}
