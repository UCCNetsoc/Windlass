package connections

import (
	"fmt"

	lxd "github.com/lxc/lxd/client"

	"github.com/Strum355/viper"
	"github.com/UCCNetworkingSociety/Windlass/auth"
	"github.com/UCCNetworkingSociety/Windlass/auth/provider"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type ServerGroup struct {
	Database sqlbuilder.Database
	LXD      lxd.ContainerServer
	Auth     provider.AuthProvider
}

type ServerGroupError struct {
	component string
	err       error
}

func (e ServerGroupError) Error() string {
	return fmt.Sprintf("%s: %v", e.component, e.err)
}

func (e ServerGroupError) Component() string {
	return e.component
}

var Group ServerGroup

func EstablishConnections() error {
	var (
		err          error
		lxdConn      lxd.ContainerServer
		mysqlConn    sqlbuilder.Database
		authProvider provider.AuthProvider
	)

	lxdConn, err = lxd.ConnectLXDUnix(viper.GetString("LXD_SOCKET"), &lxd.ConnectionArgs{
		UserAgent: "Windlass",
	})
	if err != nil {
		return ServerGroupError{"LXD", err}
	}

	opts := mysql.ConnectionURL{
		Host:     viper.GetString("DB_HOST"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASS"),
		Database: viper.GetString("DB_NAME"),
	}

	mysqlConn, err = mysql.Open(opts)
	if err != nil {
		return ServerGroupError{"MySQL", err}
	}

	authProvider, err = auth.GetProvider()
	if err != nil {
		return ServerGroupError{"Auth", err}
	}

	Group = ServerGroup{
		Auth:     authProvider,
		LXD:      lxdConn,
		Database: mysqlConn,
	}

	return nil
}

func (s *ServerGroup) Close() {
	s.Database.Close()
	s.Auth.Close()
}
