package grpc

// ErrUninitialized is raised when grpc service serve after wrong/no dial.
type ErrUninitialized struct{}

// Error implementation for ErrUninitialized.
func (err ErrUninitialized) Error() string {
	return "grpc service is not initialized"
}
