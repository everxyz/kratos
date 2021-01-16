package registry

import "context"

// Registrar is registrar interface.
type Registrar interface {
	Register(ctx context.Context, svc Service) error
	Deregister(ctx context.Context, svc Service) error
}

// Discovery is service discovery interface.
type Discovery interface {
	// GetService return the services in memory according to the service name.
	GetService(name string) ([]Service, error)
	// GetService return all services in memory.
	ListServices() ([]Service, error)
	// Resolve creates a watcher according to the service name.
	Resolve(ctx context.Context, name string) (Watcher, error)
}

// Watcher is service watcher.
type Watcher interface {
	// Watch Return services in the following three cases:
	// 1.the first time to watch and the service list is not empty
	// 2.any service changes found
	// 3.reached the context.Deadline or conetxt.Cancel
	// If the above three conditions are not met, it will block
	Watch(ctx context.Context) ([]Service, error)
	Close()
}

// Service is service interface.
type Service interface {
	// ID is the unique registration ID of service.
	ID() string
	// Name is the registration name of service.
	Name() string
	Version() string
	Metadata() map[string]string
	// Host is service hostname.
	Host() string
	// Endpoint schema: `grpc://127.0.0.1:8080?isSecure=false`.
	Endpoints() []string
}