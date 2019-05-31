package services

import (
	"context"
	"fmt"

	host "github.com/UCCNetworkingSociety/Windlass/app/repositories/containerHost"
	"github.com/UCCNetworkingSociety/Windlass/app/repositories/containerHost/lxd"
	"github.com/spf13/viper"
)

type ContainerHostService struct {
	repo host.ContainerHostRepository
}

func NewContainerHostService() *ContainerHostService {
	host := &ContainerHostService{}

	provider := viper.GetString("container.host")

	if provider == "lxd" {
		host.repo = lxd.NewLXDRepository()
	} else {
		panic(fmt.Sprintf("invalid container host %s", provider))
	}

	return host
}

func (service *ContainerHostService) WithContext(ctx context.Context) *ContainerHostService {
	service.repo.WithContext(ctx)
	return service
}

func (service *ContainerHostService) CreateHost(name string) {

}
