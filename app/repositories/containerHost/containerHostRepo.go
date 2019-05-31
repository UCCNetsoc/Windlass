package host

import (
	"context"
)

type ContainerHostRepository interface {
	WithContext(context.Context) ContainerHostRepository
	Ping() error
	CreateContainerHost(ContainerHostCreateOptions) error
}

type ContainerHostCreateOptions struct {
	Name string
}
