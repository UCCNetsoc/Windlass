package connections

import (
	consul "github.com/hashicorp/consul/api"
	vault "github.com/hashicorp/vault/api"

	lxd "github.com/lxc/lxd/client"

	"github.com/UCCNetworkingSociety/Windlass/app/auth"
	"github.com/UCCNetworkingSociety/Windlass/app/auth/provider"
	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"
	"github.com/spf13/viper"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type Connections struct {
	database sqlbuilder.Database
	lxd      lxd.ContainerServer
	auth     provider.AuthProvider
	consul   *consul.Client
	vault    *vault.Client
}

var group Connections

func EstablishConnections() error {
	var (
		err error
	)

	if _, err = GetConsul(); err != nil {
		return err
	}

	if _, err = GetMySQL(); err != nil {
		return err
	}

	if _, err = auth.GetProvider(); err != nil {
		return err
	}

	if _, err = GetLXD(); err != nil {
		return err
	}

	log.Debug("connections established")
	return nil
}

func Close() {
	group.database.Close()
	group.auth.Close()
}

func GetAuth() (provider.AuthProvider, error) {
	if group.auth != nil {
		return group.auth, nil
	}

	auth, err := auth.GetProvider()
	if err != nil {
		return nil, NewConnectionError(err, "Auth")
	}

	group.auth = auth

	return auth, nil
}

func GetConsul() (*consul.Client, error) {
	if group.consul != nil {
		return group.consul, nil
	}

	config := consul.Config{
		Address: viper.GetString("consul.host"),
	}

	client, err := consul.NewClient(&config)
	if err != nil {
		return nil, NewConnectionError(err, "Consul")
	}

	group.consul = client

	return client, nil
}

func GetMySQL() (sqlbuilder.Database, error) {
	if group.database != nil {
		return group.database, nil
	}

	opts := mysql.ConnectionURL{
		Host:     viper.GetString("db.host"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.pass"),
		Database: viper.GetString("db.name"),
	}

	mysqlConn, err := mysql.Open(opts)
	if err != nil {
		return nil, NewConnectionError(err, "MySQL")
	}

	group.database = mysqlConn

	return mysqlConn, nil
}

func GetLXD() (lxd.ContainerServer, error) {
	if group.lxd != nil {
		return group.lxd, nil
	}

	lxdConn, err := lxd.ConnectLXDUnix(viper.GetString("lxd.socket"), &lxd.ConnectionArgs{
		UserAgent: "Windlass",
	})
	if err != nil {
		return nil, NewConnectionError(err, "LXD")
	}

	group.lxd = lxdConn

	return lxdConn, nil
}
