package repo

import (
	"context"
	"fmt"
	"strconv"

	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"

	"github.com/UCCNetworkingSociety/Windlass/app/connections"
	"github.com/hashicorp/consul/api"
)

type ConsulRepository struct {
	consul *api.Client
}

func NewConsulRepository() *ConsulRepository {
	consul, err := connections.GetConsul()
	if err != nil {
		panic(fmt.Sprintf("error getting consul client: %v", err))
	}

	return &ConsulRepository{
		consul: consul,
	}
}

func (c *ConsulRepository) SelectWorker(ctx context.Context) (string, error) {
	service, _, err := c.consul.Catalog().Service("windlass-worker", "", new(api.QueryOptions).WithContext(ctx))
	if err != nil {
		return "", err
	}

	log.Info("addresses %+v", service)

	return service[0].ServiceAddress + ":" + strconv.Itoa(service[0].ServicePort), nil
}
