package grpcweb

// ErrUninitialized is raised when grpc service serve after wrong/no dial.
type ErrUninitialized struct{}

// Error implementation for ErrUninitialized.
func (err ErrUninitialized) Error() string {
	return "grpc web service is not initialized"
}
