package connections

import (
	"fmt"
	"github.com/UCCNetworkingSociety/Windlass/utils/logging"

	lxd "github.com/lxc/lxd/client"

	"github.com/Strum355/viper"
	"github.com/UCCNetworkingSociety/Windlass/app/auth"
	"github.com/UCCNetworkingSociety/Windlass/app/auth/provider"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type Connections struct {
	Database sqlbuilder.Database
	LXD      lxd.ContainerServer
	Auth     provider.AuthProvider
}

type ConnectionsError struct {
	component string
	err       error
}

func (e ConnectionsError) Error() string {
	return fmt.Sprintf("%s: %v", e.component, e.err)
}

func (e ConnectionsError) Component() string {
	return e.component
}

var Group Connections

func EstablishConnections() error {
	var (
		err          error
		lxdConn      lxd.ContainerServer
		mysqlConn    sqlbuilder.Database
		authProvider provider.AuthProvider
	)

	lxdConn, err = connectToLXD()
	if err != nil {
		return err
	}

	mysqlConn, err = connectToMySQL()
	if err != nil {
		return err
	}

	authProvider, err = auth.GetProvider()
	if err != nil {
		return ConnectionsError{"Auth", err}
	}

	Group = Connections{
		Auth:     authProvider,
		LXD:      lxdConn,
		Database: mysqlConn,
	}

	log.Debug("connections established")
	return nil
}

func (s *Connections) Close() {
	s.Database.Close()
	s.Auth.Close()
}

func connectToLXD() (lxd.ContainerServer, error) {
	lxdConn, err := lxd.ConnectLXDUnix(viper.GetString("LXD_SOCKET"), &lxd.ConnectionArgs{
		UserAgent: "Windlass",
	})
	if err != nil {
		return nil, ConnectionsError{"LXD", err}
	}
	return lxdConn, nil
}

func connectToMySQL() (sqlbuilder.Database, error) {
	opts := mysql.ConnectionURL{
		Host:     viper.GetString("DB_HOST"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASS"),
		Database: viper.GetString("DB_NAME"),
	}

	mysqlConn, err := mysql.Open(opts)
	if err != nil {
		return nil, ConnectionsError{"MySQL", err}
	}
	return mysqlConn, nil
}
