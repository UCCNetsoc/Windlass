package connections

import (
	consul "github.com/hashicorp/consul/api"
	vault "github.com/hashicorp/vault/api"

	"github.com/Strum355/log"
	"github.com/spf13/viper"
)

type Connections struct {
	consul *consul.Client
	vault  *vault.Client
}

var group Connections

func EstablishConnections() error {
	var (
		err error
	)

	if _, err = GetConsul(); err != nil {
		return err
	}

	log.Debug("connections established")
	return nil
}

func Close() {
}

func GetConsul() (*consul.Client, error) {
	if group.consul != nil {
		return group.consul, nil
	}

	config := consul.Config{
		Address: viper.GetString("consul.host"),
		Token:   viper.GetString("consul.token"),
	}

	client, err := consul.NewClient(&config)
	if err != nil {
		return nil, NewConnectionError(err, "Consul")
	}

	group.consul = client

	return client, nil
}
