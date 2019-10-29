package repo

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Strum355/log"

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
	rand.Seed(time.Now().Unix())

	services, _, err := c.consul.Health().Service("windlass_worker", "", false, new(api.QueryOptions).WithContext(ctx))
	if err != nil {
		return "", err
	}

	if len(services) == 0 {
		return "", errors.New("no windlass_worker services registered in Consul")
	}

	addrString := make([]string, 0, len(services))
	for _, s := range services {
		addrString = append(addrString, s.Service.Address+":"+strconv.Itoa(s.Service.Port))
	}

	log.WithFields(log.Fields{
		"workerAddresses": addrString,
	}).Info("fetched worker addresses from Consul")

	random := rand.Intn(len(services))

	randomService := services[random]

	return randomService.Service.Address + ":" + strconv.Itoa(randomService.Service.Port), nil
}
