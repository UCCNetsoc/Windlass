package connections

import (
	"fmt"

	lxd "github.com/lxc/lxd/client"

	"github.com/Strum355/viper"
	"github.com/UCCNetworkingSociety/Windlass/auth"
	"github.com/UCCNetworkingSociety/Windlass/auth/provider"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/sqlite"
)

type ServerGroup struct {
	Database sqlbuilder.Database
	SQLite   sqlbuilder.Database
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
		sqliteConn   sqlbuilder.Database
		authProvider provider.AuthProvider
	)

	lxdConn, err = lxd.ConnectLXDUnix(viper.GetString("LXD_SOCKET"), &lxd.ConnectionArgs{
		UserAgent: "Windlass",
	})
	if err != nil {
		return ServerGroupError{"LXD", err}
	}

	mysqlConn, err = mysql.Open(mysql.ConnectionURL{
		Host:     viper.GetString("DB_HOST"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASS"),
		Database: viper.GetString("DB_NAME"),
	})
	if err != nil {
		return ServerGroupError{"MySQL", err}
	}

	authProvider, err = auth.GetProvider()
	if err != nil {
		return ServerGroupError{"Auth", err}
	}

	sqliteConn, err = sqlite.Open(sqlite.ConnectionURL{
		Database: "./sqlite.db",
	})
	if err != nil {
		return ServerGroupError{"SQLite", err}
	}

	Group = ServerGroup{
		Auth:     authProvider,
		SQLite:   sqliteConn,
		LXD:      lxdConn,
		Database: mysqlConn,
	}

	return nil
}

func (s *ServerGroup) Close() {
	//s.SQLite.Close()
	s.Database.Close()
	s.Auth.Close()
}